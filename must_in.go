package validator_ext

import (
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"reflect"
	"strconv"
	"strings"
)

const MustInSeq = "#"

func mustIn(fl validator.FieldLevel) bool {
	field := fl.Field()
	kind := field.Kind()
	var value string
	switch kind {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		value = strconv.FormatInt(field.Int(), 10)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		value = strconv.FormatInt(int64(field.Uint()), 10)
	default:
		value = field.String()
	}

	if value == "" {
		return true
	}

	arr := strings.Split(fl.Param(), MustInSeq)
	for _, v := range arr {
		if v == value {
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
