package http

type GenericResponse struct {
	Header Header      `json:"header"`
	Data   interface{} `json:"data"`
}

type Header struct {
	Status       bool   `json:"status"`
	ErrorMessage string `json:"error_message"`
}
