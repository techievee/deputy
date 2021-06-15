package data

import (
	"fmt"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator"
)

// Roles data model
type Roles struct {
	Id     uint64 `json:"Id" validate:"required,numeric"`
	Name   string `json:"Name" validate:"required"`
	Parent uint64 `json:"Parent" validate:"gte=0,numeric"`
}

// Validate Validate the input during the load using the Validator
func (r *Roles) Validate() []string {

	var validationErrors []string

	v := validator.New()
	english := en.New()
	uni := ut.New(english, english)
	trans, _ := uni.GetTranslator("en")

	if err := v.Struct(r); err != nil {
		errors, _ := err.(validator.ValidationErrors)
		for _, err := range errors {
			translatedErr := fmt.Errorf(err.Translate(trans))
			validationErrors = append(validationErrors, translatedErr.Error())
		}
	}

	return validationErrors

}
