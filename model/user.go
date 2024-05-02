package model

import (
	"time"
	"github.com/golang-jwt/jwt"
)

type User struct {
	ID 			int				`gorm:"primaryKey;autoIncrement" json:"id"`
	Username	string			`gorm:"notNull" json:"username"`	
	Email		string			`gorm:"notNull" json:"email"`
	Password	string			`gorm:"notNull" json:"password"`
	Role		string			`gorm:"notNull" json:"role"`
}

type Session struct {
	ID     int       			`gorm:"primaryKey;autoIncrement" json:"id"`
	Token  string    			`json:"token"`
	Email  string    			`json:"email"`
	Expiry time.Time 			`json:"expiry"`
}

var JwtKey = []byte("secret-key")

type Claims struct {
	Email string `json:"email"`
	jwt.StandardClaims
}