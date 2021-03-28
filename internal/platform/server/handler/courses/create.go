package courses

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type createRequest struct {
	ID	string	`json:"id" binding:"required"`
	Name string `json:"name" binding:"required"`
	Duration string `json:"duration" binding:"required"`
}

func CreateHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req createRequest
		if err := ctx.BindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, err.Error())
		}

		ctx.Status(http.StatusCreated)
	}
}
