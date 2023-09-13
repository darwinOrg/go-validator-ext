package validator_ext

import (
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"regexp"
	"strings"
)

func regex(fl validator.FieldLevel) bool {
	s := fl.Field().String()
	if s == "" {
		return true
	}

	return regexp.MustCompile(fl.Param()).MatchString(s)
}

func registerRegexEn(validate *validator.Validate, trans ut.Translator) {
	validate.RegisterTranslation(REGEX, trans, func(ut ut.Translator) error {
		return ut.Add(REGEX, "[{0}] must match {1}", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T(REGEX, strings.ToLower(fe.Field()), fe.Param())

		return t
	})
}

func registerRegexZh(validate *validator.Validate, trans ut.Translator) {
	validate.RegisterTranslation(REGEX, trans, func(ut ut.Translator) error {
		return ut.Add(REGEX, "【{0}】必须匹配 {1}", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T(REGEX, strings.ToLower(fe.Field()), fe.Param())

		return t
	})
}
