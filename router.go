// Package router defines the API for routing and handling requests.
//
// Requests are routed based on trees of routing functions which consume individual path segments.
// The request path is split into path segments by splitting on slashes, ignoring any leading or
// trailing slash. Any encoded characters in the path (including slashes encoded as %2f) will be
// decoded before splitting.
package router

import (
	"errors"
	"strings"

	"request"
	"response"
	"status"
)

func extendVariables(base map[string]string, key, value string) map[string]string {
	base[key] = value
	return base
}

// A Handler handles a single request.
type Handler func(req request.Request, variables map[string]string) response.Response

// A RoutingRunc takes segments of the request path (split on slashes) to locate the appropriate
// Handler, returning nil if no appropriate handler is found. The second return value is a map of
// route variables extracted from the pathSegments, which will be passed to the Handler.
type RoutingFunc func(pathSegments []string) (Handler, map[string]string)

// CheckRoutes returns the first handler from the given routes which matches the given path
// segments. This is primarily implementation detail of the standard routing functions, but is
// exported to allow for writing custom routing functions with child routes.
func CheckRoutes(pathSegments []string, routes []RoutingFunc) (Handler, map[string]string) {
	for _, route := range routes {
		if h, vars := route(pathSegments); h != nil {
			return h, vars
		}
	}
	return nil, nil
}

// Constant consumes a single constant path segment, and then defers to child routes, using the
// first which matches.
func Constant(pathSegment string, children ...RoutingFunc) RoutingFunc {
	return func(pathSegments []string) (Handler, map[string]string) {
		if len(pathSegments) > 0 && pathSegments[0] == pathSegment {
			if h, vars := CheckRoutes(pathSegments[1:], children); h != nil {
				return h, vars
			}
		}
		return nil, nil
	}
}

// Variable consumes a single path segment, storing it as a variable with the given key, and
// then defers to child routes, using the first which matches.
func Variable(key string, children ...RoutingFunc) RoutingFunc {
	return func(pathSegments []string) (Handler, map[string]string) {
		if len(pathSegments) > 0 {
			if h, vars := CheckRoutes(pathSegments[1:], children); h != nil {
				return h, extendVariables(vars, key, pathSegments[0])
			}
		}
		return nil, nil
	}
}

// Leaf places an individual Handler at the leaf of a routing tree, verifying that no path segments
// are left unconsumed.
func Leaf(h Handler) RoutingFunc {
	return func(pathSegments []string) (Handler, map[string]string) {
		if len(pathSegments) > 0 {
			return nil, nil
		}
		return h, map[string]string{}
	}
}

// PathLeaf places an individual Handler at the leaf of a routing tree, storing the remaining
// unconsumed path segments as a variable with the given key (joined with slashes). This will not
// match if there are no unconsumed path segments.
func PathLeaf(h Handler, key string) RoutingFunc {
	return func(pathSegments []string) (Handler, map[string]string) {
		if len(pathSegments) == 0 {
			return nil, nil
		}
		return h, map[string]string{
			key: strings.Join(pathSegments, "/"),
		}
	}
}

// Router routes an incoming request.
type Router interface{
	Handle(req request.Request) response.Response
	NotFoundHandler(h func(req request.Request) response.Response)
}

type router struct{
	routes []RoutingFunc
	notFoundHandler func(req request.Request) response.Response
}

func defaultNotFoundHandler(_ request.Request) response.Response {
	return response.WrapErr(status.NotFound, errors.New("Not Found"))
}

// New creates a new router with the given routes. Incoming requests will be handled by the first
// route which matches.
func New(routes ...RoutingFunc) Router {
	return &router{routes, defaultNotFoundHandler}
}

func (r *router) Handle(req request.Request) response.Response {
	var pathSegments []string
	path := strings.Trim(req.URI().Path, "/")
	if path != "" {
		pathSegments = strings.Split(path, "/")
	}
	if h, vars := CheckRoutes(pathSegments, r.routes); h != nil {
		return h(req, vars)
	}
	return r.notFoundHandler(req)
}

func (r *router) NotFoundHandler(h func(req request.Request) response.Response) {
	r.notFoundHandler = h
}
