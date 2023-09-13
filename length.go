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
		return int64(utf8.RuneCountInString(field.String())) <= p
	case reflect.Slice, reflect.Map, reflect.Array:
		p := asInt(param)
		return int64(field.Len()) <= p
	}

	return true
}

func minLength(fl validator.FieldLevel) bool {
	field := fl.Field()
	param := fl.Param()

	switch field.Kind() {
	case reflect.String:
		p := asInt(param)
		return int64(utf8.RuneCountInString(field.String())) >= p
	case reflect.Slice, reflect.Map, reflect.Array:
		p := asInt(param)
		return int64(field.Len()) >= p
	}

	return true
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
	validate.RegisterTranslation(MAX_LENGTH, trans, func(ut ut.Translator) error {
		return ut.Add(MAX_LENGTH, "[{0}] max length is {1}", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T(MAX_LENGTH, strings.ToLower(fe.Field()), fe.Param())

		return t
	})
}

func registerMaxLengthZh(validate *validator.Validate, trans ut.Translator) {
	validate.RegisterTranslation(MAX_LENGTH, trans, func(ut ut.Translator) error {
		return ut.Add(MAX_LENGTH, "【{0}】最大长度是{1}", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T(MAX_LENGTH, strings.ToLower(fe.Field()), fe.Param())

		return t
	})
}

func registerMinLengthEn(validate *validator.Validate, trans ut.Translator) {
	validate.RegisterTranslation(MAX_LENGTH, trans, func(ut ut.Translator) error {
		return ut.Add(MAX_LENGTH, "[{0}] min length is {1}", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T(MAX_LENGTH, strings.ToLower(fe.Field()), fe.Param())

		return t
	})
}

func registerMinLengthZh(validate *validator.Validate, trans ut.Translator) {
	validate.RegisterTranslation(MAX_LENGTH, trans, func(ut ut.Translator) error {
		return ut.Add(MAX_LENGTH, "【{0}】最小长度是{1}", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T(MAX_LENGTH, strings.ToLower(fe.Field()), fe.Param())

		return t
	})
}
