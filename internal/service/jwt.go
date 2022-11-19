package service

import (
	"InnoTaxi-Driver/internal/config"
	"InnoTaxi-Driver/internal/model"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type Claims struct {
	Username string
	jwt.StandardClaims
}

func SignIn(login model.Login, c *gin.Context) error {
	cfg, err := config.GetCfg()
	if err != nil {
		return err
	}
	duration, _ := time.ParseDuration(cfg.JWT_EXPIRATION_TIME)

	expirationTime := time.Now().Add(duration)

	claims := &Claims{
		Username: login.Phone,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(cfg.SECRET))
	if err != nil {
		return err
	}
	http.SetCookie(c.Writer, &http.Cookie{
		Name:    "token",
		Value:   tokenString,
		Expires: expirationTime,
	})
	http.SetCookie(c.Writer, &http.Cookie{
		Name:    "phone",
		Value:   login.Phone,
		Expires: time.Now().Add(time.Hour * 24),
	})

	return nil
}
