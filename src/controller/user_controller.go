package controller

import (
	"github.com/fu-hui/am/src/dao"
	"github.com/fu-hui/am/src/model"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

// create user
func CreateUser(context *gin.Context) {
	// get req param
	var user model.User
	if err := context.ShouldBindJSON(&user); err != nil {
		context.JSON(http.StatusBadRequest, model.HttpResponse{
			Code: model.ReqParamError,
			Msg:  "parse req json data fail",
		})
		log.Printf("parse req json data fail, err:%v", err)
		return
	}

	// todo: check user req param

	// save user
	if err := dao.CreateUser(user); err != nil {
		context.JSON(http.StatusBadRequest, model.HttpResponse{
			Code: model.MysqlInsertError,
			Msg:  err.Error(),
		})
		log.Printf("create user fail, err:%v", err)
		return
	}

	context.JSON(http.StatusBadRequest, model.HttpResponse{
		Code: model.Ok,
		Msg:  "parse req json data fail",
	})
}

// user auth v1, just judge password equal
func UserAuthV1(context *gin.Context) {
	var reqUser model.User
	if err := context.ShouldBindJSON(&reqUser); err != nil {
		context.JSON(http.StatusBadRequest, model.HttpResponse{
			Code: model.ReqParamError,
			Msg:  "parse req json data fail",
		})
		log.Printf("parse req json data fail, err:%v", err)
		return
	}

	user, err := dao.QueryUserByUsername(reqUser.UserName)
	if err != nil {
		context.JSON(http.StatusBadRequest, model.HttpResponse{
			Code: model.MysqlQueryError,
			Msg:  "query user fail",
		})
		log.Printf("query user fail, err:%v", err)
		return
	}

	if user == nil {
		context.JSON(http.StatusBadRequest, model.HttpResponse{
			Code: model.ReqParamError,
			Msg:  "user not found",
		})
		log.Printf("user not found")
		return
	}

	if reqUser.Password != user.Password {
		context.JSON(http.StatusBadRequest, model.HttpResponse{
			Code: model.ReqParamError,
			Msg:  "user password invalid",
		})
		log.Printf("user password invalid")
		return
	}

	context.JSON(http.StatusOK, model.HttpResponse{
		Code: model.Ok,
		Msg:  "user auth success",
	})
}
