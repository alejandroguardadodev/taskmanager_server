package validation

import (
	"reflect"

	"github.com/go-playground/validator/v10"
)

type ErrField struct {
	FieldName  string `json:"field_name"`
	ErrorTitle string `json:"error_title"`
	Value      string `json:"value"`
}

var Validate *validator.Validate

func registerCustomValidations() {

}

func ValidationInit() {
	Validate = validator.New(validator.WithRequiredStructEnabled())

	registerCustomValidations()
}

func getJsonTag(fieldname string, val reflect.Value) string {
	for i := 0; i < val.Type().NumField(); i++ {
		f := val.Type().Field(i)

		if f.Name == fieldname {
			return f.Tag.Get("json")
		}
	}

	return fieldname
}

func GetValidateInformation(err error, element any) map[string]string {

	fields := map[string]string{}
	elemtReflect := reflect.ValueOf(element)

	for _, err := range err.(validator.ValidationErrors) {
		fields[getJsonTag(err.Field(), elemtReflect)] = err.ActualTag()
	}

	return fields
}
