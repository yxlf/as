package validator

import (
	"as/internal/service"
	"context"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

var UniqueStoreName validator.Func = func(fl validator.FieldLevel) bool {
	if k, ok := fl.Field().Interface().(string); ok {
		s := service.New(context.Background())
		s.CheckStoreNameIsUnique(k)
		return true
	}
	return false
}

func init() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		err := v.RegisterValidation("unique_store_name", UniqueStoreName)
		if err != nil {
			return
		}
	}
}
