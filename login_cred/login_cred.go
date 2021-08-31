package logincred

type LoginScredentials struct {
	Email    string `form:"email"`
	Password string `form:"password"`
}
