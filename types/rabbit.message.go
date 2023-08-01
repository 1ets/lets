package types

type IRabbitBody interface {
	SetEvent(string)
	SetData(interface{})
	GetEvent() string
	GetData() interface{}
	Setup()
}

type RabbitBody struct {
	Payload interface{}
	Event   string      `json:"event"`
	Data    interface{} `json:"data"`
}

// Set Event Name
func (rb *RabbitBody) SetEvent(event string) {
	rb.Event = event
}

// Set Data API
func (rb *RabbitBody) SetData(data interface{}) {
	rb.Data = data
}

// Get Event Name
func (rb *RabbitBody) GetEvent() string {
	return rb.Event
}

// Get Data API
func (rb *RabbitBody) GetData() interface{} {
	return rb.Data
}

func (rb *RabbitBody) Setup() {}
