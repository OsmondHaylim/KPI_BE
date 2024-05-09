package service

import (
	"goreact/middleware"
	"goreact/model"

	"net/http"
	"time"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

// type UserService interface {
// 	Login(user model.User_login) (*string, error)
// 	Register(user model.RegisterInput) error
// 	Logout(claim *model.Claims) error
// 	AddUser(u *gin.Context)
// 	UpdateUser(u *gin.Context)
// 	DeleteUser(u *gin.Context)
// 	ChangePassword(u *gin.Context)
// 	GetUserByID(u *gin.Context)
// 	GetUserList(u *gin.Context)
// 	// GetPrivileged(u *gin.Context)
// 	Profile(u *gin.Context)
// }

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

func (ua *userService) Profile(u *gin.Context) {
	email := u.Keys["email"].(string)
	compare, boo := ua.userRepo.GetByEmail(email)
	if !boo {
		u.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Trouble finding user"})
	}
	userResult := compare.ToResponse()
	u.JSON(http.StatusOK, userResult)
}