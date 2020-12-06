// Package response provides utilities for handlers to respond to requests.
package response

import (
	"bytes"
	"fmt"
	"net"
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

func WrapErr(code status.Code, err error) Response {
	return Content(code, "text/plain", err.Error())
}

func DumpRequest(req request.Request) Response {
	b := &bytes.Buffer{}
	fmt.Fprintf(b, "Method: %s\r\n", req.Method())
	fmt.Fprintf(b, "URI: %s\r\n", req.URI())
	fmt.Fprintf(b, "HTTP Version: %s\r\n", req.HTTPVersion())
	if headers := req.Headers(); len(headers) > 0 {
		fmt.Fprint(b, "Headers:\r\n")
		for k, v := range headers {
			fmt.Fprintf(b, "  %s: %s\r\n", k, v)
		}
	}
	if body := req.Body(); body != "" {
		fmt.Fprintf(b, "Body:\r\n\r\n%s", body)
	}
	return Content(status.OK, "text/plain", b.String())
}
