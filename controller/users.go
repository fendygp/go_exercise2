package controller

import (
	"bitbucket.org/bridce/go_exercise2/usecase"
	"github.com/gin-gonic/gin"
)

//UsersControllerStruct controller to handle users domain
type UsersControllerStruct struct {
	usersUsecase usecase.UsersInterface
}

func CreateUsersControllerImp(router *gin.RouterGroup, usersUsecase usecase.UsersInterface) {
	//define your service path here

}

//CreateUsers function to handle creeate user request
func (g *UsersControllerStruct) CreateUsers(c *gin.Context) {

}
