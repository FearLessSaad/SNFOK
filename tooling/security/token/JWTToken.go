package token

import (
	"github.com/FearLessSaad/SNFOK/controllers/auth/dto"
	"github.com/FearLessSaad/SNFOK/tooling/logger"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const (
	accessTokenSecret  = "dlkasjdlkasdj9r2378rhdbjksasfh"
	refreshTokenSecret = "dlkasjdlkasdj9r2378rhdbdalkj539jfh"
	signUpTokenSecret  = "dlkasjdlkasdj98rhdbdalkj539jksasfh"

	accessTokenExpiryTime  = 15 * time.Minute
	refreshTokenExpiryTime = 7 * 24 * time.Hour
	signUpTokenExpiryTime  = 15 * time.Minute
)

type TokenClaims struct {
	UserID string   `json:"user_id"`
	Roles  []string `json:"roles,omitempty"`
	jwt.RegisteredClaims
}

// GenerateToken creates a JWT with the given secret and expiry time
func GenerateToken(userID string, roles []string, secret string, expiry time.Duration) (string, error) {
	claims := TokenClaims{
		UserID: userID,
		Roles:  roles,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(expiry)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}

// GenerateJWTTokens generates access and refresh token pairs
func GenerateJWTTokens(userID string, roles []string) (dto.JWTToken, string) {
	accessToken, err := GenerateToken(userID, roles, accessTokenSecret, accessTokenExpiryTime)
	if err != nil {
		return dto.JWTToken{}, fmt.Sprintf("Failed to sign access token: %v", err)
	}
	refreshToken, err := GenerateToken(userID, nil, refreshTokenSecret, refreshTokenExpiryTime)
	if err != nil {
		return dto.JWTToken{}, fmt.Sprintf("Failed to sign refresh token: %v", err)
	}
	return dto.JWTToken{AccessToken: accessToken, RefreshToken: refreshToken}, ""
}

// GenerateSignUpJWTTokens generates a signup token
func GenerateSignUpJWTTokens(userID string) (string, string) {
	token, err := GenerateToken(userID, nil, signUpTokenSecret, signUpTokenExpiryTime)
	if err != nil {
		return "", fmt.Sprintf("Failed to sign signup token: %v", err)
	}
	return token, ""
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

// VerifySignUpTokenAndGetClaims verifies a signup token
func VerifySignUpTokenAndGetClaims(tokenString string) (bool, string, *TokenClaims) {
	claims := &TokenClaims{}
	isValid, msg, err := verifyToken(tokenString, signUpTokenSecret, claims)
	if err != nil {
		logger.Log(logger.DEBUG, fmt.Sprintf("Signup token verification failed: %v", err))
	}
	if isValid {
		logger.Log(logger.DEBUG, "Token is valid")
	} else {
		logger.Log(logger.DEBUG, "Token is not valid")
	}
	return isValid, msg, claims
}
