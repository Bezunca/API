package validators

import (
	"errors"
	"github.com/go-playground/locales/pt_BR"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	ptBRTranslations "github.com/go-playground/validator/v10/translations/pt_BR"
)

func getValidator() (*validator.Validate, ut.Translator, error) {

	translator := pt_BR.New()
	uni := ut.New(translator, translator)

	trans, found := uni.GetTranslator("pt_BR")
	if !found {
		return nil, nil, errors.New("translator not found")
	}

	v := validator.New()

	err := ptBRTranslations.RegisterDefaultTranslations(v, trans)
	if err != nil {
		return nil, nil, err
	}

	return v, trans, nil
}

func ValidateStruct(s interface{}) error {

	v, trans, err := getValidator()
	if err != nil {
		return err
	}

	err = v.Struct(s)
	if err != nil {
		validationErrors := err.(validator.ValidationErrors)
		if len(validationErrors) > 0 {
			return errors.New(validationErrors[0].Translate(trans))
		}
	}

	return nil
}
