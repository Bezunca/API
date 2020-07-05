package validators

import (
	"github.com/go-playground/locales/pt_BR"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	ptBRTranslations "github.com/go-playground/validator/v10/translations/pt_BR"
	"reflect"
)

func getValidator() (*validator.Validate, ut.Translator) {

	translator := pt_BR.New()
	uni := ut.New(translator, translator)

	trans, _ := uni.GetTranslator("pt_BR")

	v := validator.New()

	_ = ptBRTranslations.RegisterDefaultTranslations(v, trans)

	v.RegisterTagNameFunc(func(fld reflect.StructField) string {
		return fld.Tag.Get("json")
	})

	RegisterCustomTranslations(v, trans)

	return v, trans
}

func ValidateStruct(s interface{}) map[string]string {

	v, trans := getValidator()

	err := v.Struct(s)
	if err != nil {
		validationErrors := err.(validator.ValidationErrors)
		if len(validationErrors) > 0 {
			vErrors := map[string]string{}
			for _, e := range validationErrors {
				vErrors[e.Field()] = e.Translate(trans)
			}
			return vErrors
		}
	}

	return nil
}
