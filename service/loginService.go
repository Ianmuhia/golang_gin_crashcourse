package service

type LoginService interface {
	LoginUser(username, password string) bool
}

type LoginInformation struct {
	email    string
	password string
}

func StaticLoginService() LoginService {
	return &LoginInformation{
		email:    "ianmuhia3@gmail.com",
		password: "*#*Johnte2536",
	}

}

func (info *LoginInformation) LoginUser(email, password string) bool {

	return info.email == email && info.password == password

}
