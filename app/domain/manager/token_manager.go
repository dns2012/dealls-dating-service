package manager

import (
	"context"
	"github.com/dns2012/dealls-dating-service/app/domain/entity"
	"github.com/dns2012/dealls-dating-service/app/domain/exception"
	"github.com/golang-jwt/jwt/v5"
	"google.golang.org/grpc/metadata"
	"strings"
	"time"
)

type TokenManager interface {
	GetToken(ctx context.Context) (string, error)
	CreateToken(ctx context.Context, user *entity.User) (string, error)
	VerifyToken(ctx context.Context) (*UserClaims, error)
}
type tokenManagerImplementation struct {
	secretKey []byte
}

type UserClaims struct {
	ID			uint   `json:"id"`
	Nickname    string   `json:"nickname"`
	Email 		string `json:"email"`
	Role uint `json:"role"`
	jwt.RegisteredClaims
}

func (m *tokenManagerImplementation) GetToken(ctx context.Context) (tokenString string, err error) {
	meta, ok := metadata.FromIncomingContext(ctx)

	if !ok {
		return tokenString, exception.Unauthenticated("Authorization header is invalid. Please set authorization header with bearer token format.")
	}

	authorizationHeader := meta.Get("authorization")

	if len(authorizationHeader) == 0 {
		return tokenString, exception.Unauthenticated("Authorization header is invalid. Please set authorization header with bearer token format.")
	}

	bearerToken := authorizationHeader[0]

	if !strings.HasPrefix(bearerToken, "Bearer ") {
		return tokenString, exception.Unauthenticated("Authorization header is invalid. Please set authorization header with bearer token format.")
	}

	tokenString = strings.ReplaceAll(bearerToken, "Bearer ", "")

	return tokenString, nil
}

func (m *tokenManagerImplementation) CreateToken(ctx context.Context, user *entity.User) (string, error) {
	accessTokenClaim := &UserClaims{
		ID:	user.ID,
		Nickname: user.Nickname,
		Email: user.Email,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer: "dealls",
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, accessTokenClaim)

	tokenString, err := token.SignedString(m.secretKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func (m *tokenManagerImplementation) VerifyToken(ctx context.Context) (*UserClaims, error) {
	tokenString, err := m.GetToken(ctx)
	if err != nil {
		return nil, err
	}

	userClaims := new(UserClaims)

	token, err := jwt.ParseWithClaims(tokenString, userClaims, func(token *jwt.Token) (interface{}, error) {
		return m.secretKey, nil
	})
	if err != nil {
		if strings.Contains(err.Error(), "token is expired") {
			return nil, exception.Unauthenticated("Access token is expired. Please get the new access token.")
		}
		return nil, exception.Unauthenticated("Access token is invalid. Please pass a valid access token.")
	}

	if !token.Valid {
		return nil, exception.Unauthenticated("Access token is invalid. Please pass a valid access token.")
	}

	return userClaims, nil
}

func NewTokenManager(secretKey []byte) TokenManager {
	return &tokenManagerImplementation{secretKey: secretKey}
}