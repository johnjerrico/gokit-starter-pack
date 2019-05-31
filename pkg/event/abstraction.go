package event

import (
	"github.com/go-kit/kit/endpoint"
)

//IEvent ...
type IEvent interface {
	Store(
		domain string,
		model string,
		eventType string,
		f endpoint.Endpoint,
	) endpoint.Endpoint
}