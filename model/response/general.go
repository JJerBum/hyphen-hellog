package response

type Genreal struct {
	Status  uint   `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}
