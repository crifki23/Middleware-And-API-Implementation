package helpers

import (
	"chapter3-sesi2/pkg/errs"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetParamId(ctx *gin.Context, key string) (int, errs.MessageErr) {
	value := ctx.Param(key)
	id, err := strconv.Atoi(value)
	if err != nil {
		return 0, errs.NewBadRequest("invalid parameter id")
	}
	return id, nil
}
