package models

type Response struct {
	Message string `json:"message"`
}

type ResponseError struct {
    Response
	Error interface{} `json:"error"`
}

type ResponseSuccess struct {
    Response
	Payload interface{} `json:"payload"`
}
