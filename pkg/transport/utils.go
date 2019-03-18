package transport

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

func convertParamToInt(paramName string, ctx *gin.Context) (int, error) {
	// get the id parameter off the context
	sid := ctx.Param(paramName)
	// it comes off as a string, so convert to int
	id, err := strconv.Atoi(sid)
	if err != nil {
		return 0, err
	}

	return id, nil
}
