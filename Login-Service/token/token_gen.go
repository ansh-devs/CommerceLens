package token

import (
	"fmt"
	"github.com/aead/chacha20poly1305"
	"github.com/o1egl/paseto"
	_ "github.com/o1egl/paseto"
	"time"
)

type TokenMaker struct {
	paseto       *paseto.V2
	symmetricKey []byte
}

func newTokenMaker(key string) (*TokenMaker, error) {
	if len(key) != chacha20poly1305.KeySize {
		return nil, fmt.Errorf("[ERROR]: invalid key size provided, must exactly be exactly of size %d\n", chacha20poly1305.KeySize)
	}
	return &TokenMaker{
		paseto:       paseto.NewV2(),
		symmetricKey: []byte(key),
	}, nil
}

func (tokenmaker *TokenMaker) createTokenGen(userId string, duration time.Duration) (string, error) {
	payload, err := NewPayload(userId, duration)
	if err != nil {
		return "", fmt.Errorf("[ERROR]: can't create payload : %s\n", err)
	}
	return tokenmaker.paseto.Encrypt(tokenmaker.symmetricKey, payload, nil)
}
func (t *TokenMaker) decryptTokenGen(token string) (*Payload, error) {
	payload := &Payload{}
	if err := t.paseto.Decrypt(token, t.symmetricKey, &payload, nil); err != nil {
		return nil, fmt.Errorf("[ERROR]: can't decrypt token : %s\n", err)
	}
	if err := payload.IsValid(); err != nil {
		return nil, err
	}
	return payload, nil
}
