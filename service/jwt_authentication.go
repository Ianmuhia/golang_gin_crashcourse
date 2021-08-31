// package service

// import (
// 	"fmt"
// 	"os"
// 	"time"

// 	"github.com/dgrijalva/jwt-go"
// )

// type JWTService interface {
// 	GenerateToken(email string, isUser bool) string
// 	ValidateToken(token string) (*jwt.Token, error)
// }

// type authCustomClaims struct {
// 	Name string `json:"name"`
// 	User bool   `json:"user"`
// 	jwt.StandardClaims
// }

// type jwtServices struct {
// 	secretKey string
// 	issure    string
// }

// func JWTAuthService() JWTService {
// 	return &jwtServices{
// 		// secretKey: "",
// 		// issure:    "ian",
// 	}
// }

// func getSecretKey() string {
// 	secret := os.Getenv("SECRET")
// 	if secret == "" {
// 		secret = "secret"

// 	}

// 	return secret

// }

// func (Service *jwtServices) GenerateToken(email string, isUser bool) string {
// 	claims := &authCustomClaims{
// 		email,
// 		isUser,
// 		jwt.StandardClaims{
// 			Audience:  "",
// 			ExpiresAt: time.Now().Add(time.Hour * 48).Unix(),
// 			Id:        "",
// 			IssuedAt:  time.Now().Unix(),
// 			Issuer:    Service.issure,
// 			NotBefore: 0,
// 			Subject:   "",
// 		},
// 	}
// 	token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)
// 	t, err := token.SignedString([]byte(Service.secretKey))
// 	if err != nil {
// 		panic(err)
// 	}
// 	return t

// }

// func (Service *jwtServices) validateToken(encodedToken string) (*jwt.Token, error) {

// 	return jwt.Parse(encodedToken, func(token *jwt.Token) (interface{}, error) {
// 		if _, isvalid := token.Method.(*jwt.SigningMethodHMAC); !isvalid {
// 			return nil, fmt.Errorf("invalid token", token.Header["alg"])
// 		}
// 		return []byte(Service.secretKey), nil
// 	})

// }

package service

import (
	"fmt"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

//jwt service
type JWTService interface {
	GenerateToken(email string, isUser bool) string
	ValidateToken(token string) (*jwt.Token, error)
}
type authCustomClaims struct {
	Name string `json:"name"`
	User bool   `json:"user"`
	jwt.StandardClaims
}

type jwtServices struct {
	secretKey string
	issure    string
}

//auth-jwt
func JWTAuthService() JWTService {
	return &jwtServices{
		secretKey: getSecretKey(),
		issure:    "Rohan",
	}
}

func getSecretKey() string {
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		secret = "secret"
	}
	return secret
}

func (service *jwtServices) GenerateToken(email string, isUser bool) string {
	claims := &authCustomClaims{
		email,
		isUser,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 48).Unix(),
			Issuer:    service.issure,
			IssuedAt:  time.Now().Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	//encoded string
	t, err := token.SignedString([]byte(service.secretKey))
	if err != nil {
		panic(err)
	}
	return t
}

func (service *jwtServices) ValidateToken(encodedToken string) (*jwt.Token, error) {
	return jwt.Parse(encodedToken, func(token *jwt.Token) (interface{}, error) {
		if _, isvalid := token.Method.(*jwt.SigningMethodHMAC); !isvalid {
			return nil, fmt.Errorf("Invalid token", token.Header["alg"])

		}
		return []byte(service.secretKey), nil
	})

}
