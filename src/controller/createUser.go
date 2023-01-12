package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/jjeanjacques10/crud-golang/src/configuration/rest_err"
)

func CreateUser(c *gin.Context) {
	err := rest_err.NewBadRequestError("wrong route")
	c.JSON(err.Code, err)
}
