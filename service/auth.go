package service

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"net/http"
	"otus-social-network-service_gen_swagger/app_error"
	"otus-social-network-service_gen_swagger/cfg"
	dto "otus-social-network-service_gen_swagger/dto/user"
	"otus-social-network-service_gen_swagger/repository"
	"time"
)

const invalidTokenMessage = "invalid token"

type AuthServiceInstance struct {
	tokenTTL          uint
	signingKey        string
	salt              string
	repositoryManager *repository.RepositoryManager
}

type tokenClaims struct {
	jwt.StandardClaims
	UserId uuid.UUID `json:"user_id"`
}

type AuthService interface {
	CreateUser(user *dto.UserRegisterBody) (uuid.UUID, *app_error.HttpError)
	GenerateToken(email, password string) (string, *app_error.HttpError)
	ParseToken(accessToken string) (uuid.UUID, *app_error.HttpError)
}

func NewAuthService(repositoryManager *repository.RepositoryManager) *AuthServiceInstance {
	return &AuthServiceInstance{
		salt:              cfg.Config().Services.Service.Auth.Salt,
		signingKey:        cfg.Config().Services.Service.Auth.SigningKey,
		repositoryManager: repositoryManager,
		tokenTTL:          cfg.Config().Services.Service.Auth.TokenTTL,
	}
}

func (s *AuthServiceInstance) GenerateToken(email, password string) (string, error) {
	user, err := s.repositoryManager.UserRepositoryInstance().GetUser(email, s.generatePasswordHash(password))
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Second * time.Duration(s.tokenTTL)).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		user.Id,
	})
	tokenString, error := token.SignedString([]byte(s.signingKey))

	if error != nil {
		return "", app_error.NewInternalServerError(error)
	}

	return tokenString, nil
}

func (s *AuthServiceInstance) ParseToken(accessToken string) (uuid.UUID, *app_error.HttpError) {
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return uuid.Nil, app_error.NewHttpError(errors.New("invalid signing method"), invalidTokenMessage, "id", http.StatusBadRequest)
		}

		return []byte(s.signingKey), nil
	})

	if err != nil {
		return uuid.Nil, app_error.NewHttpError(err, invalidTokenMessage, "id", http.StatusBadRequest)
	}

	claims, ok := token.Claims.(*tokenClaims)

	if !ok {
		return uuid.Nil, app_error.NewHttpError(errors.New("token claims are not of type tokenClaims"), invalidTokenMessage, "id", http.StatusBadRequest)
	}

	return claims.UserId, nil
}

func (s *AuthServiceInstance) generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(s.salt)))
}

func (s *AuthServiceInstance) CreateUser(user *dto.UserRegisterBody) (uuid.UUID, error) {
	user.Password = s.generatePasswordHash(user.Password)
	return s.repositoryManager.UserRepositoryInstance().CreateUser(user)
}
