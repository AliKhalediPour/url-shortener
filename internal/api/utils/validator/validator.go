package validator

import "github.com/go-playground/validator/v10"

var validate = validator.New()

type ErrorResponse struct {
	Error       bool
	FailedField string
	Tag         string
	Value       interface{}
}

func Validate[T any](data T) []ErrorResponse {
	var validationErrors []ErrorResponse = nil

	errs := validate.Struct(data)
	if errs != nil {
		validationErrors = []ErrorResponse{}
		for _, err := range errs.(validator.ValidationErrors) {
			// In this case data object is actually holding the User struct
			var elem ErrorResponse

			elem.FailedField = err.Field() // Export struct field name
			elem.Tag = err.Tag()           // Export struct tag
			elem.Value = err.Value()       // Export field value
			elem.Error = true

			validationErrors = append(validationErrors, elem)
		}
	}

	return validationErrors
}
