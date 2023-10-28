package model

type InGetSISS struct {
	Code    int    `json:"code" validate:"required"`
	Message string `json:"message" validate:"required"`
	Data    struct {
		ID string `json:"id" `
	} `json:"data"`
}

type InDeleteSISS struct {
	Code    int    `json:"code" validate:"required"`
	Message string `json:"message" validate:"required"`
	Data    any    `json:"data"`
}
