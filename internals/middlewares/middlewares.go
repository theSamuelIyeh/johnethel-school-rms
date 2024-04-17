package middlewares

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"github.com/theSamuelIyeh/johnethel-school-rms/internals/auth"
	"github.com/theSamuelIyeh/johnethel-school-rms/internals/utils"
)

type Middleware struct {
	Config *utils.ApiCfg
}

func NewMiddleware(cfg *utils.ApiCfg) *Middleware {
	return &Middleware{Config: cfg}
}

func (m *Middleware) LoggedinUserMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// check if user is logged in
		// if not redirect to login page
		ctx := c.Request().Context()
		jwtSecret := os.Getenv("JWT_SECRET")

		// get access and refresh token from cookies
		accessCookie, accessToken, accessErr := utils.GetTokenFromCookie(c, "access_token")
		_, refreshToken, refreshErr := utils.GetTokenFromCookie(c, "refresh_token")

		if accessErr != nil || refreshErr != nil {
			utils.CreateCookie(c, "Bearer Fake token", "access_token", -1)
			utils.CreateCookie(c, "Bearer Fake token", "refresh_token", -1)
			return utils.RedirectToUrl(c, "/auth/login")
		}

		// parse and validate jwt
		token, err := jwt.Parse(accessToken, func(token *jwt.Token) (interface{}, error) {
			return []byte(jwtSecret), nil
		})
		if err != nil {
			utils.CreateCookie(c, "Bearer Fake token", "access_token", -1)
			utils.CreateCookie(c, "Bearer Fake token", "refresh_token", -1)
			return utils.RedirectToUrl(c, "/auth/login")
		}

		// check if token is valid
		if token.Valid {
			// check access cookie expiry for refresh
			if accessCookie.Expires.Before(time.Now().Add(5 * time.Minute)) {
				user, err := auth.Supabase.Auth.RefreshUser(ctx, accessToken, refreshToken)
				if err != nil {
					utils.CreateCookie(c, "Bearer Fake token", "access_token", -1)
					utils.CreateCookie(c, "Bearer Fake token", "refresh_token", -1)
					return utils.RedirectToUrl(c, "/auth/login")
				}

				// Refresh the cookie with new access token
				newAccessToken := user.AccessToken
				newRefreshToken := user.RefreshToken
				utils.CreateCookie(c, newAccessToken, "access_token", user.ExpiresIn)
				utils.CreateCookie(c, newRefreshToken, "refresh_token", 0)

				// set access token and user to context
				c.Set("access_token", newAccessToken)
				c.Set("user", user.User)
				c.Set("role", user.User.UserMetadata["role"].(string))
				c.Set("userName", user.User.UserMetadata["display_name"].(string))
				c.Set("avatar", user.User.UserMetadata["passport"].(string))

				// check if change password is on
				changePassword := user.User.UserMetadata["password_change"].(bool)
				if changePassword {
					return utils.RedirectToUrl(c, "/auth/updatepassword")
				}
			} else {
				// get the user using access token

				user, err := auth.GetUserDataService(c, accessToken)
				if err != nil {
					utils.CreateCookie(c, "Bearer Fake token", "access_token", -1)
					utils.CreateCookie(c, "Bearer Fake token", "refresh_token", -1)
					return utils.RedirectToUrl(c, "/auth/login")
				}

				// set access token and user to context
				c.Set("access_token", accessToken)
				c.Set("user", user)
				c.Set("role", user.UserMetadata["role"].(string))
				c.Set("userName", user.UserMetadata["display_name"].(string))
				c.Set("avatar", user.UserMetadata["passport"].(string))

				// check if change password is on
				changePassword := user.UserMetadata["password_change"].(bool)
				if changePassword {
					return utils.RedirectToUrl(c, "/auth/updatepassword")
				}
			}
			return next(c)
		}
		utils.CreateCookie(c, "Bearer Fake token", "access_token", -1)
		utils.CreateCookie(c, "Bearer Fake token", "refresh_token", -1)
		return utils.RedirectToUrl(c, "/auth/login")
	}
}

// middleware for unprotected routes
func (m *Middleware) LoggedOutUserMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		_, err := c.Cookie("access_token")
		if err != nil {
			utils.CreateCookie(c, "Bearer Fake token", "refresh_token", -1)
			return next(c)
		}
		return utils.RedirectToUrl(c, "/dashboard/home")
	}
}

// get dashboard info
// func (m *Middleware) GetDashboardInfoMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
// 	return func(c echo.Context) error {

// 		result, err := m.Config.DB.GetAllUnlockedTerms(c.Request().Context())
// 		if err != nil {
// 			fmt.Println(err)
// 		}
// 		c.Set("terms", result)
// 		return next(c)
// 	}
// }
