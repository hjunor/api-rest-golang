package services

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type jtwServices struct {
	secretkey string
	issuer    string
}

func NewJwtServices() *jtwServices {
	return &jtwServices{
		secretkey: "hash-256",
		issuer:    "api-rest-go",
	}
}

type Clain struct {
	Sum uint `json:"sum"`
	jwt.StandardClaims
}

func (s *jtwServices) GenerateToken(id uint) (string, error) {
	claim := &Clain{
		id,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 2).Unix(),
			Issuer:    s.issuer,
			IssuedAt:  time.Now().Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	t, err := token.SignedString([]byte(s.secretkey))

	if err != nil {
		return "", err
	}
	return t, nil
}

func (s *jtwServices) ValidateToken(token string) bool {
	_, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, isValid := t.Method.(*jwt.SigningMethodHMAC); !isValid {
			return nil, fmt.Errorf("invalid token")
		}
		return []byte(s.secretkey), nil
	})
	return err == nil
}
