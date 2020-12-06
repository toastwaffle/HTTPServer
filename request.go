// Package request implements request parsing.
package request

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"strconv"
	"strings"
)

type Request interface {
	// Placeholder method to dump the request as a HTTP 200 response.
	Dump(net.Conn)
}

type request struct {
	method string
	uri string
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

	req := &request{
		method: requestParts[0],
		uri: requestParts[1],
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

func (r *request) Dump(c net.Conn) {
	fmt.Fprintf(c, "%s 200 OK\r\n", r.httpVersion)
	fmt.Fprint(c, "Content-Type: text/plain\r\n")
	fmt.Fprint(c, "\r\n")
	fmt.Fprintf(c, "Method: %s\r\n", r.method)
	fmt.Fprintf(c, "URI: %s\r\n", r.uri)
	fmt.Fprintf(c, "HTTP Version: %s\r\n", r.httpVersion)
	if len(r.headers) > 0 {
		fmt.Fprint(c, "Headers:\r\n")
		for k, v := range r.headers {
			fmt.Fprintf(c, "  %s: %s\r\n", k, v)
		}
	}
	if r.body != "" {
		fmt.Fprintf(c, "Body:\r\n\r\n%s", r.body)
	}
}
