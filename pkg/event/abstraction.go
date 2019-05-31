package event

import (
	"github.com/go-kit/kit/endpoint"
	stan "github.com/nats-io/stan.go"
)

//IEvent ...
type IEvent interface {
	Store(
		domain string,
		model string,
		eventType string,
		f endpoint.Endpoint,
	) endpoint.Endpoint
	Subscribe() *stan.Subscription
}