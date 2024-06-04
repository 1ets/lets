package types

import (
	"encoding/json"
	"fmt"
	"reflect"

	"github.com/1ets/lets"
)

type IEvent interface {
	GetName() string
	GetData() interface{}
	GetReplyTo() IReplyTo
	GetCorrelationId() string
	GetExchange() string
	GetRoutingKey() string
	GetBody() []byte
	GetDebug() bool
}

type Event struct {
	Name          string
	Exchange      string // Service exchange.
	RoutingKey    string // Service routing key.
	Data          interface{}
	ReplyTo       IReplyTo
	CorrelationId string
	Debug         bool
	Body          IRabbitBody
}

func (m *Event) GetName() string {
	return m.Name
}

func (m *Event) GetData() interface{} {
	return m.Data
}

func (m *Event) GetReplyTo() IReplyTo {
	return m.ReplyTo
}

func (m *Event) GetCorrelationId() string {
	return m.CorrelationId
}

func (m *Event) GetExchange() string {
	return m.Exchange
}

func (m *Event) GetRoutingKey() string {
	return m.RoutingKey
}

func (m *Event) GetDebug() bool {
	return m.Debug
}

func (m *Event) NilBody() bool {
	if m.Body == nil {
		return true
	}

	switch reflect.TypeOf(m.Body).Kind() {
	case reflect.Ptr, reflect.Map, reflect.Array, reflect.Chan, reflect.Slice:
		return reflect.ValueOf(m.Body).IsNil()
	}

	return false
}

func (m *Event) GetBody() []byte {
	// Check if body is not set
	if m.NilBody() {
		m.Body = &RabbitBody{}
	}

	m.Body.SetEvent(m.Name)
	m.Body.SetData(m.Data)
	m.Body.Setup()

	body, err := json.Marshal(m.Body)
	if err != nil {
		lets.LogE("RabbitEvent: %s", err.Error())
		return nil
	}

	return body
}

type IReplyTo interface {
	SetExchange(string)
	GetExchange() string
	SetRoutingKey(string)
	GetRoutingKey() string
	Get() string
}

type ReplyTo struct {
	Exchange   string `json:"exchange"`
	RoutingKey string `json:"routing_key"`
}

func (r *ReplyTo) SetExchange(exchange string) {
	r.Exchange = exchange
}

func (r *ReplyTo) GetExchange() string {
	return r.Exchange
}

func (r *ReplyTo) SetRoutingKey(routingKey string) {
	r.RoutingKey = routingKey
}

func (r *ReplyTo) GetRoutingKey() string {
	return r.RoutingKey
}

func (r *ReplyTo) Get() string {
	return r.GetJson()
}

func (r *ReplyTo) GetJson() string {
	data, err := json.Marshal(r)
	if err != nil {
		fmt.Println("Marshal ERR: ", err.Error())
	}

	return string(data)
}
