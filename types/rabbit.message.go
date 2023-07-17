package types

type IRabbitBody interface {
	SetEvent(string)
	SetData(interface{})
	Setup()
}

type RabbitBody struct {
	Event string      `json:"event"`
	Data  interface{} `json:"data"`
}

func (rb *RabbitBody) SetEvent(event string) {
	rb.Event = event
}

func (rb *RabbitBody) SetData(data interface{}) {
	rb.Data = data
}

func (rb *RabbitBody) Setup() {}
