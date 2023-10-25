package response

type GetSISS struct {
	Code    int    `json:"code" validate:"required"`
	Message string `json:"message" validate:"required"`
	Data    struct {
		ID string `json:"id" `
	} `json:"data"`
}

type DeleteSISS struct {
	Code    int    `json:"code" validate:"required"`
	Message string `json:"message" validate:"required"`
	Data    any    `json:"data"`
}
