package api

import (
	// "fmt"
	"goreact/model"
	"goreact/service"
	"net/http"
	"strconv"
	// "strings"

	"github.com/gin-gonic/gin"
)

type userAPI interface{
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

type analisaAPI struct {
	crudService service.CrudService,
	
}