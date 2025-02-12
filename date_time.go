package validator_ext

import (
	"fmt"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"reflect"
	"strings"
	"time"
)

var (
	// 指定日期格式
	dateFormat = "2006-01-02"

	// 指定时间格式
	dateTimeFormat = "2006-01-02 15:04:05"
)

// isDate is the validation function for validating if the current field's value is a valid date string.
func isDate(fl validator.FieldLevel) bool {
	return isDateOrDateTime(fl, dateFormat)
}

// isDateTime is the validation function for validating if the current field's value is a valid datetime string.
func isDateTime(fl validator.FieldLevel) bool {
	return isDateOrDateTime(fl, dateTimeFormat)
}

func isDateOrDateTime(fl validator.FieldLevel, format string) bool {
	field := fl.Field()
	s := field.String()
	if s == "" {
		return true
	}

	if field.Kind() == reflect.String {
		_, err := time.Parse(format, field.String())

		return err == nil
	}

	panic(fmt.Sprintf("Bad field type %T", field.Interface()))
}

func registerIsDateEn(validate *validator.Validate, trans ut.Translator) {
	err := validate.RegisterTranslation(IsDate, trans, func(ut ut.Translator) error {
		return ut.Add(IsDate, "[{0}] date format must be yyyy-MM-dd", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T(IsDate, strings.ToLower(fe.Field()))

		return t
	})
	if err != nil {
		panic(err)
	}
}

func registerIsDateZh(validate *validator.Validate, trans ut.Translator) {
	err := validate.RegisterTranslation(IsDate, trans, func(ut ut.Translator) error {
		return ut.Add(IsDate, "【{0}】日期格式必须是 yyyy-MM-dd", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T(IsDate, strings.ToLower(fe.Field()))

		return t
	})
	if err != nil {
		panic(err)
	}
}

func registerIsDateTimeEn(validate *validator.Validate, trans ut.Translator) {
	err := validate.RegisterTranslation(IsDatetime, trans, func(ut ut.Translator) error {
		return ut.Add(IsDatetime, "[{0}] datetime format must be yyyy-MM-dd", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T(IsDatetime, strings.ToLower(fe.Field()))

		return t
	})
	if err != nil {
		panic(err)
	}
}

func registerIsDateTimeZh(validate *validator.Validate, trans ut.Translator) {
	err := validate.RegisterTranslation(IsDatetime, trans, func(ut ut.Translator) error {
		return ut.Add(IsDatetime, "【{0}】日期时间格式必须是 yyyy-MM-dd hh:mm:ss", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T(IsDatetime, strings.ToLower(fe.Field()))

		return t
	})
	if err != nil {
		panic(err)
	}
}
