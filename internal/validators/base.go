package validators

import (
	"github.com/go-playground/locales/pt_BR"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	ptBRTranslations "github.com/go-playground/validator/v10/translations/pt_BR"
	"reflect"
	"strings"
)

func getValidator() (*validator.Validate, ut.Translator) {

	translator := pt_BR.New()
	uni := ut.New(translator, translator)

	trans, _ := uni.GetTranslator("pt_BR")

	v := validator.New()

	_ = ptBRTranslations.RegisterDefaultTranslations(v, trans)

	v.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})

	_ = v.RegisterTranslation("required", trans, func(ut ut.Translator) error {
		return ut.Add("required", "{0}:Obrigatório", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("required", fe.Field())
		return t
	})

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
				eArray := strings.Split(e.Translate(trans), ":")
				vErrors[eArray[0]] = eArray[1]
			}
			return vErrors
		}
	}

	return nil
}
