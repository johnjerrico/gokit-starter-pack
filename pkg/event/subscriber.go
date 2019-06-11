package event

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/go-kit/kit/log"
	stan "github.com/nats-io/stan.go"
)

// Subscriber wraps subscription to topic a URL and provides a method that implements endpoint.Endpoint.
type Subscriber struct {
	conn       stan.Conn
	subject    string
	queueGroup string
	durable    string
	startAt    string //avaliable ops : all, seqno, time, since (for more information: https://golang.org/pkg/time/#ParseDuration)
	handler    stan.MsgHandler
	logger     log.Logger
}

//NewSubscriber to create new Subscriber
func NewSubscriber(conn stan.Conn, subject string, group string, durable string, startAt string, logger log.Logger, handler stan.MsgHandler) *Subscriber {
	return &Subscriber{
		conn:       conn,
		subject:    subject,
		logger:     logger,
		queueGroup: group,
		durable:    durable,
		startAt:    startAt,
		handler:    handler,
	}
}

//Subscribe to subscribe topic to nats
func (s *Subscriber) Subscribe() *stan.Subscription {
	var startOpt stan.SubscriptionOption
	if s.startAt == "all" {
		startOpt = stan.DeliverAllAvailable()
	} else if s.startAt == "since" {
		var option = strings.Split(s.startAt, ":")
		ago, err := time.ParseDuration(option[1])
		if err != nil {
			s.logger.Log("nats", fmt.Sprintf("Error when subscribing topic %s", s.subject))
			s.logger.Log("err", err)
			return nil
		}
		startOpt = stan.StartAtTimeDelta(ago)
	} else if s.startAt == "time" {
		var option = strings.Split(s.startAt, ":")
		intTimestamp, err := strconv.ParseInt(option[1], 10, 64)
		if err != nil {
			s.logger.Log("nats", fmt.Sprintf("Error when subscribing topic %s", s.subject))
			s.logger.Log("err", err)
			return nil
		}
		startOpt = stan.StartAtTime(time.Unix(intTimestamp, 0))
	} else if s.startAt == "seqno" {
		var option = strings.Split(s.startAt, ":")
		intSeq, err := strconv.ParseUint(option[1], 10, 64)
		if err != nil {
			s.logger.Log("nats", fmt.Sprintf("Error when subscribing topic %s", s.subject))
			s.logger.Log("err", err)
			return nil
		}
		startOpt = stan.StartAtSequence(intSeq)
	}
	sub, err := s.conn.QueueSubscribe(s.subject, s.queueGroup, s.handler, stan.DurableName(s.durable), startOpt)
	if err != nil {
		s.logger.Log("nats", fmt.Sprintf("Error when subscribing topic %s", s.subject))
		return nil
	}
	s.logger.Log("nats", fmt.Sprintf("Subscribed topic %s with durable %s and start option %s", s.subject, s.durable, s.startAt))
	return &sub
}
