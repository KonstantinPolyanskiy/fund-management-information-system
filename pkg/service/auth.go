package service

import (
	"crypto/sha1"
	"errors"
	"fmt"
	internal_types "fund-management-information-system/internal_types"
	"fund-management-information-system/pkg/repository"
	"github.com/dgrijalva/jwt-go"
	"strings"
	"time"
)

const (
	salt       = "gdfokihierufghuoiw"
	signingKey = "cvobuoiywetr2345>>v"
)

type AuthService struct {
	repo repository.Authorization
}

type TokenClaims struct {
	jwt.StandardClaims
	UserId int `json:"user_id"`
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateClient(client internal_types.SignUpClient) (int, error) {
	if client.Email != "" {
		if err := checkValidEmail(client.Email); err != nil {
			return 0, err
		}
	}
	client.Password = generatePasswordHash(client.Password)
	return s.repo.CreateClient(client)
}

func (s *AuthService) CreateManagerAccount(managerAccount internal_types.ManagerAccount) (int, error) {
	if managerAccount.Email != "" {
		if err := checkValidEmail(managerAccount.Email); err != nil {
			return 0, err
		}
	}
	managerAccount.Password = generatePasswordHash(managerAccount.Password)
	return s.repo.CreateManagerAccount(managerAccount)
}

func (s *AuthService) GenerateToken(login, password string) (string, error) {
	account, err := s.repo.GetAccount(login, generatePasswordHash(password))
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &TokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(24 * time.Hour).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		account.Id,
	})

	return token.SignedString([]byte(signingKey))
}

func (s *AuthService) ParseToken(accessToken string) (int, error) {
	token, err := jwt.ParseWithClaims(accessToken, &TokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid")
		}

		return []byte(signingKey), nil
	})
	if err != nil {
		return 0, err
	}
	claims, ok := token.Claims.(*TokenClaims)
	if !ok {
		return 0, errors.New("claims is not a type")
	}

	return claims.UserId, nil
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}

func checkValidEmail(email string) error {
	if strings.Contains(email, "@mail.ru") {
		return nil
	}
	return errors.New("невалидная почта")
}
