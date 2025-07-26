package token

import (
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const TenantIDKey = "X-Tenant-ID"

var jwtSecretKey = []byte("My Super Secret Key | BuLu LuLu KuLu MuLu SuLu")

type JwtClaims struct {
	TenantID string `json:"tenant_id"`
	jwt.RegisteredClaims
}

func GenerateToken(tenantID string) (string, error) {
	claims := JwtClaims{
		TenantID: tenantID,
		RegisteredClaims: jwt.RegisteredClaims{
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(jwtSecretKey)
}

func ParseJWT(tokenStr string) (*JwtClaims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &JwtClaims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecretKey, nil
	})
	if err != nil || !token.Valid {
		return nil, err
	}

	claims, ok := token.Claims.(*JwtClaims)
	if !ok {
		return nil, jwt.ErrTokenInvalidClaims
	}

	return claims, nil
}

func GetTenantID(r *http.Request) string {
	return r.Header.Get(TenantIDKey)
}
