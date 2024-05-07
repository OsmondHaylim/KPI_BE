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

type UserService interface {
	Login(user model.User_login) (*string, error)
	Register(user model.RegisterInput) error
	Logout(claim *model.Claims) error
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

func (us *userService) Login(user model.User_login) (*string, error){
	if user.Username == "" || user.Password == "" {return nil, errors.New("username or password empty")}
	dbUser, _ := us.userRepo.GetByName(user.Username)
	if dbUser.Username == "" || dbUser.ID == 0 {return nil, errors.New("user with username" + user.Username + " not found")}
	if !middleware.CheckPasswordHash(user.Password, dbUser.Password) {return nil, errors.New("wrong email or password")}

	expirationTime := time.Now().Add(20 * time.Minute)
	claims := &model.Claims{
		Email: dbUser.Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := t.SignedString(model.JwtKey)
	if err != nil {return nil, errors.New("error signing claims")}

	session := model.Session{
		Token:  tokenString,
		Email:  dbUser.Email,
		Expiry: expirationTime,
	}
	_, err = us.sessionRepo.SessionAvailEmail(session.Email)
	if err != nil {
		err = us.sessionRepo.AddSessions(session)
	} else {
		err = us.sessionRepo.UpdateSessions(session)
	}
	if err != nil {return nil, err}
	_, err = model.CheckValidation(tokenString)
	return &tokenString, err
}

func (us *userService) Register(user model.RegisterInput) error{
	if user.Email == "" || user.Password == "" || user.Username == "" {
		return errors.New("register data is empty")
	} else if user.Password != user.Confirm_password {
		return errors.New("password and confirm password doesn't match")
	}
	_, exists := us.userRepo.GetByEmail(user.Email)
	if exists {return errors.New("email already exists")}

	hashedPw, err := middleware.HashPassword(user.Password)
	if err != nil {return err}

	var result model.User = model.User{
		Username: user.Username,
		Email:    user.Email,
		Password: hashedPw,
		Role:     "Member",
	}
	err = us.userRepo.Store(&result)
	if err != nil {return err}
	return nil
}

func (us *userService) Logout(claim *model.Claims) error{
	expirationTime := time.Now()
	claim.StandardClaims.ExpiresAt = expirationTime.Unix()
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	tokenString, err := t.SignedString(model.JwtKey)
	if err != nil {return errors.New("error signing claims")}

	session := model.Session{
		Token:  tokenString,
		Email:  claim.Email,
		Expiry: expirationTime,
	}
	_, err = us.sessionRepo.SessionAvailEmail(session.Email)
	if err != nil {
		return err
	} else {
		err = us.sessionRepo.DeleteSession(tokenString)
	}
	if err != nil {return err}
	return nil
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
			u.JSON(http.StatusInternalServerError, err)
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