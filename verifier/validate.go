package verifier

import (
	"hyphen-hellog/cerrors"

	"github.com/go-playground/validator"
)

var validate *validator.Validate

func init() {
	validate = validator.New()
}

func Validate(model interface{}) {
	if err := validate.Struct(model); err != nil {
		panic(cerrors.ErrValidation)
	}

	// var messages []map[string]interface{}
	// for _, err := range err.(validator.ValidationErrors) {
	// 	messages = append(messages, map[string]interface{}{
	// 		"field":   err.Field(),
	// 		"message": "this field is " + err.Tag(),
	// 	})
	// }

	// jsonMessage, errJson := json.Marshal(messages)
	// exception.Sniff(errJson)

	// panic(cerrors.ValidationError{
	// 	ErrorMessage: string(jsonMessage),
	// })
}
