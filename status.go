// Package status provides HTTP status codes
package status

type Code int

const (
	// Informational
	Continue Code = 100
	SwitchingProtocols = 101
	Processing = 102
	EarlyHints = 103
	// Success
	OK = 200
	Created = 201
	Accepted = 202
	NonAuthoritativeInformation = 203
	NoContent = 204
	ResetContent = 205
	PartialContent = 206
	Multi = 207
	AlreadyReported = 208
	IMUsed = 226
	// Redirection
	MultipleChoices = 300
	MovedPermanently = 301
	Found = 302
	SeeOther = 303
	NotModified = 304
	UseProxy = 305
	SwitchProxy = 306
	TemporaryRedirect = 307
	PermanentRedirect = 308
	// Client Errors
	BadRequest = 400
	Unauthorized = 401
	PaymentRequired = 402
	Forbidden = 403
	NotFound = 404
	MethodNotAllowed = 405
	NotAcceptable = 406
	ProxyAuthenticationRequired = 407
	RequestTimeout = 408
	Conflict = 409
	Gone = 410
	LengthRequired = 411
	PreconditionFailed = 412
	PayloadTooLarge = 413
	URITooLong = 414
	UnsupportedMediaType = 415
	RangeNotSatisfiable = 416
	ExpectationFailed = 417
	IAmATeapot = 418
	MisdirectedRequest = 421
	UnprocessableEntity = 422
	Locked = 423
	FailedDependency = 424
	TooEarly = 425
	UpgradeRequired = 426
	PreconditionRequired = 428
	TooManyRequests = 429
	RequestHeaderFieldsTooLarge = 431
	UnavailableForLegalReasons = 451
	// Server Errors
	InternalServerError = 500
	NotImplemented = 501
	BadGateway = 502
	ServiceUnavailable = 503
	GatewayTimeout = 504
	HTTPVersionNotSupported = 505
	VariantAlsoNegotiates = 506
	InsufficientStorage = 507
	LoopDetected = 508
	NotExtended = 510
	NetworkAuthenticationRequired = 511
)

var names = map[Code]string{
	Continue: "Continue",
	SwitchingProtocols: "Switching Protocols",
	Processing: "Processing",
	EarlyHints: "Early Hints",
	OK: "OK",
	Created: "Created",
	Accepted: "Accepted",
	NonAuthoritativeInformation: "Non-Authoritative Information",
	NoContent: "No Content",
	ResetContent: "Reset Content",
	PartialContent: "Partial Content",
	Multi: "Multi",
	AlreadyReported: "Already Reported",
	IMUsed: "IM Used",
	MultipleChoices: "Multiple Choices",
	MovedPermanently: "Moved Permanently",
	Found: "Found",
	SeeOther: "See Other",
	NotModified: "Not Modified",
	UseProxy: "Use Proxy",
	SwitchProxy: "Switch Proxy",
	TemporaryRedirect: "Temporary Redirect",
	PermanentRedirect: "Permanent Redirect",
	BadRequest: "Bad Request",
	Unauthorized: "Unauthorized",
	PaymentRequired: "Payment Required",
	Forbidden: "Forbidden",
	NotFound: "Not Found",
	MethodNotAllowed: "Method Not Allowed",
	NotAcceptable: "Not Acceptable",
	ProxyAuthenticationRequired: "Proxy Authentication Required",
	RequestTimeout: "Request Timeout",
	Conflict: "Conflict",
	Gone: "Gone",
	LengthRequired: "Length Required",
	PreconditionFailed: "Precondition Failed",
	PayloadTooLarge: "Payload Too Large",
	URITooLong: "URI Too Long",
	UnsupportedMediaType: "Unsupported Media Type",
	RangeNotSatisfiable: "Range Not Satisfiable",
	ExpectationFailed: "Expectation Failed",
	IAmATeapot: "I'm A Teapot",
	MisdirectedRequest: "Misdirected Request",
	UnprocessableEntity: "Unprocessable Entity",
	Locked: "Locked",
	FailedDependency: "Failed Dependency",
	TooEarly: "Too Early",
	UpgradeRequired: "Upgrade Required",
	PreconditionRequired: "Precondition Required",
	TooManyRequests: "Too Many Requests",
	RequestHeaderFieldsTooLarge: "Request Header Fields Too Large",
	UnavailableForLegalReasons: "Unavailable For Legal Reasons",
	InternalServerError: "Internal Server Error",
	NotImplemented: "Not Implemented",
	BadGateway: "Bad Gateway",
	ServiceUnavailable: "Service Unavailable",
	GatewayTimeout: "Gateway Timeout",
	HTTPVersionNotSupported: "HTTP Version Not Supported",
	VariantAlsoNegotiates: "Variant Also Negotiates",
	InsufficientStorage: "Insufficient Storage",
	LoopDetected: "Loop Detected",
	NotExtended: "Not Extended",
	NetworkAuthenticationRequired: "Network Authentication Required",
}

func Name(c Code) string {
	return names[c]
}
