package validator_ext

import (
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"reflect"
	"strconv"
	"strings"
	"unicode/utf8"
)

func maxLength(fl validator.FieldLevel) bool {
	field := fl.Field()
	param := fl.Param()

	switch field.Kind() {
	case reflect.String:
		p := asInt(param)
		length := int64(utf8.RuneCountInString(field.String()))
		return length == 0 || length <= p
	case reflect.Slice, reflect.Map, reflect.Array:
		p := asInt(param)
		length := int64(field.Len())
		return length == 0 || length <= p
	default:
		return true
	}
}

func minLength(fl validator.FieldLevel) bool {
	field := fl.Field()
	param := fl.Param()

	switch field.Kind() {
	case reflect.String:
		p := asInt(param)
		length := int64(utf8.RuneCountInString(field.String()))
		return length == 0 || length >= p
	case reflect.Slice, reflect.Map, reflect.Array:
		p := asInt(param)
		length := int64(field.Len())
		return length == 0 || length >= p
	default:
		return true
	}
}

// asInt returns the parameter as a int64
// or panics if it can't convert
func asInt(param string) int64 {
	i, err := strconv.ParseInt(param, 0, 64)
	panicIf(err)

	return i
}

func panicIf(err error) {
	if err != nil {
		panic(err.Error())
	}
}

func registerMaxLengthEn(validate *validator.Validate, trans ut.Translator) {
	err := validate.RegisterTranslation(MaxLength, trans, func(ut ut.Translator) error {
		return ut.Add(MaxLength, "[{0}] max length is {1}", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T(MaxLength, strings.ToLower(fe.Field()), fe.Param())

		return t
	})
	if err != nil {
		panic(err)
	}
}

func registerMaxLengthZh(validate *validator.Validate, trans ut.Translator) {
	err := validate.RegisterTranslation(MaxLength, trans, func(ut ut.Translator) error {
		return ut.Add(MaxLength, "【{0}】最大长度是{1}", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T(MaxLength, strings.ToLower(fe.Field()), fe.Param())

		return t
	})
	if err != nil {
		panic(err)
	}
}

func registerMinLengthEn(validate *validator.Validate, trans ut.Translator) {
	err := validate.RegisterTranslation(MinLength, trans, func(ut ut.Translator) error {
		return ut.Add(MinLength, "[{0}] min length is {1}", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T(MinLength, strings.ToLower(fe.Field()), fe.Param())

		return t
	})
	if err != nil {
		panic(err)
	}
}

func registerMinLengthZh(validate *validator.Validate, trans ut.Translator) {
	err := validate.RegisterTranslation(MinLength, trans, func(ut ut.Translator) error {
		return ut.Add(MinLength, "【{0}】最小长度是{1}", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T(MinLength, strings.ToLower(fe.Field()), fe.Param())

		return t
	})
	if err != nil {
		panic(err)
	}
}
