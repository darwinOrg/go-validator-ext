package validator_ext

import (
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_trans "github.com/go-playground/validator/v10/translations/en"
	zh_trans "github.com/go-playground/validator/v10/translations/zh"
	"strings"
)

var uniTranslator *ut.UniversalTranslator

func RegisterTranslations(validate *validator.Validate) {
	eg := en.New()
	uniTranslator = ut.New(eg, eg, zh.New())

	transEn, _ := uniTranslator.GetTranslator("en")
	en_trans.RegisterDefaultTranslations(validate, transEn)

	transZh, _ := uniTranslator.GetTranslator("zh")
	zh_trans.RegisterDefaultTranslations(validate, transZh)

	registerIsCountryEn(validate, transEn)
	registerIsDateEn(validate, transEn)
	registerIsDateTimeEn(validate, transEn)
	registerMaxLengthEn(validate, transEn)
	registerMinLengthEn(validate, transEn)
	registerIsMobileEn(validate, transEn)
	registerMustInEn(validate, transEn)
	registerRegexEn(validate, transEn)

	registerIsCountryZh(validate, transZh)
	registerIsDateZh(validate, transZh)
	registerIsDateTimeZh(validate, transZh)
	registerMaxLengthZh(validate, transZh)
	registerMinLengthZh(validate, transZh)
	registerIsMobileZh(validate, transZh)
	registerMustInZh(validate, transZh)
	registerRegexZh(validate, transZh)
}

func TranslateError(err error, locale string) map[string]string {
	if err == nil {
		return nil
	}

	errs, ok := err.(validator.ValidationErrors)
	if !ok {
		return nil
	}

	if locale == "" {
		locale = "zh"
	} else if strings.Contains(locale, "-") {
		locale = strings.Split(locale, "-")[0]
	}

	trans, found := uniTranslator.GetTranslator(locale)
	if !found {
		locale = "zh"
		trans, _ = uniTranslator.GetTranslator(locale)
	}

	return errs.Translate(trans)
}

func TranslateValidateError(err error, lng string) string {
	if err == nil {
		return ""
	}

	transErrors := TranslateError(err, lng)
	if len(transErrors) == 0 {
		return err.Error()
	}

	errMsgs := make([]string, 0, len(transErrors))
	for _, msg := range transErrors {
		errMsgs = append(errMsgs, msg)
	}

	return strings.Join(errMsgs, "\n")
}
