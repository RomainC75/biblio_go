package errors

import (
	"errors"
	"strings"

	// "gitub.com/RomainC75/biblio/internal/providers/validation"

	"github.com/go-playground/validator/v10"
)

var errorsList = make(map[string]string)

func Init() {
	errorsList = map[string]string{}
}

func SetFromErrors(err error) {
	var validationErrors validator.ValidationErrors

	if errors.As(err, &validationErrors) {
		for _, fieldError := range validationErrors {
			Add(fieldError.Field(), GetErrorMsg(fieldError.Tag()))
			// ["name": "required"]
		}
	}
}

func Add(key string, value string) {
	errorsList[strings.ToLower(key)] = value
}

func GetErrorMsg(tag string) string {
	// return validation.ErrorMessages()[tag]
	return ""
}

func Get() map[string]string {
	return errorsList
}
