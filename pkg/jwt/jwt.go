package jwt

import (
	"foxomni/pkg/config"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type Service struct {
	conf config.AuthConfig
}

type JwtCustomClaims struct {
	UserID int    `json:"user_id"`
	UID    string `json:"uuid"`
	jwt.RegisteredClaims
}

type CachedTokens struct {
	AccessUID  string `json:"access"`
	RefreshUID string `json:"refresh"`
}

func NewService(conf config.AuthConfig) *Service {
	return &Service{
		conf: conf,
	}
}

func (svc *Service) GenerateTokenPair(userID int) (
	accessToken, refreshToken string, err error,
) {
	if accessToken, _, err = svc.createToken(userID, svc.conf.AccessExp,
		svc.conf.AccessKey); err != nil {
		return
	}

	if refreshToken, _, err = svc.createToken(userID, svc.conf.RefreshExp,
		svc.conf.RefreshKey); err != nil {
		return
	}

	return
}

func (svc *Service) createToken(userID int, expMin int, secret string) (
	token, uid string, err error,
) {
	exp := time.Now().Add(time.Minute * time.Duration(expMin))
	uid = uuid.New().String()
	claims := &JwtCustomClaims{
		UserID: userID,
		UID:    uid,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(exp),
		},
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	token, err = jwtToken.SignedString([]byte(secret))

	return
}
