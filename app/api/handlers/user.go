package handlers

import (
	"gosplitwise/app/models"
	usermdl "gosplitwise/app/modules/user"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func Register(c *gin.Context) {
	//TODO: add password encrytion
	validate := validator.New()
	var user models.User
	if err := c.Bind(&user); err != nil {
		c.JSON(http.StatusBadRequest, models.Response{
			Error:  err.Error(),
			Status: models.STATUS_FAILED,
			Result: nil,
		})
		return
	}
	if err := validate.Struct(user); err != nil {
		c.JSON(http.StatusBadRequest, models.Response{
			Error:   err.Error(),
			Status:  models.STATUS_FAILED,
			Result:  nil,
			Message: "validation failed",
		})
		return
	}
	user.Id = primitive.NewObjectID()
	user.Friends = []primitive.ObjectID{}
	if err := usermdl.Register(user); err != nil {
		c.JSON(http.StatusInternalServerError, models.Response{
			Error:   err.Error(),
			Status:  models.STATUS_FAILED,
			Result:  nil,
			Message: "error while registering user",
		})
		return
	}
	c.JSON(http.StatusOK, models.Response{
		Error:   "",
		Status:  models.STATUS_SUCCESS,
		Result:  user,
		Message: "registration successfull",
	})
}

func Login(c *gin.Context) {
	validate := validator.New()

	var userCred models.Credential
	if err := c.Bind(&userCred); err != nil {
		c.JSON(http.StatusBadRequest, models.Response{
			Error:   err.Error(),
			Message: "",
			Result:  nil,
			Status:  models.STATUS_FAILED,
		})
		return
	}
	if err := validate.Struct(userCred); err != nil {
		c.JSON(http.StatusBadRequest, models.Response{
			Error:   err.Error(),
			Message: "email and password is required",
			Result:  nil,
			Status:  models.STATUS_FAILED,
		})
		return
	}
	user, err := usermdl.Login(userCred)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			c.JSON(http.StatusNotFound, models.Response{
				Error:   err.Error(),
				Message: "user not found",
				Result:  nil,
				Status:  models.STATUS_FAILED,
			})
		}
		return
	}
	c.JSON(http.StatusOK, models.Response{
		Error:   "",
		Status:  models.STATUS_SUCCESS,
		Result:  user,
		Message: "logged in",
	})
}
