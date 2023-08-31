package handler

import (
	"net/http"
	"reqwizard/internal/domain"
	"reqwizard/internal/routes/auth"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	useCase auth.UseCase
}

func NewHandler(useCase auth.UseCase) *Handler {
	return &Handler{
		useCase: useCase,
	}
}

// SignUp
// @Tags user
// @Description sign-up
// @Param user body auth.SignUpInput true "User body"
// @Success 200 {object} domain.Response
// @Failure 400 {object} domain.BadResponse
// @Failure 401 {object} domain.BadResponse
// @Failure 403 {object} domain.BadResponse
// @Failure 405 {object} domain.BadResponse
// @Failure 500 {object} domain.BadResponse
// @Router /api/auth/v1/sign-up [post].
func (h *Handler) SignUp(c *gin.Context) {
	inp := new(auth.SignUpInput)

	if err := c.BindJSON(inp); err != nil {
		c.JSON(http.StatusBadRequest, domain.BadResponse{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	if err := auth.ValidateSignUpInput(inp); err != nil {
		c.JSON(http.StatusNotAcceptable, domain.BadResponse{
			Status:  http.StatusNotAcceptable,
			Message: err.Error(),
		})

		return
	}

	if err := h.useCase.SignUp(c.Request.Context(), inp); err != nil {
		c.JSON(http.StatusNotAcceptable, domain.BadResponse{
			Status:  http.StatusNotAcceptable,
			Message: err.Error(),
		})
		return
	}

	c.Status(http.StatusOK)
}

// SendVerifyCode
// @Tags user
// @Description send verify code
// @Param user body auth.SendVerifyCodeInput true "user body"
// @Success 200 {object} domain.Response
// @Failure 400 {object} domain.BadResponse
// @Failure 401 {object} domain.BadResponse
// @Failure 403 {object} domain.BadResponse
// @Failure 405 {object} domain.BadResponse
// @Failure 500 {object} domain.BadResponse
// @Router /api/auth/v1/send-verify-code [post].
func (h *Handler) SendVerifyCode(c *gin.Context) {
	inp := new(auth.SendVerifyCodeInput)

	if err := c.BindJSON(inp); err != nil {
		c.JSON(http.StatusBadRequest, domain.BadResponse{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	if err := auth.ValidateSendVerifyCodeInput(inp); err != nil {
		c.JSON(http.StatusNotAcceptable, domain.BadResponse{
			Status:  http.StatusNotAcceptable,
			Message: err.Error(),
		})

		return
	}

	if err := h.useCase.SendVerifyCode(c.Request.Context(), inp); err != nil {
		c.JSON(http.StatusNotAcceptable, domain.BadResponse{
			Status:  http.StatusNotAcceptable,
			Message: err.Error(),
		})
		return
	}

	c.Status(http.StatusOK)
}

// CheckVerifyCode
// @Tags user
// @Description check verify code
// @Param user body auth.CheckVerifyCodeInput true "user body"
// @Success 200 {object} domain.Response
// @Failure 400 {object} domain.BadResponse
// @Failure 401 {object} domain.BadResponse
// @Failure 403 {object} domain.BadResponse
// @Failure 405 {object} domain.BadResponse
// @Failure 500 {object} domain.BadResponse
// @Router /api/auth/v1/check-verify-code [post].
func (h *Handler) CheckVerifyCode(c *gin.Context) {
	inp := new(auth.CheckVerifyCodeInput)

	if err := c.BindJSON(inp); err != nil {
		c.JSON(http.StatusBadRequest, domain.BadResponse{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	if err := auth.ValidateCheckVerifyCodeInput(inp); err != nil {
		c.JSON(http.StatusNotAcceptable, domain.BadResponse{
			Status:  http.StatusNotAcceptable,
			Message: err.Error(),
		})

		return
	}

	token, err := h.useCase.CheckVerifyCode(c.Request.Context(), inp)
	if err != nil {
		c.JSON(http.StatusNotAcceptable, domain.BadResponse{
			Status:  http.StatusNotAcceptable,
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, domain.Response{
		Status:  http.StatusOK,
		Payload: token,
	})
}

// SignIn
// @Tags user
// @Description sign-in
// @Param user body auth.SignInInput true "user body"
// @Success 200 {object} domain.Response
// @Failure 400 {object} domain.BadResponse
// @Failure 401 {object} domain.BadResponse
// @Failure 403 {object} domain.BadResponse
// @Failure 405 {object} domain.BadResponse
// @Failure 500 {object} domain.BadResponse
// @Router /api/auth/v1/sign-in [post].
func (h *Handler) SignIn(c *gin.Context) {
	inp := new(auth.SignInInput)

	if err := c.BindJSON(inp); err != nil {
		c.JSON(http.StatusBadRequest, domain.BadResponse{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	if err := auth.ValidateSignInInput(inp); err != nil {
		c.JSON(http.StatusNotAcceptable, domain.BadResponse{
			Status:  http.StatusNotAcceptable,
			Message: err.Error(),
		})

		return
	}

	if err := h.useCase.SignIn(c.Request.Context(), inp); err != nil {
		c.JSON(http.StatusUnauthorized, domain.BadResponse{
			Status:  http.StatusUnauthorized,
			Message: err.Error(),
		})

		return
	}

	c.Status(http.StatusOK)
}

// GetProfile
// @Tags user
// @Description get user profile
// @Success 200 {object} domain.ResponseUser
// @Failure 400 {object} domain.BadResponse
// @Failure 401 {object} domain.BadResponse
// @Failure 403 {object} domain.BadResponse
// @Failure 405 {object} domain.BadResponse
// @Failure 500 {object} domain.BadResponse
// @Router /api/auth/v1/get-profile [get].
func (h *Handler) GetProfile(c *gin.Context) {
	inp := new(auth.GetProfileInput)

	// c токена вытаскиваем
	if user, exist := c.Get(auth.CtxUserKey); exist {
		inp.ID = user.(*domain.User).ID
		inp.Email = user.(*domain.User).Email
	}

	user, err := h.useCase.GetProfile(c.Request.Context(), inp)
	if err != nil {
		c.JSON(http.StatusUnauthorized, domain.BadResponse{
			Status:  http.StatusUnauthorized,
			Message: err.Error(),
		})

		return
	}

	c.JSON(http.StatusOK, domain.ResponseUser{
		Status:  http.StatusOK,
		Payload: user,
	})
}