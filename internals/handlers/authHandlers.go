package handlers

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/theSamuelIyeh/johnethel-school-rms/internals/auth"
	"github.com/theSamuelIyeh/johnethel-school-rms/internals/utils"
	"github.com/theSamuelIyeh/johnethel-school-rms/views/fragments"
	"github.com/theSamuelIyeh/johnethel-school-rms/views/pages"
)

// login form struct
type LoginDetails struct {
	Email    string `form:"email" validate:"required,email"`
	Password string `form:"password" validate:"required,min=6"`
}

// signup form struct
type SignupDetails struct {
	Email     string `form:"email" validate:"required,email"`
	Password  string `form:"password" validate:"required,min=6"`
	FirstName string `form:"first-name" validate:"required"`
	LastName  string `form:"last-name" validate:"required"`
	Phone     string `form:"phone" validate:"required"`
}

// forgot password details struct
type ForgotPasswordDetails struct {
	Email string `form:"email" validate:"required,email"`
}

// reset password  query struct
type ResetPasswordDetails struct {
	TokenHash  string `query:"token_hash" validate:"required"`
	VerifyType string `query:"type" validate:"required"`
}

// update new password struct
type UpdatePasswordDetails struct {
	NewPassword        string `form:"password1" validate:"required"`
	ConfirmNewPassword string `form:"password2" validate:"required"`
}

// get login form
func (h *Handler) GetLoginFormHandler(c echo.Context) error {
	title := "Login Page"

	templates := utils.TemplateStruct{
		Page:      pages.LoginPage(fragments.LoginFormComponent(), title),
		Title:     title,
		Component: fragments.LoginFormComponent(),
	}
	return utils.RenderOrdinaryTemplate(c, templates)
}

// get update password form
func (h *Handler) GetUpdatePasswordFormHandler(c echo.Context) error {
	title := "Update Password Page"

	templates := utils.TemplateStruct{
		Page:      pages.LoginPage(fragments.UpdatePasswordComponent(), title),
		Title:     title,
		Component: fragments.UpdatePasswordComponent(),
	}
	return utils.RenderOrdinaryTemplate(c, templates)
}

// get forgot pasword form
func (h *Handler) GetForgotPasswordFormHandler(c echo.Context) error {
	title := "Forgot Password Page"

	templates := utils.TemplateStruct{
		Page:      pages.LoginPage(fragments.ForgotPasswordComponent(), title),
		Title:     title,
		Component: fragments.ForgotPasswordComponent(),
	}
	return utils.RenderOrdinaryTemplate(c, templates)
}

// post signup api handler
func (h *Handler) CreateUserHandler(c echo.Context) error {
	u := new(SignupDetails)
	if err := c.Bind(u); err != nil {
		return c.HTML(http.StatusBadRequest, "<p x-show=\"showFormError\">An Error Occured</p>")
	}
	if err := c.Validate(u); err != nil {
		return c.HTML(http.StatusBadRequest, "<p x-show=\"showFormError\">Invalid Details</p>")
	}
	displayName := fmt.Sprintf("%s %s", u.FirstName, u.LastName)
	_, err := auth.SignupService(c, u.Email, u.Password, displayName, u.Phone)
	if err != nil {
		return c.HTML(http.StatusInternalServerError, "<p x-show=\"showFormError\">"+err.Error()+"</p>")
	}
	return utils.RedirectToUrl(c, "/dashboard")
}

// post login handler
func (h *Handler) LoginHandler(c echo.Context) error {
	// create new user struct
	u := new(LoginDetails)

	// bind incoming data to struct
	if err := c.Bind(u); err != nil {
		return c.HTML(http.StatusBadRequest, "<p x-show=\"showFormError\">An Error Occured</p>")
	}

	// validate incoming data
	if err := c.Validate(u); err != nil {
		return c.HTML(http.StatusBadRequest, "<p x-show=\"showFormError\">Invalid Details</p>")
	}

	// call login service
	user, err := auth.LoginService(c, u.Email, u.Password)
	if err != nil {
		errMsg := strings.Split(err.Error(), ":")
		errMsgPart := errMsg[1]
		c.Response().Header().Set("Hx-reswap", "#innerHTML")
		return c.HTML(http.StatusInternalServerError, "<p x-show=\"showFormError\">"+errMsgPart+"</p>")
	}

	// set cookies
	utils.CreateCookie(c, user.AccessToken, "access_token", user.ExpiresIn)
	utils.CreateCookie(c, user.RefreshToken, "refresh_token", 0)
	return utils.RedirectToUrl(c, "/dashboard/home")
}

// post logout handler
func (h *Handler) LogoutHandler(c echo.Context) error {
	accessToken := c.Get("access_token").(string)
	err := auth.LogoutService(c, accessToken)
	if err == nil {
		utils.CreateCookie(c, "Bearer fake_token", "access_token", -1)
		utils.CreateCookie(c, "Bearer fake_token", "refresh_token", -1)
	}
	utils.CreateCookie(c, "Bearer fake_token", "access_token", -1)
	utils.CreateCookie(c, "Bearer fake_token", "refresh_token", -1)
	return utils.RedirectToUrl(c, "/auth/login")
}

// post forgot password handler (sends email link)
func (h *Handler) ForgotPasswordHandler(c echo.Context) error {
	u := new(ForgotPasswordDetails)
	if err := c.Bind(u); err != nil {
		return c.HTML(http.StatusBadRequest, "<p x-show=\"showFormError\">An Error Occured</p>")

	}
	if err := c.Validate(u); err != nil {
		return c.HTML(http.StatusBadRequest, "<p x-show=\"showFormError\">Invalid Details</p>")

	}
	err := auth.Supabase.Auth.ResetPasswordForEmail(c.Request().Context(), u.Email)
	if err != nil {
		return c.HTML(http.StatusBadRequest, "<p x-show=\"showFormError\">"+err.Error()+"</p>")

	}
	return utils.RenderOrdinaryTemplate(c, utils.TemplateStruct{
		Title:     "Password Reset Confirm",
		Page:      pages.LoginPage(fragments.ConfirmEmailSentComponent(), "Password Reset Confirm"),
		Component: fragments.ConfirmEmailSentComponent()})
}

// post forgot password handler (called from supabase link in email)
func (h *Handler) ResetPasswordHandler(c echo.Context) error {
	u := new(ResetPasswordDetails)
	if err := c.Bind(u); err != nil {
		return c.HTML(http.StatusBadRequest, "<p x-show=\"showFormError\">An Error Occured</p>")
	}
	if err := c.Validate(u); err != nil {
		return c.HTML(http.StatusBadRequest, "<p x-show=\"showFormError\">Invalid Details</p>")
	}
	user, err := auth.VerifyOtpService(c, u.TokenHash, u.VerifyType)
	if err != nil {
		return c.HTML(http.StatusBadRequest, "<p x-show=\"showFormError\">"+err.Error()+"</p>")
	}
	accessToken := user.AccessToken

	// get user info
	newUser, err := auth.GetUserDataService(c, accessToken)
	if err != nil {
		return c.HTML(http.StatusBadRequest, "<p x-show=\"showFormError\">"+err.Error()+"</p>")
	}
	firstName := newUser.UserMetadata["first_name"].(string)
	lastName := newUser.UserMetadata["last_name"].(string)
	// reset user password
	_, err = auth.UpdateUserDataService(c, accessToken, map[string]interface{}{
		"password": firstName + lastName + "@123",
		"data": map[string]interface{}{
			"change_password": true,
		}})
	if err != nil {
		return c.HTML(http.StatusBadRequest, "<p x-show=\"showFormError\">"+err.Error()+"</p>")
	}
	return utils.RedirectToUrl(c, "/dashboard")
}

// update new password
func (h *Handler) UpdatePasswordHandler(c echo.Context) error {
	u := new(UpdatePasswordDetails)
	if err := c.Bind(u); err != nil {
		return c.HTML(http.StatusBadRequest, "<p x-show=\"showFormError\">An Error Occured</p>")
	}
	if err := c.Validate(u); err != nil {
		return c.HTML(http.StatusBadRequest, "<p x-show=\"showFormError\">Invalid Details</p>")
	}
	if u.NewPassword != u.ConfirmNewPassword {
		return c.HTML(http.StatusBadRequest, "<p x-show=\"showFormError\">Passwords do not match</p>")
	}
	accessToken := c.Get("access_token").(string)
	details := map[string]interface{}{
		"password": u.ConfirmNewPassword,
		"data": map[string]interface{}{
			"change_password": false,
		}}
	_, err := auth.UpdateUserDataService(c, accessToken, details)
	if err != nil {
		return utils.RedirectToUrl(c, "/auth/login")
	}
	return utils.RedirectToUrl(c, "/dashboard")
}
