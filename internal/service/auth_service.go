package service

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/olenka--91/reminder-app/internal/domain"
	"github.com/olenka--91/reminder-app/internal/repository"
)

const (
	salt       = "jdhfiaw;em"
	signingKey = "dksfiewi29uejic"
	tokenTTL   = 24 * time.Hour
)

type tokenClaims struct {
	jwt.StandardClaims
	UserID int `json:"user_id"`
}

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(r repository.Authorization) *AuthService {
	return &AuthService{repo: r}
}

func (r *AuthService) CreateUser(user domain.User) (int, error) {
	user.Password = generatePasswordHash(user.Password)
	return r.repo.CreateUser(user)
}

func (r *AuthService) GenerateToken(username, password string) (string, error) {
	user, err := r.repo.GetUser(username, generatePasswordHash(password))
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		user.ID,
	})

	return token.SignedString([]byte(signingKey))

}

func (r *AuthService) ParseToken(accessToken string) (int, error) {
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}
		return []byte(signingKey), nil
	})
	if err != nil {
		return 0, err
	}

	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return 0, errors.New("bad claims format")
	}
	return claims.UserID, nil
}

func generatePasswordHash(pass string) string {
	h := sha1.New()
	h.Write([]byte(pass))

	return fmt.Sprintf("%x", h.Sum([]byte(salt)))
}
