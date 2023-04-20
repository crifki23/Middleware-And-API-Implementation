package helpers

import (
	"chapter3-sesi2/pkg/errs"

	"github.com/asaskevich/govalidator"
)

func ValidateStruct(payload interface{}) errs.MessageErr {
	_, err := govalidator.ValidateStruct(payload)
	if err != nil {
		return errs.NewBadRequest(err.Error())
	}
	return nil
}
