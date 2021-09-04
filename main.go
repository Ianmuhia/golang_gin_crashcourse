package main

import (
	"gin_jwt/contollers"
	"gin_jwt/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	 loginService := service.StaticLoginService()
	jwtService := service.JWTAuthService()
	loginController := contollers.LoginHandler(jwtService, loginService)

	server := gin.Default()

	server.POST("/login", func(context *gin.Context) {
		token := loginController.Login(context)
		if token != "" {
			context.JSON(http.StatusOK, gin.H{
				"token": token,
			})
		} else {
			context.JSON(http.StatusUnauthorized, nil)
		}
	})

	port :=":8080"
	host:="192.168.0.103"
	err := server.Run( host+port)
	if err != nil {
	panic(err)
	}
}


