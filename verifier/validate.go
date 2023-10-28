package verifier

import (
	"encoding/json"
	"hyphen-hellog/cerrors"
	"hyphen-hellog/cerrors/exception"

	"github.com/go-playground/validator"
)

var validate *validator.Validate

func init() {
	validate = validator.New()
}

func Validate(model interface{}) {
	validate := validator.New()
	err := validate.Struct(model)
	if err != nil {
		var messages []map[string]interface{}
		for _, err := range err.(validator.ValidationErrors) {
			messages = append(messages, map[string]interface{}{
				"field":   err.Field(),
				"message": "this field is " + err.Tag(),
			})
		}

		jsonMessage, errJson := json.Marshal(messages)
		exception.Sniff(errJson)

		panic(cerrors.ValidationErr{
			Err: string(jsonMessage),
		})
	}
}
