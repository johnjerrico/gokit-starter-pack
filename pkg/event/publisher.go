package event

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	stan "github.com/nats-io/stan.go"
)

// Publisher wraps a URL and provides a method that implements endpoint.Endpoint.
type Publisher struct {
	publisher stan.Conn
	logger    log.Logger
}

//NewPublisher to create new Publisher
func NewPublisher(conn stan.Conn, logger log.Logger) *Publisher {

	return &Publisher{
		publisher: conn,
		logger:    logger,
	}
}

//Store for publish event (begin and commit) to nats and data wrapping as a middleware
func (p *Publisher) Store(domain, model, eventType, subject string, f endpoint.Endpoint) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, errResponse error) {
		var requestData map[string]interface{}
		var requestBundle = make(map[string]interface{})
		data, err := json.Marshal(request)
		if err != nil {
			return nil, err
		}
		if err = json.Unmarshal(data, &requestData); err != nil {
			return nil, err
		}

		requestBundle["domain"] = domain
		requestBundle["model"] = model
		requestBundle["event_type"] = eventType
		requestBundle["data"] = requestData

		subjectNew := fmt.Sprintf("%s.%s", subject, "begin")

		dataBundle, err := json.Marshal(requestBundle)
		if err != nil {
			return nil, err
		}
		p.publisher.Publish(subjectNew, dataBundle)
		p.logger.Log("nats", "Published message on channel: "+subjectNew)
		p.logger.Log("nats", fmt.Sprintf("data : %s", requestBundle))

		defer func(_ time.Time) {
			if errResponse == nil {
				var resultData map[string]interface{}
				var resultBundle = make(map[string]interface{})
				dataResult, err := json.Marshal(response)
				if err != nil {
					p.logger.Log("error_publish_commit", err)
				}

				if err = json.Unmarshal(dataResult, &resultData); err != nil {
					p.logger.Log("error_publish_commit", err)
				}

				resultBundle["domain"] = domain
				resultBundle["model"] = model
				resultBundle["event_type"] = eventType
				resultBundle["data"] = resultData

				dataBundle, err := json.Marshal(resultBundle)
				if err != nil {
					p.logger.Log("error_publish_commit", err)
				}
				subjectNew := fmt.Sprintf("%s.%s", subject, "commit")
				p.publisher.Publish(subjectNew, dataBundle)
				p.logger.Log("nats", "Published message on channel: "+subjectNew)
				p.logger.Log("nats", fmt.Sprintf("response after : %s", resultBundle))
			}
		}(time.Now())

		return f(ctx, request)
	}
}
