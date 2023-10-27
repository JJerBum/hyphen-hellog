package response

type Genreal struct {
	Code    uint   `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}
