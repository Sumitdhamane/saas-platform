	package auth

import (
	"errors"

	"golang.org/x/crypto/bcrypt"

	"github.com/sumitdhamane/saas-platform/configs"
	"github.com/sumitdhamane/saas-platform/internal/user"
)

func Login(
	email string,
	password string,
	cfg *configs.Config,
) (string, error) {

	u, hashedPassword, err := user.GetUserByEmail(email)
	if err != nil {
		return "", errors.New("invalid credentials")
	}

	err = bcrypt.CompareHashAndPassword(
		[]byte(hashedPassword),
		[]byte(password),
	)

	if err != nil {
		return "", errors.New("invalid credentials")
	}

	token, err := GenerateToken(
		u.ID,
		u.TenantID,
		u.Email,
		cfg.JWTSecret,
	)

	if err != nil {
		return "", err
	}

	return token, nil
}