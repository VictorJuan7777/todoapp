package api

import (
	"fmt"
	db "todoapp/db/sqlc"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type callUser struct {
	ID       int64
	Username string
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}

func HashPassword(password string) (string, error) {
	hashedpassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("failed to hash password: %w", err)
	}
	return string(hashedpassword), nil
}

func CheckPassword(password string, hashedpassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedpassword), []byte(password))
}

func CallbackUser(originUser db.Users) *callUser {
	return &callUser{
		ID:       originUser.ID,
		Username: originUser.Username,
	}
}
