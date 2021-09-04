package contollers

import (
	logincred "gin_jwt/login_cred"
	"gin_jwt/service"

	"github.com/gin-gonic/gin"
)

type LoginController interface {
	Login(ctx *gin.Context) string
}
type loginController struct {
	loginService service.LoginService
	jwtService   service.JWTService
}
func LoginHandler(loginService service.JWTService, jWtService service.LoginService) LoginController {
	return &loginController{
		loginService: jWtService,
		jwtService:   loginService,
	}
}

func (controller *loginController) Login(ctx *gin.Context) string {
	var credentials logincred.LoginScredentials
	err :=ctx.ShouldBind(&credentials)
	if err != nil {
		return "no data found"
	}
isAuthenticated := controller.loginService.LoginUser(credentials.Email,credentials.Password)
if isAuthenticated{
	return controller.jwtService.GenerateToken(credentials.Email, true)
}
	return ""
}