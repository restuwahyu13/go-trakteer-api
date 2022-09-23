package packages

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type UsersToken struct {
	AccessToken         string `json:"access_token"`
	RefreshToken        string `json:"refresh_token"`
	AccessTokenExpired  string `json:"access_token_expired"`
	RefreshTokenExpired string `json:"refresh_token_expired"`
}

func SignToken(payload map[string]interface{}, expiredAt time.Duration) string {
	options := jwt.MapClaims{}

	options["jti"] = uuid.NewString()
	options["exp"] = time.Now().Add(time.Duration(expiredAt)).Unix()
	options["aud"] = "go-trakteer"

	for i, v := range payload {
		options[i] = v
	}

	jwtClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, options)
	token, err := jwtClaims.SignedString([]byte(viper.GetString("JWT_SECRET")))

	if err != nil {
		logrus.Errorf("Generate access token error: %v", err)
	}

	return token
}

func VerifyToken(tokenString string) (string, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(viper.GetString("JWT_SECRET")), nil
	})
	return token.Raw, err
}

func ParseToken(tokenString string) (jwt.Claims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(viper.GetString("JWT_SECRET")), nil
	})
	return token.Claims, err
}
