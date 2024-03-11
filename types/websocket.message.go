package types

type WebSocketBody struct {
	Event string      `json:"event"`
	Data  interface{} `json:"data"`
}
