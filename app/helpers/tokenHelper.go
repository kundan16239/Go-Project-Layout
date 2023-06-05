package helpers

import (
	"errors"
	"go-folder-sample/app/models"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/spf13/viper"
)

// ValidateToken validates the provided JWT token and returns the claims or an error message.
func ValidateAccessToken(tokenString string) (*models.CustomClaims, error) {
	// Define your JWT secret key
	secretKey := viper.GetString("jwtAccessSecret")

	// Parse the JWT token
	token, err := jwt.ParseWithClaims(tokenString, &models.CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})

	// Handle parsing errors
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return nil, errors.New("invalid token signature")
		}
		return nil, errors.New("failed to parse token")
	}

	// Validate the token
	if claims, ok := token.Claims.(*models.CustomClaims); ok && token.Valid {
		if claims.ExpiresAt < time.Now().Unix() {
			return nil, errors.New("token has expired")
		}
		return claims, nil
	}

	return nil, errors.New("invalid token")
}

func ValidateRefreshToken(tokenString string) (*models.CustomClaims, error) {
	// Define your JWT secret key
	refreshKey := viper.GetString("jwtRefreshSecret")

	// Parse the JWT token
	token, err := jwt.ParseWithClaims(tokenString, &models.CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(refreshKey), nil
	})

	// Handle parsing errors
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return nil, errors.New("invalid token signature")
		}
		return nil, errors.New("failed to parse token")
	}

	// Validate the token
	if claims, ok := token.Claims.(*models.CustomClaims); ok && token.Valid {
		if claims.ExpiresAt < time.Now().Unix() {
			return nil, errors.New("token has expired")
		}
		return claims, nil
	}

	return nil, errors.New("invalid token")
}

func GenerateAllTokens(identifier string) (signedToken string, signedRefreshToken string, err error) {
	jwtAccessSecret := viper.GetString("jwtAccessSecret")
	jwtRefreshSecret := viper.GetString("jwtRefreshSecret")

	accessToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Id:        identifier,
		ExpiresAt: time.Now().Local().Add(time.Minute * viper.GetDuration("accessTokenExpiryMinute")).Unix(),
	}).SignedString([]byte(jwtAccessSecret))
	if err != nil {
		return "", "", err
	}

	refreshToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Id:        identifier,
		ExpiresAt: time.Now().Local().Add(time.Hour * viper.GetDuration("refreshTokenExpiryHour")).Unix(),
	}).SignedString([]byte(jwtRefreshSecret))
	if err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil
}
