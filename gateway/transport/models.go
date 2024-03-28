package transport

type ResponseWithHash struct {
	Hash string `json:"hash"`
}

type ResponseExists struct {
	Exists bool `json:"exists"`
}

type RequestWithPayload struct {
	Payload string `json:"payload"`
}
