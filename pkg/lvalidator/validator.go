package lvalidator

import (
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
)

type lvalidator interface {
	Struct(interface{}) []ErrorMsg
}

type ErrorMsg struct {
	Field   string
	Messege string
}

type Lvalidator struct {
	V  *validator.Validate
	ut *ut.UniversalTranslator
}

func New() *Lvalidator {
	en := en.New()
	return &Lvalidator{
		V:  validator.New(),
		ut: ut.New(en, en),
	}
}

func (l *Lvalidator) Struct(s interface{}) []ErrorMsg {

	msgs := []ErrorMsg{}
	trans, _ := l.ut.GetTranslator("en")
	en_translations.RegisterDefaultTranslations(l.V, trans)
	err := l.V.Struct(s)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			msgs = append(msgs, ErrorMsg{
				Field:   err.StructField(),
				Messege: err.Translate(trans),
			})
		}
		return msgs

	}

	return nil
}
