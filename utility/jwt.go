package utility

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
)

type JWTPayload struct {
	Id        int
	Role      string
	ExpiredAt time.Time
}

func GenerateJWT(payload JWTPayload, secret string, expired int16) (token string, err error) {
	claims := jwt.MapClaims{
		"id":        payload.Id,
		"expiredAt": time.Now().Add(time.Duration(expired) * time.Minute),
	}

	tok := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	token, err = tok.SignedString([]byte(secret))
	if err != nil {
		return
	}
	return
}
func VerifyToken(tokenString string, secret string) (token JWTPayload, err error) {
	tokens, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method")
		}

		return []byte(secret), nil
	})

	if err != nil {
		return token, err
	}

	claims, ok := tokens.Claims.(jwt.MapClaims)
	if ok && tokens.Valid {
		expired, _ := time.Parse(time.RFC3339, fmt.Sprintf("%v", claims["expiredAt"]))
		id := claims["id"].(int)

		token = JWTPayload{
			Id:        id,
			Role:      fmt.Sprintf("%v", claims["role"]),
			ExpiredAt: expired,
		}
		return token, nil
	}

	return token, fmt.Errorf("unable to extract claims")
}
