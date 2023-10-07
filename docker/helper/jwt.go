package helper

import (
	"time"
	"github.com/golang-jwt/jwt/v5"
)

type customclaim struct {
	ID uint `json:"id:`
	Name string `json:"name"`
	jwt.RegisteredClaims
}

func GenerateJWT(id uint, name string) string{
	var loaddata customclaim
	loaddata.ID = id
	loaddata.Name = name
	loaddata.ExpiresAt = jwt.NewNumericDate(time.Now().Add(time.Hour *10))
	
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, loaddata)
	t, _ := token.SignedString([]byte("5347876725"))
	return t
}
