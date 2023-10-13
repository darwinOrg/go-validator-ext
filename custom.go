package validator_ext

import (
	"errors"
	"github.com/go-playground/validator/v10"
)

var customAliasMap = map[string]string{
	IS_CURRENCY: "iso4217",
	IS_COUNTRY:  "iso3166_1_alpha2|iso3166_1_alpha3|iso3166_1_alpha_numeric|" + EXT_COUNTRY_CODE,
}

var customValidatorMap = map[string]validator.Func{
	MAX_LENGTH:       maxLength,
	MIN_LENGTH:       minLength,
	IS_DATE:          isDate,
	IS_DATETIME:      isDateTime,
	REGEX:            regex,
	IS_MOBILE:        isMobile,
	MUST_IN:          mustIn,
	EXT_COUNTRY_CODE: extCountryCode,
}

var CustomValidator *validator.Validate

func NewCustomValidator() *validator.Validate {
	if CustomValidator != nil {
		return CustomValidator
	}

	CustomValidator = validator.New()
	CustomValidator.SetTagName("binding")
	AddCustomRules(CustomValidator)

	return CustomValidator
}

func AddCustomRules(validate *validator.Validate) {
	for k, v := range customAliasMap {
		validate.RegisterAlias(k, v)
	}

	for k, v := range customValidatorMap {
		validate.RegisterValidation(k, v)
	}

	RegisterTranslations(validate)
}

func ValidateDefault(obj any) error {
	err := CustomValidator.Struct(obj)
	errMsg := TranslateValidateError(err, "")
	if errMsg != "" {
		return errors.New(errMsg)
	}
	return nil
}
