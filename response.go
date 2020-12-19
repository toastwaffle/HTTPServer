// Package response provides utilities for handlers to respond to requests.
package response

import (
	"fmt"
	"io"
	"net"
	"sort"
	"strings"

	"request"
	"status"
)

type Response interface{
	WriteResponse(net.Conn, string)
	Header(string, string) Response
}

type baseResponse struct {
	code status.Code
	headers map[string]string
}

func newBase(code status.Code) *baseResponse {
	return &baseResponse{
		code: code,
		headers: map[string]string{},
	}
}

func (br *baseResponse) WriteResponse(c net.Conn, httpVersion string) {
	fmt.Fprintf(c, "%s %d %s\r\n", httpVersion, br.code, status.Name(br.code))
	for k, v := range br.headers {
		fmt.Fprintf(c, "%s: %s\r\n", k, v)
	}
	fmt.Fprint(c, "\r\n")
}

func (br *baseResponse) Header(key, value string) Response {
	br.headers[strings.ToLower(key)] = value
	return br
}

type contentResponse struct {
	// Is a *baseResponse, but Response to make the builder pattern work.
	Response
	body string
}

func Content(code status.Code, mimeType, body string) Response {
	return &contentResponse{
		Response: newBase(code).Header("content-type", mimeType),
		body: body,
	}
}

func (cr *contentResponse) WriteResponse(c net.Conn, httpVersion string) {
	cr.Response.WriteResponse(c, httpVersion)
	fmt.Fprint(c, cr.body)
}

type streamingResponse struct {
	// Is a *baseResponse, but Response to make the builder pattern work.
	Response
	respFunc func(w io.Writer)
}

func Streaming(code status.Code, mimeType string, respFunc func(w io.Writer)) Response {
	return &streamingResponse{
		Response: newBase(code).Header("content-type", mimeType),
		respFunc: respFunc,
	}
}

func (sr *streamingResponse) WriteResponse(c net.Conn, httpVersion string) {
	sr.Response.WriteResponse(c, httpVersion)
	sr.respFunc(c)
}

func WrapErr(code status.Code, err error) Response {
	return Content(code, "text/plain", err.Error())
}

func printSortedMap(w io.Writer, m map[string]string) {
	var keys []string
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		fmt.Fprintf(w, "  %s: %s\r\n", k, m[k])
	}
}

func DumpRequest(req request.Request, variables map[string]string) Response {
	return Streaming(status.OK, "text/plain", func (w io.Writer) {
		fmt.Fprintf(w, "Method: %s\r\n", req.Method())
		fmt.Fprintf(w, "URI: %s\r\n", req.URI())
		fmt.Fprintf(w, "HTTP Version: %s\r\n", req.HTTPVersion())
		if headers := req.Headers(); len(headers) > 0 {
			fmt.Fprint(w, "Headers:\r\n")
			printSortedMap(w, headers)
		}
		if len(variables) > 0 {
			fmt.Fprint(w, "Variables:\r\n")
			printSortedMap(w, variables)
		}
		if body := req.Body(); body != "" {
			fmt.Fprintf(w, "Body:\r\n\r\n%s", body)
		}
	})
}
