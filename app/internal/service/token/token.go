package token

import (
	"github.com/golang-jwt/jwt"
	"github.com/timurzdev/mentorship-test-task/internal/entity"
	"github.com/timurzdev/mentorship-test-task/internal/generated"
	"time"
)

type TokenService struct {
	key []byte
	ttl time.Duration
}

func NewTokenService(key []byte, ttl time.Duration) *TokenService {
	return &TokenService{
		key: key,
		ttl: ttl,
	}
}

func (t *TokenService) GenerateToken(userId int64) string {
	return "not implemented"
}

func (t *TokenService) GenerateTokenForDummy(role generated.UserType) (*entity.Token, error) {
	claims := jwt.MapClaims{
		"role": role,
		"ttl":  t.ttl.Seconds(),
	}

	generateToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(t.key)
	if err != nil {
		return nil, err
	}

	return &entity.Token{Token: generateToken}, nil
}
