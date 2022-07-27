package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/luigiescalante/proyect-template/core/domain"
	"github.com/luigiescalante/proyect-template/core/repository/api"
	"github.com/luigiescalante/proyect-template/core/repository/database"
	"github.com/luigiescalante/proyect-template/core/service"
	"net/http"
	"strconv"
)

func GetUsers(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"system": "ping",
	})
}

func GetUserById(c *gin.Context) {
	userId, _ := strconv.Atoi(c.Param("id"))
	userSrv := service.UserServiceFactory(nil, database.UserRepositoryFactory())
	user, err := userSrv.GetById(userId)
	if err != nil {
		api.Error(c, err.Error())
		return
	}
	api.Success(c, user)
}

func Save(c *gin.Context) {
	var user *domain.Users
	err := c.BindJSON(&user)
	if err != nil {
		return
	}
	userSrv := service.UserServiceFactory(user, database.UserRepositoryFactory())
	err = userSrv.Save()
	if err != nil {
		api.Error(c, err.Error())
		return
	}
	api.Success(c, user)
}

func Update(c *gin.Context) {
	api.Success(c, "kaiba")
}

func Delete(c *gin.Context) {
	userId, _ := strconv.Atoi(c.Param("id"))
	userSrv := service.UserServiceFactory(nil, database.UserRepositoryFactory())
	user, err := userSrv.GetById(userId)
	userSrv.User = user
	if err != nil {
		api.Error(c, err.Error())
		return
	}
	err = userSrv.Delete()
	if err != nil {
		return
	}
	api.Success(c, user)
}
