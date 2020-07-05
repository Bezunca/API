package validators

import (
	"github.com/go-playground/locales"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"strconv"
)

type Translations struct {
	tag             string
	customRegisFunc validator.RegisterTranslationsFunc
	customTransFunc validator.TranslationFunc
}

func RegisterCustomTranslations(v *validator.Validate, trans ut.Translator) {

	translations := []Translations{
		{
			tag: "required",
			customRegisFunc: func(ut ut.Translator) error {
				return ut.Add("required", "Campo obrigatório", true)
			},
			customTransFunc: func(ut ut.Translator, fe validator.FieldError) string {
				t, _ := ut.T("required", fe.Field())
				return t
			},
		},
		{
			tag: "email",
			customRegisFunc: func(ut ut.Translator) error {
				return ut.Add("email", "Email inválido", true)
			},
			customTransFunc: func(ut ut.Translator, fe validator.FieldError) string {
				t, _ := ut.T("email", fe.Field())
				return t
			},
		},
		{
			tag: "min",
			customRegisFunc: func(ut ut.Translator) error {
				if err := ut.AddCardinal("min", "Campo deve conter no mínimo {0} caracter", locales.PluralRuleOne, false); err != nil {
					return err
				}
				return ut.AddCardinal("min", "Campo deve conter no mínimo {0} caracteres", locales.PluralRuleOther, false)
			},
			customTransFunc: func(ut ut.Translator, fe validator.FieldError) string {
				number, _ := strconv.ParseFloat(fe.Param(), 64)
				t, _ := ut.C("min", number, 0, ut.FmtNumber(number, 0))
				return t
			},
		},
		{
			tag: "max",
			customRegisFunc: func(ut ut.Translator) error {
				if err := ut.AddCardinal("max", "Campo deve conter no máximo {0} caracter", locales.PluralRuleOne, false); err != nil {
					return err
				}
				return ut.AddCardinal("max", "Campo deve conter no máximo {0} caracteres", locales.PluralRuleOther, false)
			},
			customTransFunc: func(ut ut.Translator, fe validator.FieldError) string {
				number, _ := strconv.ParseFloat(fe.Param(), 64)
				t, _ := ut.C("max", number, 0, ut.FmtNumber(number, 0))
				return t
			},
		},

	}

	for _, t := range translations {
		_ = v.RegisterTranslation(t.tag, trans, t.customRegisFunc, t.customTransFunc)
	}
}