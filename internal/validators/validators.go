package validators

import (
	"regexp"

	validator "github.com/go-playground/validator/v10"
)

func RegisterCustomValidators(v *validator.Validate) {
	v.RegisterValidation("pan", validatePAN)
	v.RegisterValidation("mobile", validateMobile)
}

func validatePAN(fl validator.FieldLevel) bool {
	panRegex := `^[A-Z]{5}[0-9]{4}[A-Z]$` // five letters + four digits + one letter (ex. ABCDE1234F)
	return regexp.MustCompile(panRegex).MatchString(fl.Field().String())
}

func validateMobile(fl validator.FieldLevel) bool {
	mobileRegex := `^[0-9]{10}$` // 10 digit number
	return regexp.MustCompile(mobileRegex).MatchString(fl.Field().String())
}
