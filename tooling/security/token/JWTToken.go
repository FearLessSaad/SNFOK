package token

import (
	"fmt"
	"time"

	"github.com/FearLessSaad/SNFOK/controllers/auth/dto"
	"github.com/FearLessSaad/SNFOK/tooling/logger"

	"github.com/golang-jwt/jwt/v5"
)

const (
	accessTokenSecret  = "dlkasjdlkasdj9r2378rhdbjksasfh"
	refreshTokenSecret = "dlkasjdlkasdj9r2378rhdbdalkj539jfh"

	accessTokenExpiryTime  = 2 * 60 * time.Minute
	refreshTokenExpiryTime = 7 * 24 * time.Hour
)

type TokenClaims struct {
	UserID      string `json:"user_id"`
	Designation string `json:"designation,omitempty"`
	jwt.RegisteredClaims
}

// GenerateToken creates a JWT with the given secret and expiry time
func GenerateToken(userID string, designation string, secret string, expiry time.Duration) (string, error) {
	claims := TokenClaims{
		UserID:      userID,
		Designation: designation,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(expiry)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}

// GenerateJWTTokens generates access and refresh token pairs
func GenerateJWTTokens(userID string, designation string) (dto.JWTToken, string) {
	accessToken, err := GenerateToken(userID, designation, accessTokenSecret, accessTokenExpiryTime)
	if err != nil {
		return dto.JWTToken{}, fmt.Sprintf("Failed to sign access token: %v", err)
	}
	refreshToken, err := GenerateToken(userID, "", refreshTokenSecret, refreshTokenExpiryTime)
	if err != nil {
		return dto.JWTToken{}, fmt.Sprintf("Failed to sign refresh token: %v", err)
	}
	return dto.JWTToken{AccessToken: accessToken, RefreshToken: refreshToken}, ""
}

// verifyToken validates a JWT and returns its claims
func verifyToken(tokenString, secret string, claims jwt.Claims) (bool, string, error) {
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method")
		}
		return []byte(secret), nil
	})
	if err != nil || !token.Valid {
		return false, "Your session has been expired!", err
	}
	if exp, err := claims.GetExpirationTime(); err == nil && exp.Before(time.Now()) {
		return false, "Your session has been expired!", nil
	}
	return true, "", nil
}

// VerifyAccessTokenAndGetClaims verifies an access token
func VerifyAccessTokenAndGetClaims(tokenString string) (bool, string, *TokenClaims) {
	claims := &TokenClaims{}
	isValid, msg, err := verifyToken(tokenString, accessTokenSecret, claims)
	if err != nil {
		logger.Log(logger.DEBUG, fmt.Sprintf("Access token verification failed: %v", err))
	}
	return isValid, msg, claims
}

// VerifyRefreshTokenAndGetClaims verifies a refresh token
func VerifyRefreshTokenAndGetClaims(tokenString string) (bool, string, *TokenClaims) {
	claims := &TokenClaims{}
	isValid, msg, err := verifyToken(tokenString, refreshTokenSecret, claims)
	if err != nil {
		logger.Log(logger.DEBUG, fmt.Sprintf("Refresh token verification failed: %v", err))
	}
	return isValid, msg, claims
}
