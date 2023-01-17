package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jjeanjacques10/crud-golang/src/configuration/logger"
	"github.com/jjeanjacques10/crud-golang/src/configuration/validation"
	"github.com/jjeanjacques10/crud-golang/src/controller/model/request"
	"github.com/jjeanjacques10/crud-golang/src/controller/model/response"
	"go.uber.org/zap"
)

func CreateUser(c *gin.Context) {
	logger.Info("Init CreateUser controller",
		zap.String("journey", "createUser"),
	)
	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		logger.Error("Error trying to validate user info", err,
			zap.String("journey", "createUser"))
		restErr := validation.ValidateUserError(err)

		c.JSON(restErr.Code, restErr)
		return
	}

	response := response.UserResponse{
		ID:    "test",
		Email: userRequest.Email,
		Name:  userRequest.Name,
		Age:   userRequest.Age,
	}

	logger.Info("User created sucessfully",
		zap.String("journey", "createUser"))

	c.JSON(http.StatusOK, response)
}
