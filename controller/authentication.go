package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Registration Documentation godoc
// @Summary Endpoint to send user data for registration
// @Description Send user data to the main authentication endpoint.
// @Tags Authorization
// @Accept json
// @Produce json
// @Param RegistrationData body RegistrationData true  "The User Registration data"
// @Success 200 {object} UserData "User Registration default data"
// @Failure 400 {object} ResponseObject "API response object"
// @Router /apicore/authentication/registration [post]
func (c *Controller) Registration(ctx *gin.Context) {
	ctx.String(http.StatusOK, "pong")
}

// Registration Verification Documentation godoc
// @Summary Endpoint to verify user registration link
// @Description Send user data and verification code to verify validity of verification.
// @Tags Authorization
// @Accept json
// @Produce json
// @Param RegistrationData body VerifyRegistration true  "The User Registration data and code"
// @Success 200 {object} RegistrationObject "User Registration default data"
// @Failure 400 {object} ResponseObject "API response object"
// @Router /apicore/authentication/registration/verify [post]
func (c *Controller) RegistrationVerification(ctx *gin.Context) {
	ctx.String(http.StatusOK, "pong")
}

// Login Documentation godoc
// @Summary Endpoint to send user data for login
// @Description Send user data to the main authentication login endpoint endpoint.
// @Tags Authorization
// @Accept json
// @Produce json
// @Param LoginData body LoginData true  "The User Login data"
// @Success 200 {object} LoginObject "User Registration default data"
// @Failure 400 {object} ResponseObject "API response object"
// @Router /apicore/authentication/login [post]
func (c *Controller) Login(ctx *gin.Context) {
	ctx.String(http.StatusOK, "pong")
}

// Forget Password Documentation godoc
// @Summary Endpoint to request for new user password
// @Description When a user forgets password, forgetPassword endpoint handles sending a new link for password recovery
// @Tags Authorization
// @Accept json
// @Produce json
// @Param string body string true  "The user email"
// @Success 200 {object} ResponseObject "Check email for password mail and link"
// @Failure 400 {object} ResponseObject "API response object"
// @Router /apicore/authentication/pasword/forget [post]
func (c *Controller) ForgetPassword(ctx *gin.Context) {
	ctx.String(http.StatusOK, "pong")
}

// Change Password Documentation godoc
// @Summary Change password updates a user password with a current one
// @Description When a user decides to change password, frontend call api with old and new password
// @Tags Authorization
// @Accept json
// @Produce json
// @Param ChangePassword body ChangePassword true  "The new password changes"
// @Success 200 {object} ResponseObject "Check email for password mail and link"
// @Failure 400 {object} ResponseObject "API response object"
// @Router /apicore/authentication/pasword/change [post]
func (c *Controller) ChangePassword(ctx *gin.Context) {
	ctx.String(http.StatusOK, "pong")
}

// VerifyToken checks if a user token is valid
// @Summary Checks if an authenticated user token is valid.
// @Description Pass token in the Authorization header. Verify token checks if the token is verified
// @Tags Authorization
// @Accept json
// @Produce json
// @Success 200 {object} VerifyToken "Check validity"
// @Failure 400 {object} ResponseObject "API response object"
// @Router /apicore/token/verify [get]
func (c *Controller) VerifyToken(ctx *gin.Context) {
	ctx.String(http.StatusOK, "pong")
}

// GetUserDataFromToken gets a user information from the
// @Summary Checks if an authenticated user token is valid, and send user data of authenticated user
// @Description Pass token in the Authorization header. Verify token checks if the token is verified and send the authenticated user application related data
// @Tags Authorization
// @Accept json
// @Produce json
// @Success 200 {object} LoginObject "Check validity"
// @Failure 400 {object} ResponseObject "API response object"
// @Router /apicore/token/user [get]
func (c *Controller) GetUserDataFromToken(ctx *gin.Context) {
	ctx.String(http.StatusOK, "pong")
}

// GetUserDataFromUID gets a user information from the unique user id
// @Summary Retrieve user data from user id
// @Description Get user uid from url and retrieve user data with the uid string
// @Tags Authorization
// @Accept json
// @Produce json
// @Success 200 {object} UserData "User related data"
// @Failure 400 {object} ResponseObject "API response object"
// @Router /apicore/user/{uid} [get]
func (c *Controller) GetUserDataFromUID(ctx *gin.Context) {
	ctx.String(http.StatusOK, "pong")
}

// UpdateUserData gets a user information from the unique user id
// @Summary Update user information
// @Description Update user data. This happens when a user wants to update user information
// @Tags Authorization
// @Param UserData body UserData true  "Update user data json object"
// @Accept json
// @Produce json
// @Success 200 {object} UserData "User related data"
// @Failure 400 {object} ResponseObject "API response object"
// @Router /apicore/user/ [patch]
func (c *Controller) UpdateUserData(ctx *gin.Context) {
	ctx.String(http.StatusOK, "pong")
}

// ResetToken sends a new token authentication
// @Summary Reset token renews old token string
// @Description Pass token in the Authorization header. Verify the token and send a new one with updated credentials
// @Tags Authorization
// @Accept json
// @Produce json
// @Success 200 {object} VerifyToken "Check validity"
// @Failure 400 {object} ResponseObject "API response object"
// @Router /apicore/token/reset [get]
func (c *Controller) ResetToken(ctx *gin.Context) {
	ctx.String(http.StatusOK, "pong")
}
