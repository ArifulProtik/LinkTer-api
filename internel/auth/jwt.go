package auth

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

func (a *Auth) VerifyToken(token string) (uuid.UUID, error) {
	claims := jwt.RegisteredClaims{}
	decoded, err := jwt.ParseWithClaims(token, &claims, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("False")
		}
		return []byte(a.cfg.Tokens.PrivateKey), nil
	})
	if err != nil {
		return uuid.UUID{}, err
	}
	if !decoded.Valid {
		return uuid.UUID{}, errors.New("Not a Valid Token")
	}
	u, _ := uuid.Parse(claims.Subject)
	return u, nil

}

func (a *Auth) GenAccesstoken(id *uuid.UUID) (string, error) {
	duration, err := time.ParseDuration(a.cfg.Tokens.AcTokenEXP)
	if err != nil {
		return "", err
	}
	expire := time.Now().Add(duration)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		ExpiresAt: jwt.NewNumericDate(expire),
		NotBefore: jwt.NewNumericDate(time.Now()),
		Subject:   id.String(),
	})
	tokenstring, err := token.SignedString([]byte(a.cfg.PrivateKey))
	if err != nil {
		return "", err
	}
	return tokenstring, nil
}
func (a *Auth) Refreshtoken(id *uuid.UUID) (string, error) {
	duration, err := time.ParseDuration(a.cfg.Tokens.RfTokenEXP)
	if err != nil {
		return "", err
	}
	expire := time.Now().Add(duration)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		ExpiresAt: jwt.NewNumericDate(expire),
		NotBefore: jwt.NewNumericDate(time.Now()),
		Subject:   id.String(),
	})
	tokenstring, err := token.SignedString([]byte(a.cfg.PrivateKey))
	if err != nil {
		return "", err
	}
	return tokenstring, nil
}
