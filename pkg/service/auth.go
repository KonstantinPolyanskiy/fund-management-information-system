package service

import (
	"crypto/sha1"
	"errors"
	"fmt"
	internal_types "fund-management-information-system/internal-types"
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

func (s *AuthService) CreateClient(client internal_types.Client) (int, error) {
	if err := checkValidEmail(client.Email); err != nil {
		return 0, err
	}
	client.Password = generatePasswordHash(client.Password)
	return s.repo.CreateClient(client)
}

func (s *AuthService) CreateManager(manager internal_types.Manager) (int, error) {
	if err := checkValidEmail(manager.Email); err != nil {
		return 0, err
	}
	manager.Password = generatePasswordHash(manager.Password)
	return s.repo.CreateManager(manager)
}

func (s *AuthService) GenerateToken(login, password string) (string, error) {
	user, err := s.repo.User(login, generatePasswordHash(password))
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &TokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(24 * time.Hour).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		user.Id,
	})

	return token.SignedString([]byte(signingKey))
}

func (s *AuthService) ParseToken(token string) (int, error) {
	//TODO implement me
	panic("implement me")
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
	return errors.New("Невалидная почта")
}
