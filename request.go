// Package request implements request parsing.
package request

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"net/url"
	"strconv"
	"strings"
)

type Request interface {
	Method() string
	URI() *url.URL
	HTTPVersion() string
	Headers() map[string]string
	Body() string
}

type request struct {
	method string
	uri *url.URL
	httpVersion string
	headers map[string]string
	body string
}

func Parse(c net.Conn) (Request, error) {
	r := bufio.NewReader(c)

	// Request-Line = Method SP Request-URI SP HTTP-Version CRLF
	// https://tools.ietf.org/html/rfc2616#section-5.1
	requestLine, err := r.ReadString(0x0a)
	if err != nil {
		return nil, err
	}
	requestParts := strings.Split(strings.TrimSpace(requestLine), " ")
	if len(requestParts) != 3 {
		return nil, fmt.Errorf("request line %q did not contain 3 parts", requestLine)
	}
	uri, err := url.ParseRequestURI(requestParts[1])
	if err != nil {
		return nil, err
	}

	req := &request{
		method: requestParts[0],
		uri: uri,
		httpVersion: requestParts[2],
		headers: map[string]string{},
	}

	// Read headers line by line
	for {
		line, err := r.ReadString(0x0a)
		if err != nil {
			return nil, err
		}
		line = strings.TrimSpace(line)
		if line == "" {
			// Headers terminated with a blank line
			break
		}
		parts := strings.SplitN(line, ":", 2)
		if len(parts) != 2 {
			return nil, fmt.Errorf("header line %q did not contain 2 parts", line)
		}
		req.headers[strings.ToLower(strings.TrimSpace(parts[0]))] = strings.TrimSpace(parts[1])
	}

	// Body is present iff Content-Length or Transfer-Encoding header is present
	// https://tools.ietf.org/html/rfc2616#section-4.4
	// TODO: support Transfer-Encoding.
	if lengthStr, ok := req.headers["content-length"]; ok {
		length, err := strconv.Atoi(lengthStr)
		if err != nil {
			return nil, err
		}
		body := make([]byte, length, length)
		if _, err := io.ReadFull(r, body); err != nil {
			return nil, err
		}
		req.body = string(body)
	}

	return req, nil
}

func (r *request) Method() string {
	return r.method
}

func (r *request) URI() *url.URL {
	return r.uri
}

func (r *request) HTTPVersion() string {
	return r.httpVersion
}

func (r *request) Headers() map[string]string {
	return r.headers
}

func (r *request) Body() string {
	return r.body
}
