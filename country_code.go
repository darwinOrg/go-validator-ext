package validator_ext

import (
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"strings"
)

var extCountryCodes = map[string]bool{
	"BONAIRE_SINT_EUSTATIUS_AND_SABA": true,
	"BOSNIA_AND_HERZEGOVINA":          true,
	"GEORGIA":                         true,
	"GUINEA":                          true,
	"HOLY_SEE__THE_":                  true,
	"JERSEY":                          true,
	"JORDAN":                          true,
	"SINT_MAARTEN__DUTCH_PART_":       true,
	"SVALBARD_AND_JAN_MAYEN":          true,
	"SYRIAN_ARAB_REPUBLIC__THE_":      true,
	"UNITED_STATES_MINOR_OUTLYING_ISLANDS__THE_": true,
	"WALLIS_AND_FUTUNA":                          true,
}

func extCountryCode(fl validator.FieldLevel) bool {
	s := fl.Field().String()
	if s == "" {
		return true
	}

	return extCountryCodes[s]
}

func registerIsCountryEn(validate *validator.Validate, trans ut.Translator) {
	err := validate.RegisterTranslation(IsCountry, trans, func(ut ut.Translator) error {
		return ut.Add(IsCountry, "[{0}] must be valid country code", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T(IsCountry, strings.ToLower(fe.Field()))

		return t
	})
	if err != nil {
		panic(err)
	}
}

func registerIsCountryZh(validate *validator.Validate, trans ut.Translator) {
	err := validate.RegisterTranslation(IsCountry, trans, func(ut ut.Translator) error {
		return ut.Add(IsCountry, "【{0}】必须是有效的国家编码", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T(IsCountry, strings.ToLower(fe.Field()))

		return t
	})
	if err != nil {
		panic(err)
	}
}
