package validator_ext

import (
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"strings"
)

const MustInSeq = "#"

func mustIn(fl validator.FieldLevel) bool {
	s := fl.Field().String()
	if s == "" {
		return true
	}

	arr := strings.Split(fl.Param(), MustInSeq)
	for _, v := range arr {
		if v == s {
			return true
		}
	}

	return false
}

func registerMustInEn(validate *validator.Validate, trans ut.Translator) {
	validate.RegisterTranslation(MUST_IN, trans, func(ut ut.Translator) error {
		return ut.Add(MUST_IN, "[{0}] must in [{1}]", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T(MUST_IN, strings.ToLower(fe.Field()), fe.Param())

		return t
	})
}

func registerMustInZh(validate *validator.Validate, trans ut.Translator) {
	validate.RegisterTranslation(MUST_IN, trans, func(ut ut.Translator) error {
		return ut.Add(MUST_IN, "【{0}】必须在【{1}】", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T(MUST_IN, strings.ToLower(fe.Field()), fe.Param())

		return t
	})
}
