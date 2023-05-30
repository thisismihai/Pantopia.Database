package api

import (
	"github.com/go-playground/validator/v10"
	"github.com/mbaxamb3/pantopia/util"
)

var validIndustry validator.Func = func(fl validator.FieldLevel) bool {
	industry := fl.Field().String()
	return util.IsSupportedIndustry(industry)
}
