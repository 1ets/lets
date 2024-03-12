package types

type IWebSocketBody interface {
	SetAction(string)
	SetData(interface{})
	GetAction() string
	GetData() interface{}
	Setup()
}

type WebSocketBody struct {
	Action string      `json:"action"`
	Data   interface{} `json:"data"`
}

func (wsb *WebSocketBody) SetAction(action string) {
	wsb.Action = action
}

func (wsb *WebSocketBody) SetData(data interface{}) {
	wsb.Data = data
}

func (wsb *WebSocketBody) GetAction() string {
	return wsb.Action
}

func (wsb *WebSocketBody) GetData() interface{} {
	return wsb.Data
}

func (wsb *WebSocketBody) Setup() {
}
