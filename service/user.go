package service

import (
	"goreact/repository"
	"goreact/middleware"
	"goreact/model"

	"net/http"
	"strconv"
	"time"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

type UserRepo interface {
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

type userService struct {
	userRepo    repository.UserRepo
	sessionRepo repository.SessionRepo
}

func NewUserAPI(userRepo repository.UserRepo, sessionRepo repository.SessionRepo) *userService {
	return &userService{userRepo, sessionRepo}
}

func (ua *userService) Login(u *gin.Context) {
	var user model.User_login
	if err := u.BindJSON(&user); err != nil {
		u.JSON(http.StatusBadRequest, errors.New("invalid decode json"))
		return
	}
	if user.Username == "" || user.Password == "" {
		u.JSON(http.StatusBadRequest, errors.New("login data is empty"))
		return
	}
	dbUser, _ := ua.userRepo.GetByName(user.Username)
	if dbUser.Username == "" || dbUser.ID == 0 {
		u.JSON(http.StatusBadRequest, errors.New("user not found"))
		return
	}
	if !middleware.CheckPasswordHash(user.Password, dbUser.Password) {
		u.JSON(http.StatusBadRequest, errors.New("wrong email or password"))
		return
	}
	expirationTime := time.Now().Add(20 * time.Minute)
	claims := &model.Claims{
		Email: dbUser.Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := t.SignedString(model.JwtKey)
	if err != nil {
		u.JSON(http.StatusInternalServerError, errors.New("error signing claims"))
		return
	}
	session := model.Session{
		Token:  tokenString,
		Email:  dbUser.Email,
		Expiry: expirationTime,
	}
	_, err = ua.sessionRepo.SessionAvailEmail(session.Email)
	if err != nil {
		err = ua.sessionRepo.AddSessions(session)
	} else {
		err = ua.sessionRepo.UpdateSessions(session)
	}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return model.JwtKey, nil
	})

	if err != nil {
		u.JSON(http.StatusInternalServerError, errors.New("error internal server"))
		return
	}

	if !token.Valid {
		u.JSON(http.StatusUnauthorized, errors.New("invalid token"))
		return
	}

	u.JSON(http.StatusOK, gin.H{
		"message": "login success",
		"data": gin.H{
			"apiKey": tokenString,
			"user": gin.H{
				"id":       dbUser.ID,
				"username": dbUser.Username,
				"email":    dbUser.Email,
				"role":     dbUser.Role,
			},
		},
	})
}

func (ua *userService) Register(u *gin.Context) {
	var user model.RegisterInput
	if err := u.BindJSON(&user); err != nil {
		u.JSON(http.StatusBadRequest, errors.New("invalid decode json"))
		return
	}
	if user.Email == "" || user.Password == "" || user.Username == "" {
		u.JSON(http.StatusBadRequest, errors.New("register data is empty"))
		return
	} else if user.Password != user.Confirm_password {
		u.JSON(http.StatusBadRequest, errors.New("password and confirm password doesn't match"))
		return
	}
	_, exists := ua.userRepo.GetByEmail(user.Email)
	if exists {
		u.JSON(http.StatusBadRequest, errors.New("email already exists"))
		return
	}

	hashedPw, err := middleware.HashPassword(user.Password)
	if err != nil {
		u.JSON(http.StatusInternalServerError, errors.New("Skill Issue at Hashing"))
	}

	var result model.User = model.User{
		Username: user.Username,
		Email:    user.Email,
		Password: hashedPw,
		Role:     "member",
	}
	err = ua.userRepo.Store(&result)
	if err != nil {
		u.JSON(http.StatusInternalServerError, errors.New("Error Storing Data"))
		return
	}
	u.JSON(http.StatusCreated, model.NewSuccessResponse("register success"))
}

func (ua *userService) Logout(u *gin.Context) {
	u.JSON(http.StatusOK, model.NewSuccessResponse("logout success"))
}

func (ua *userService) AddUser(u *gin.Context) {
	var newUser model.User
	if err := u.ShouldBindJSON(&newUser); err != nil {
		u.JSON(http.StatusBadRequest, model.ErrorResponse{Error: err.Error()})
		return
	}

	if newUser.Email == "" || newUser.Password == "" || newUser.Username == "" {
		u.JSON(http.StatusBadRequest, errors.New("register data is empty"))
		return
	}

	err := ua.userRepo.Store(&newUser)
	if err != nil {
		u.JSON(http.StatusInternalServerError, model.ErrorResponse{Error: err.Error()})
		return
	}

	u.JSON(http.StatusOK, model.SuccessResponse{Message: "add User success"})
}

func (ua *userService) UpdateUser(u *gin.Context) {
	var newUser model.User
	if err := u.ShouldBindJSON(&newUser); err != nil {
		u.JSON(http.StatusBadRequest, model.ErrorResponse{Error: err.Error()})
		return
	}
	UserID, err := strconv.Atoi(u.Param("id"))
	if err != nil {
		u.JSON(http.StatusBadRequest, model.ErrorResponse{Error: "invalid User ID"})
		return
	}
	err = ua.userRepo.Update(UserID, newUser)
	if err != nil {
		u.JSON(http.StatusInternalServerError, model.ErrorResponse{Error: err.Error()})
		return
	}
	u.JSON(http.StatusOK, model.SuccessResponse{Message: "User update success"})
}

func (ua *userService) DeleteUser(u *gin.Context) {
	UserID, err := strconv.Atoi(u.Param("id"))
	if err != nil {
		u.JSON(http.StatusBadRequest, model.ErrorResponse{Error: "invalid User ID"})
		return
	}
	err = ua.userRepo.Delete(UserID)
	if err != nil {
		u.JSON(http.StatusInternalServerError, model.ErrorResponse{Error: err.Error()})
		return
	}
	u.JSON(http.StatusOK, model.SuccessResponse{Message: "User delete success"})
}

func (ua *userService) ChangePassword(u *gin.Context) {
	email := u.Keys["email"].(string)

	compare, boo := ua.userRepo.GetByEmail(email)
	if !boo {
		u.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Trouble finding user"})
	}

	curr := u.Param("current_password")
	newp := u.Param("new_password")
	if !middleware.CheckPasswordHash(curr, compare.Email) {
		u.JSON(http.StatusInternalServerError, errors.New("wrong password"))
		return
	} else {
		hashedPw, err := middleware.HashPassword(newp)
		if err != nil {
			u.JSON(http.StatusInternalServerError, errors.New("Skill Issue at Hashing"))
		}
		compare.Password = hashedPw
		err = ua.userRepo.Update(compare.ID, compare)
		if err != nil {
			u.JSON(http.StatusInternalServerError, model.ErrorResponse{Error: err.Error()})
			return
		}
		u.JSON(http.StatusOK, model.SuccessResponse{Message: "User update success"})
	}
}

func (ua *userService) GetUserByID(u *gin.Context) {
	UserID, err := strconv.Atoi(u.Param("id"))
	if err != nil {
		u.JSON(http.StatusBadRequest, model.ErrorResponse{Error: "Invalid User ID"})
		return
	}

	User, err := ua.userRepo.GetByID(UserID)
	if err != nil {
		u.JSON(http.StatusInternalServerError, model.ErrorResponse{Error: err.Error()})
		return
	}

	var result model.UserResponse
	result.User = User.ToCompact()
	result.Message = "User with ID " + strconv.Itoa(UserID) + " Found"

	u.JSON(http.StatusOK, result)
}

func (ua *userService) GetUserList(u *gin.Context) {
	name := u.Query("name")
	// if name != ""{
	User, err := ua.userRepo.SearchName(name)
	if err != nil {
		u.JSON(http.StatusInternalServerError, model.ErrorResponse{Error: err.Error()})
		return
	}

	var result model.UserArrayResponse
	userResult := []model.User_compact{}

	for _, data := range User {
		userResult = append(userResult, data.ToCompact())
	}

	result.Users = userResult
	result.Message = "Getting All Users From Search Result Success"

	u.JSON(http.StatusOK, result)
	// }else{
	// 	User, err := ua.userRepo.GetList()
	// 	if err != nil {
	// 		u.JSON(http.StatusInternalServerError, model.ErrorResponse{Error: err.Error()})
	// 		return
	// 	}
	// 	var result model.UserArrayResponse
	// 	var userResult []model.User_compact
	// 	for _, data := range User{
	// 		userResult = append(userResult, model.ToCompact(data))
	// 	}
	// 	result.Users = userResult
	// 	result.Message = "Getting All Users Success"
	// 	u.JSON(http.StatusOK, result)
	// }
}

// func (ua *userService) GetPrivileged(u *gin.Context) {
// 	User, err := ua.userRepo.GetPrivileged()
// 	if err != nil {
// 		u.JSON(http.StatusInternalServerError, model.ErrorResponse{Error: err.Error()})
// 		return
// 	}
// 	var result model.UserArrayResponse
// 	var userResult []model.User_compact
// 	for _, data := range User {
// 		userResult = append(userResult, model.ToCompact(data))
// 	}
// 	result.Users = userResult
// 	result.Message = "Getting All Privileged Users Success"
// 	u.JSON(http.StatusOK, result)
// }

func (ua *userService) Profile(u *gin.Context) {
	email := u.Keys["email"].(string)

	compare, boo := ua.userRepo.GetByEmail(email)
	if !boo {
		u.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Trouble finding user"})
	}

	userResult := compare.ToCompact()

	u.JSON(http.StatusOK, gin.H{
		"message": "Get User Profile Success",
		"data":    userResult,
	})
}