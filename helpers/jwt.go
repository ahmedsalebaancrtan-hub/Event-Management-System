package helpers

import (
	"time"

	"github.com/ahmedsaleban/eventManagementsystem/infra"
	"github.com/ahmedsaleban/eventManagementsystem/models"
	"github.com/golang-jwt/jwt"
)

func GenerateJwt(role models.Role, sub string, ExpireIn int64, isrefreshToken bool) (string, error) {
	config := infra.Configuration

	var jwtsecret []byte

	if isrefreshToken {
		jwtsecret = []byte(config.Refresh_jwt_token)
	} else {
		jwtsecret = []byte(config.Access_jwt_Token)
	}
	claims := jwt.MapClaims{
		"sub":            sub,
		"npf":            time.Now(),
		"exp":            ExpireIn,
		"isrefreshToken": isrefreshToken,
		"role":           role,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtsecret)

}
