package api

import (
	// "fmt"
	"goreact/model"
	"goreact/service"
	"net/http"
	// "strconv"

	"strings"

	"github.com/gin-gonic/gin"
)

type UserAPI interface{
	Login(u *gin.Context)
	Register(u *gin.Context)
	Logout(u *gin.Context)
	AddUser(u *gin.Context)
	UpdateUser(u *gin.Context)
	DeleteUser(u *gin.Context)
	ChangePassword(u *gin.Context)
	GetUserByID(u *gin.Context)
	GetUserList(u *gin.Context)
	// GetPrivileged(u *gin.Context)
	Profile(u *gin.Context)
}

type userAPI struct {
	crudService service.CrudService
	userService service.UserService
}

func (ua *userAPI) Login (u *gin.Context) {
	var user model.User_login
	err := u.BindJSON(&user)
	if model.ErrorCheck(u, err){return}
	tokenString, err := ua.userService.Login(user)
	if model.ErrorCheck(u, err){return}
	u.JSON(http.StatusFound, gin.H{
		"data":"Successfully Logging In",
		"token":&tokenString,
	})
}

func (ua *userAPI) Register (u *gin.Context) {
	var user model.RegisterInput
	err := u.BindJSON(&user)
	if model.ErrorCheck(u, err){return}
	err = ua.userService.Register(user)
	if model.ErrorCheck(u, err){return}
	u.JSON(http.StatusFound, model.SuccessResponse{Message: "Register User success"})
}

func (ua *userAPI) Logout (u *gin.Context) {
	res := strings.Split(u.GetHeader("Authorization"), " ")
	if len(res) != 2 || res[0] != "Bearer" {
		u.JSON(http.StatusUnauthorized, model.ErrorResponse{Error: "coba lagi bang"})
		return
	}
	claims, err := model.CheckValidation(res[1])
	u.Header("Authorization", "")
	err = ua.userService.Logout(claims)
	if model.ErrorCheck(u, err){return}
	u.JSON(http.StatusFound, model.SuccessResponse{Message: "Logout success"})
}