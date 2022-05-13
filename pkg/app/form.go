package app

import (
	"github.com/gin-gonic/gin"
	val "github.com/go-playground/validator/v10"
	"strings"
)

type ValidError struct {
	Key     string
	Message string
}

type ValidErrors []*ValidError

func (v *ValidError) Error() string {
	return v.Message
}

func (v ValidErrors) Error() string {
	return strings.Join(v.Errors(), ",")
}

func (v ValidErrors) Errors() []string {
	var errs []string
	for _, err := range v {
		errs = append(errs, err.Error())
	}

	return errs
}

func BindAndValid(c *gin.Context, v interface{}) (bool, ValidErrors) {
	var errs ValidErrors
	err := c.ShouldBind(v)
	if err != nil {
		verrs, ok := err.(val.ValidationErrors)
		if !ok {
			return false, errs
		}

		for _, verr := range verrs {
			err := verr.(val.FieldError)
			errs = append(errs, &ValidError{
				Key:     err.Namespace(),
				Message: err.Error(),
			})
		}

		return false, errs
	}

	return true, nil
}

