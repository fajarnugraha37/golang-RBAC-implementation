package jwt

import (
	"os"
	"strconv"
	"time"

	"bitbucket.org/nugrahafajar37/fajar-rgb-golang-test/common/util"
	jwtLib "github.com/golang-jwt/jwt/v4"
)

var jwtHelperInstance *JwtHelper

type JwtHelper struct {
	secretKey     []byte
	expiresInHour int
	signingMethod *jwtLib.SigningMethodHMAC
}

func NewJwtNelper() *JwtHelper {
	return util.Singleton(jwtHelperInstance, func() *JwtHelper {
		expiresInHour := 24
		if res, err := strconv.Atoi(os.Getenv("JWT_EXPIRES_IN")); err != nil {
			expiresInHour = res
		}
		return &JwtHelper{secretKey: []byte(os.Getenv("JWT_SECRET")), expiresInHour: expiresInHour, signingMethod: jwtLib.SigningMethodHS256}
	})
}

func (helper *JwtHelper) GenerateToken(userId int, roles ...string) (token string, err error) {
	jwtToken := jwtLib.NewWithClaims(
		jwtLib.SigningMethodHS256,
		Claims{
			RegisteredClaims: jwtLib.RegisteredClaims{
				ID:        strconv.Itoa(userId),
				ExpiresAt: jwtLib.NewNumericDate(helper.getExpiredIn()),
				Issuer:    "personal",
				Subject:   "personal",
			},
			UserID: userId,
			Roles:  roles,
		},
	)
	token, err = jwtToken.SignedString(helper.secretKey)

	return
}

func (helper *JwtHelper) ValidateToken(token string) (*Claims, error) {
	claims := &Claims{}
	_, err := jwtLib.ParseWithClaims(token, claims, helper.keyFunc)

	return claims, err
}

func (helper *JwtHelper) getExpiredIn() time.Time {
	return time.Now().Add(time.Duration(helper.expiresInHour) * time.Hour)
}

func (helper *JwtHelper) keyFunc(token *jwtLib.Token) (interface{}, error) {
	return helper.secretKey, nil
}
