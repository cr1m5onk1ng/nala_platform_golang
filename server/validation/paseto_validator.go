package validation

import (
	"errors"
	"fmt"
	"time"

	"github.com/aead/chacha20poly1305"
	"github.com/google/uuid"
	"github.com/o1egl/paseto"
)

type Payload struct {
	TokenID   uuid.UUID `json:"token_id"`
	UserID    string    `json:"user_id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	IssuedAt  time.Time `json:"issued_at"`
	ExpiredAt time.Time `json:"expired_at"`
}

var (
	ErrInvalidToken = errors.New("token is invalid")
	ErrExpiredToken = errors.New("token has expired")
)

func (payload *Payload) Valid() error {
	if time.Now().After(payload.ExpiredAt) {
		return ErrExpiredToken
	}
	return nil
}

type TokenManager interface {
	CreateToken(username string, email string, duration time.Duration) (string, error)
	VerifyToken(token string) (*Payload, error)
}

type PasetoTokenManager struct {
	paseto       *paseto.V2
	symmetricKey []byte
}

func (tokenManager *PasetoTokenManager) CreateToken(username string, email string, duration time.Duration) (string, error) {
	payload, err := NewPayload(username, email, duration)
	if err != nil {
		return "", err
	}
	return tokenManager.paseto.Encrypt(tokenManager.symmetricKey, payload, nil)
}

func (tokenManager *PasetoTokenManager) VerifyToken(token string) (*Payload, error) {
	payload := &Payload{}

	err := tokenManager.paseto.Decrypt(token, tokenManager.symmetricKey, payload, nil)
	if err != nil {
		return nil, err
	}

	err = payload.Valid()
	if err != nil {
		return nil, err
	}

	return payload, nil
}

func NewPayload(username string, email string, duration time.Duration) (*Payload, error) {
	tokenID, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}
	payload := &Payload{
		TokenID:   tokenID,
		Username:  username,
		IssuedAt:  time.Now(),
		ExpiredAt: time.Now().Add(duration),
	}
	return payload, nil
}

func NewPasetoTokenManager(symmetricKey string) (TokenManager, error) {
	if len(symmetricKey) != chacha20poly1305.KeySize {
		return nil, fmt.Errorf("invalid key size: must be exactly %d characters", chacha20poly1305.KeySize)
	}

	tokenManager := &PasetoTokenManager{
		paseto:       paseto.NewV2(),
		symmetricKey: []byte(symmetricKey),
	}

	return tokenManager, nil
}
