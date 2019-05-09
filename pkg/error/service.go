package error

import "fmt"

//Kind ...
type Kind = int

//List list of errors
type List struct {
	//BADREQUEST represents invalid syntax
	BADREQUEST Kind
	//UNAUTHORIZED  this response means "unauthenticated". That is, the client must authenticate itself to get the requested response
	UNAUTHORIZED Kind
	//PAYMENTREQUIRED This response code is reserved for future use. Initial aim for creating this code was using it for digital payment systems
	PAYMENTREQUIRED Kind
	//FORBIDDEN The client does not have access rights to the content, i.e. they are unauthorized, so server is rejecting to give proper response. Unlike 401, the client's identity is known to the server
	FORBIDDEN Kind
	//NOTFOUND The server can not find requested resource. In the browser, this means the URL is not recognized. In an API, this can also mean that the endpoint is valid but the resource itself does not exist
	NOTFOUND Kind
	//METHODNOTALLOWED The request method is known by the server but has been disabled and cannot be used. For example, an API may forbid DELETE-ing a resource. The two mandatory methods, GET and HEAD, must never be disabled and should not return this error code
	METHODNOTALLOWED Kind
	//NOTACCEPTABLE This response is sent when the web server, after performing server-driven content negotiation, doesn't find any content following the criteria given by the user agent
	NOTACCEPTABLE Kind
	//PROXYAUTHENTICATIONREQUIRED This is similar to 401 but authentication is needed to be done by a proxy
	PROXYAUTHENTICATIONREQUIRED Kind
	//REQUESTTIMEOUT This response is sent on an idle connection by some servers, even without any previous request by the client. It means that the server would like to shut down this unused connection. This response is used much more since some browsers, like Chrome, Firefox 27+, or IE9, use HTTP pre-connection mechanisms to speed up surfing. Also note that some servers merely shut down the connection without sending this message
	REQUESTTIMEOUT Kind
	//CONFLICT This response is sent when a request conflicts with the current state of the server
	CONFLICT Kind
	//GONE this response would be sent when the requested content has been permanently deleted from server, with no forwarding address. Clients are expected to remove their caches and links to the resource. The HTTP specification intends this status code to be used for "limited-time, promotional services". APIs should not feel compelled to indicate resources that have been deleted with this status code
	GONE Kind
	//LENGTHREQUIRED Server rejected the request because the Content-Length header field is not defined and the server requires it
	LENGTHREQUIRED Kind
	//PRECONDITION The client has indicated preconditions in its headers which the server does not meet
	PRECONDITIONFAILED Kind
	//PAYLOADTOOLARGE Request entity is larger than limits defined by server; the server might close the connection or return an Retry-After header field
	PAYLOADTOOLARGE Kind
	//URITOOLONG The URI requested by the client is longer than the server is willing to interpret
	URITOOLONG Kind
	//UNSUPPORTEDMEDIATYPE The media format of the requested data is not supported by the server, so the server is rejecting the request
	UNSUPPORTEDMEDIATYPE Kind
	//REQUESTEDRANGENOTSATISFIABLE The range specified by the Range header field in the request can't be fulfilled; it's possible that the range is outside the size of the target URI's data
	REQUESTEDRANGENOTSATISFIABLE Kind
	//EXPECTATIONFAILED This response code means the expectation indicated by the Expect request header field can't be met by the server
	EXPECTATIONFAILED Kind
	//UNPROCESSABLEENTITY The request was well-formed but was unable to be followed due to semantic errors
	UNPROCESSABLEENTITY Kind
	//LOCKED The resource that is being accessed is locked
	LOCKED Kind
	//FAILEDDEPENDENCY The request failed due to failure of a previous request
	FAILEDDEPENDENCY Kind
	//TOOEARLY Indicates that the server is unwilling to risk processing a request that might be replayed
	TOOEARLY Kind
	//UPGRADEREQUIRED The server refuses to perform the request using the current protocol but might be willing to do so after the client upgrades to a different protocol. The server sends an Upgrade header in a 426 response to indicate the required protocol(s)
	UPGRADEREQUIRED Kind
	//PRECONDITIONREQUIRED The origin server requires the request to be conditional. Intended to prevent the 'lost update' problem, where a client GETs a resource's state, modifies it, and PUTs it back to the server, when meanwhile a third party has modified the state on the server, leading to a conflic
	PRECONDITIONREQUIRED Kind
	//TOOMANYREQUEST The user has sent too many requests in a given amount of time ("rate limiting")
	TOOMANYREQUEST Kind
	//REQUESTHEADERFIELDSTOOLARGE The server is unwilling to process the request because its header fields are too large. The request MAY be resubmitted after reducing the size of the request header fields
	REQUESTHEADERFIELDSTOOLARGE Kind
	//UNAVAILABLEFORLEGALREASONS The user requests an illegal resource, such as a web page censored by a government
	UNAVAILABLEFORLEGALREASONS Kind
	//INTERNALSERVERERROR The server has encountered a situation it doesn't know how to handle
	INTERNALSERVERERROR Kind
	//NOTIMPLEMENTED The request method is not supported by the server and cannot be handled. The only methods that servers are required to support (and therefore that must not return this code) are GET and HEAD
	NOTIMPLEMENTED Kind
	//BADGATEWAY This error response means that the server, while working as a gateway to get a response needed to handle the request, got an invalid response
	BADGATEWAY Kind
	//SERVICEUNAVAILABLE The server is not ready to handle the request. Common causes are a server that is down for maintenance or that is overloaded. Note that together with this response, a user-friendly page explaining the problem should be sent. This responses should be used for temporary conditions and the Retry-After: HTTP header should, if possible, contain the estimated time before the recovery of the service. The webmaster must also take care about the caching-related headers that are sent along with this response, as these temporary condition responses should usually not be cached
	SERVICEUNAVAILABLE Kind
	//GATEWAYTIMEOUT This error response is given when the server is acting as a gateway and cannot get a response in time
	GATEWAYTIMEOUT Kind
	//HTTPVERSIONNOTSUPPORTED The HTTP version used in the request is not supported by the server
	HTTPVERSIONNOTSUPPORTED Kind
	//VARIANTALSONEGOTIATES The server has an internal configuration error: transparent content negotiation for the request results in a circular reference
	VARIANTALSONEGOTIATES Kind
	//INSUFFICIENTSTORAGE The server has an internal configuration error: the chosen variant resource is configured to engage in transparent content negotiation itself, and is therefore not a proper end point in the negotiation process
	INSUFFICIENTSTORAGE Kind
	//LOOPDETECTED The server detected an infinite loop while processing the request
	LOOPDETECTED Kind
	//NOTEXTENDED Further extensions to the request are required for the server to fulfill it
	NOTEXTENDED Kind
	//NETWORKAUTHENTICATIONREQUIRED The 511 status code indicates that the client needs to authenticate to gain network access
	NETWORKAUTHENTICATIONREQUIRED Kind
}

//Enum list of error
var Enum = &List{
	BADREQUEST:                    0,
	UNAUTHORIZED:                  1,
	PAYMENTREQUIRED:               2,
	FORBIDDEN:                     3,
	NOTFOUND:                      4,
	METHODNOTALLOWED:              5,
	NOTACCEPTABLE:                 6,
	PROXYAUTHENTICATIONREQUIRED:   7,
	REQUESTTIMEOUT:                8,
	CONFLICT:                      9,
	GONE:                          10,
	LENGTHREQUIRED:                11,
	PRECONDITIONFAILED:            12,
	PAYLOADTOOLARGE:               13,
	URITOOLONG:                    14,
	UNSUPPORTEDMEDIATYPE:          15,
	REQUESTEDRANGENOTSATISFIABLE:  16,
	EXPECTATIONFAILED:             17,
	UNPROCESSABLEENTITY:           18,
	LOCKED:                        19,
	FAILEDDEPENDENCY:              20,
	TOOEARLY:                      21,
	UPGRADEREQUIRED:               22,
	PRECONDITIONREQUIRED:          23,
	TOOMANYREQUEST:                24,
	REQUESTHEADERFIELDSTOOLARGE:   25,
	UNAVAILABLEFORLEGALREASONS:    26,
	INTERNALSERVERERROR:           27,
	NOTIMPLEMENTED:                28,
	BADGATEWAY:                    29,
	SERVICEUNAVAILABLE:            30,
	GATEWAYTIMEOUT:                31,
	HTTPVERSIONNOTSUPPORTED:       32,
	VARIANTALSONEGOTIATES:         33,
	INSUFFICIENTSTORAGE:           34,
	LOOPDETECTED:                  35,
	NOTEXTENDED:                   36,
	NETWORKAUTHENTICATIONREQUIRED: 37,
}

//Error encapsulate error with type of error
type Error struct {
	err     string
	kind    Kind
	message string
}

//New create new error
func New(err error, kind Kind, message string) *Error {
	return &Error{
		err:     fmt.Sprintf("%v", err),
		kind:    kind,
		message: message,
	}
}

func (d *Error) Error() string {
	return d.err
}
