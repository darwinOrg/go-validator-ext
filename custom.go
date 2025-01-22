package validator_ext

import (
	"errors"
	"github.com/go-playground/validator/v10"
)

var customAliasMap = map[string]string{
	IsCurrency: "iso4217",
	IsCountry:  "iso3166_1_alpha2|iso3166_1_alpha3|iso3166_1_alpha_numeric|" + ExtCountryCode,
}

var customValidatorMap = map[string]validator.Func{
	MaxLength:      maxLength,
	MinLength:      minLength,
	IsDate:         isDate,
	IsDatetime:     isDateTime,
	Regex:          regex,
	IsMobile:       isMobile,
	MustIn:         mustIn,
	ExtCountryCode: extCountryCode,
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
		err := validate.RegisterValidation(k, v)
		if err != nil {
			panic(err)
		}
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
