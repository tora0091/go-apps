package domain

type ResponseData struct {
	Status int         `json:"status"`
	Data   interface{} `json:"data"`
	Err    error       `json:"error"`
}
