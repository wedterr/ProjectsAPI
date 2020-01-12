package controller

import (
	"ProjectsAPI/app"
	"net/http"

	"github.com/dgrijalva/jwt-go"
)

//Создать токен JWT
func GetToken(w http.ResponseWriter, r *http.Request) {
	tk := &app.Token{UserId: 1}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	tokenString, _ := token.SignedString([]byte(app.TokenPass))
	w.Write([]byte(tokenString))
}
