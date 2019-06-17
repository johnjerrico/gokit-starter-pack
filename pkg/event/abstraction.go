package event

import (
	"github.com/go-kit/kit/endpoint"
	stan "github.com/nats-io/stan.go"
)

//IEvent ...
type IEvent interface {
	Store(
		domain, model, eventType, subject, eventSource string,
		f endpoint.Endpoint,
		metaBuilder MetaBuilder,
	) endpoint.Endpoint
	Subscribe() *stan.Subscription
}
