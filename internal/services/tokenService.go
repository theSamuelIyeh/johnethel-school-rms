package services

import (
	"fmt"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/thesamueliyeh/johnethel-rms/internal/database"
)

// token service struct
type tokenService struct {
	db database.Database
}

// token service factory
func NewTokenService(db database.Database) *tokenService {
	return &tokenService{db: db}
}

// struct for token
type Token struct {
	AccessToken     string
	RefreshToken    string
	AccessTokenExp  time.Time
	RefreshTokenExp time.Time
	UserId          string
	Role            string
}

// token service generate token method
func (s *tokenService) GenerateToken(user *database.User) (*Token, error) {
	secret := []byte(os.Getenv("JWT_SECRET"))
	accessTokenExp := time.Now().Add(15 * time.Minute)
	refreshTokenExp := time.Now().Add(4 * time.Hour)
	accessToken := jwt.New(jwt.SigningMethodHS256)
	claims := accessToken.Claims.(jwt.MapClaims)
	claims["exp"] = accessTokenExp.Unix()
	claims["role"] = user.Role
	if user.AdmissionNo == 0 {
		claims["admission_no"] = user.Email
	} else {
		claims["admission_no"] = user.AdmissionNo
	}
	claims["id"] = user.User.Id
	accessTokenString, err := accessToken.SignedString(secret)
	if err != nil {
		return nil, err
	}
	refreshToken := jwt.New(jwt.SigningMethodHS256)
	claims = refreshToken.Claims.(jwt.MapClaims)
	claims["exp"] = refreshTokenExp.Unix()
	claims["role"] = user.Role
	if user.AdmissionNo == 0 {
		claims["admission_no"] = user.Email
	} else {
		claims["admission_no"] = user.AdmissionNo
	}
	claims["id"] = user.User.Id
	refreshTokenString, err := refreshToken.SignedString(secret)
	if err != nil {
		return nil, err
	}
	var token = Token{AccessToken: accessTokenString, RefreshToken: refreshTokenString, AccessTokenExp: accessTokenExp, RefreshTokenExp: refreshTokenExp, UserId: user.User.Id, Role: user.Role}
	return &token, nil
}

// token service set refresh token method
func (s *tokenService) SetRefreshToken(token string, expiry time.Time, user_id string) error {
	err := s.db.StoreRefreshToken(token, user_id, expiry)
	if err != nil {
		return err
	}
	return nil
}

// token service clear tokens method
func (s *tokenService) ClearTokens(c *fiber.Ctx) error {
	accessToken := c.Cookies("access_token")
	id := c.Cookies("id")
	if id == "" {
		if accessToken == "" {
			return fmt.Errorf("no token found")
		}
		// clear refresh token in db
		parsedToken, err := jwt.Parse(accessToken, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("JWT_SECRET")), nil
		})
		if err != nil {
			return err
		}
		claims, ok := parsedToken.Claims.(jwt.MapClaims)
		if !ok || !parsedToken.Valid {
			return fmt.Errorf("invalid token")
		}
		userId := claims["id"].(string)
		err = s.db.ClearRefreshToken(userId)
		if err != nil {
			return err
		}
		c.Cookie(&fiber.Cookie{
			Name:  "access_token",
			Value: "",
			// Secure:  true,
			Path:     "/",
			HTTPOnly: true,
			Expires:  time.Now().Add(-1 * time.Hour),
		})
		c.Cookie(&fiber.Cookie{
			Name:  "role",
			Value: "",
			// Secure:  true,
			Path:     "/",
			HTTPOnly: true,
			Expires:  time.Now().Add(-1 * time.Hour),
		})
		return nil
	}
	err := s.db.ClearRefreshToken(id)
	if err != nil {
		return err
	}
	if accessToken != "" {
		c.Cookie(&fiber.Cookie{
			Name:  "access_token",
			Value: "",
			// Secure:  true,
			Path:     "/",
			HTTPOnly: true,
			Expires:  time.Now().Add(-1 * time.Hour),
		})
	}
	c.Cookie(&fiber.Cookie{
		Name:  "id",
		Value: "",
		// Secure:  true,
		Path:     "/",
		HTTPOnly: true,
		Expires:  time.Now().Add(-1 * time.Hour),
	})
	c.Cookie(&fiber.Cookie{
		Name:  "role",
		Value: "",
		// Secure:  true,
		Path:     "/",
		HTTPOnly: true,
		Expires:  time.Now().Add(-1 * time.Hour),
	})
	return nil
	// clear refresh token in db
}

// token service verify token method
func (s *tokenService) VerifyToken(token string) (bool, error) {
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	if err != nil {
		return false, err
	}
	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok || !parsedToken.Valid {
		return false, fmt.Errorf("invalid token")
	}
	exp := claims["exp"].(float64)
	expTime := time.Unix(int64(exp), 0)

	if time.Until(expTime) <= 2*time.Minute {
		if time.Until(expTime) <= 0 {
			return false, fmt.Errorf("token expired")
		}
		return true, nil
	}
	return false, nil
}

// token service refresh access token method
func (s *tokenService) RefreshAccessToken(token string) (*Token, error) {
	secret := []byte(os.Getenv("JWT_SECRET"))
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return secret, nil
	})
	if err != nil {
		return nil, err
	}
	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok || !parsedToken.Valid {
		return nil, fmt.Errorf("invalid token")
	}
	userId := claims["id"].(string)
	role := claims["role"].(string)
	admissionNo := claims["admission_no"].(float64)

	accessTokenExp := time.Now().Add(15 * time.Minute)
	NewToken := jwt.New(jwt.SigningMethodHS256)
	claims = NewToken.Claims.(jwt.MapClaims)
	claims["exp"] = accessTokenExp.Unix()
	claims["role"] = role
	claims["admission_no"] = admissionNo
	claims["id"] = userId
	accessTokenString, err := NewToken.SignedString(secret)
	if err != nil {
		return nil, err
	}
	var tokenRet = Token{AccessToken: accessTokenString, AccessTokenExp: accessTokenExp, UserId: userId}
	return &tokenRet, nil
}
