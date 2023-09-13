package validator_ext

import (
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"regexp"
	"strings"
)

var mobileRegex = regexp.MustCompile("^\\+?[1-9]?[0-9]{7,14}$")

func isMobile(fl validator.FieldLevel) bool {
	s := fl.Field().String()
	if s == "" {
		return true
	}

	return mobileRegex.MatchString(s)
}

func registerIsMobileEn(validate *validator.Validate, trans ut.Translator) {
	validate.RegisterTranslation(IS_MOBILE, trans, func(ut ut.Translator) error {
		return ut.Add(IS_MOBILE, "[{0}] must be valid mobile", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T(IS_MOBILE, strings.ToLower(fe.Field()))

		return t
	})
}

func registerIsMobileZh(validate *validator.Validate, trans ut.Translator) {
	validate.RegisterTranslation(IS_MOBILE, trans, func(ut ut.Translator) error {
		return ut.Add(IS_MOBILE, "【{0}】必须是有效的手机号", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T(IS_MOBILE, strings.ToLower(fe.Field()))

		return t
	})
}
