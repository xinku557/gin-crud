package utils

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func CreateToken(userId uint, level uint) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["user_id"] = userId
	claims["level"] = level
	claims["exp"] = time.Now().Add(time.Hour * 24 * 7).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(os.Getenv("JWT_SECRECT")))
}

func TokenValid(ctx *gin.Context) error {
	tokenString := extractToken(ctx)
	_, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {

		_, ok := token.Method.(*jwt.SigningMethodHMAC)

		if !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("JWT_SECRECT")), nil
	})

	if err != nil {
		return err
	}

	return nil
}

func extractToken(ctx *gin.Context) string {

	bearerToken := ctx.Request.Header["Authorization"]
	if len(bearerToken) != 1 {
		return ""
	}
	if len(strings.Split(bearerToken[0], " ")) == 2 {
		return strings.Split(bearerToken[0], " ")[1]
	}
	return ""
}

func ExtractTokenID(ctx *gin.Context) (map[string]uint, error) {

	tokenString := extractToken(ctx)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("JWT_SECRECT")), nil
	})
	if err != nil {
		return map[string]uint{"id": 0, "level": 0}, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		uid, err := strconv.ParseUint(fmt.Sprintf("%.0f", claims["user_id"]), 10, 32)
		level, err := strconv.ParseUint(fmt.Sprintf("%.0f", claims["level"]), 10, 32)
		if err != nil {
			return map[string]uint{"id": 0, "level": 0}, err
		}
		return map[string]uint{"id": uint(uid), "level": uint(level)}, nil
	}
	return map[string]uint{"id": 0, "level": 0}, nil
}
