package validation

import (
	"encoding/json"
	"errors"

	"github.com/Kovokar/estudos-go/tree/main/apis/crud-go/src/configuration/rest_err"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translation "github.com/go-playground/validator/v10/translations/en"
)

var (
	validate = validator.New()
	transl   ut.Translator
)

func init() {
	if val, ok := binding.Validator.Engine().(*validator.Validate); ok {
		en := en.New()
		unt := ut.New(en, en)
		transl, _ = unt.GetTranslator("en")
		en_translation.RegisterDefaultTranslations(val, transl)
	}
}

func ValidateUserError(validation_err error) *rest_err.RestErr {
	var jsonErr *json.UnmarshalTypeError
	var jsonValidationError validator.ValidationErrors

	if errors.As(validation_err, &jsonErr) {
		return rest_err.NewBadRequestError("Invalid field type")
	} else if errors.As(validation_err, &jsonValidationError) {
		errorsCauses := []rest_err.Causes{}

		for _, err := range validation_err.(validator.ValidationErrors) {
			cause := rest_err.Causes{
				Field:   err.Field(),
				Message: err.Translate(transl),
			}
			errorsCauses = append(errorsCauses, cause)
		}
		return rest_err.NewBadRequestValidationError("Some fields are invalide", errorsCauses)
	} else {
		return rest_err.NewBadRequestError("Error trying to convert fields")
	}
}

// log.Println("Init CreateUser controler")
// var userRequest request.UserRequest

// if err := c.shou
