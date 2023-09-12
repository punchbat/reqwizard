package handler

import (
	"net/http"
	"reqwizard/internal/domain"
	"reqwizard/internal/routes/auth"
	"reqwizard/internal/shared/utils"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
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
// @Router /auth/v1/sign-up [post].
func (h *Handler) SignUp(c *gin.Context) {
	inp := new(auth.SignUpInput)

	inp.Email = c.PostForm("email")
	inp.Password = c.PostForm("password")
	inp.PasswordConfirm = c.PostForm("passwordConfirm")
	inp.Role = c.PostForm("role")
	inp.Name = c.PostForm("name")
	inp.Surname = c.PostForm("surname")
	inp.Gender = c.PostForm("gender")
	inp.Birthday = c.PostForm("birthday")

	avatar, header, err := c.Request.FormFile("avatar")
	if err == nil {
		defer avatar.Close()

		fileSize := header.Size
		maxSize := int64(2 * 1024 * 1024) // 2MB
		if fileSize > maxSize {
			c.JSON(http.StatusNotAcceptable, domain.BadResponse{
				Status:  http.StatusNotAcceptable,
				Message: "Avatar size exceeds the limit of 2MB",
			})
			return
		}

		if !utils.IsValidAvatarImageExtension(header.Filename) {
			c.JSON(http.StatusUnsupportedMediaType, domain.BadResponse{
				Status:  http.StatusUnsupportedMediaType,
				Message: "Invalid avatar type (allowed: png, jpeg, jpg)",
			})
			return
		}

		inp.Avatar = avatar
		inp.AvatarName = header.Filename
	} else if err != http.ErrMissingFile {
		c.JSON(http.StatusUnprocessableEntity, domain.BadResponse{
			Status:  http.StatusUnprocessableEntity,
			Message: "Error uploading avatar",
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

	status, err := h.useCase.SignUp(c.Request.Context(), inp)
	if err != nil {
		c.JSON(status, domain.BadResponse{
			Status:  status,
			Message: err.Error(),
		})
		return
	}

	c.Status(status)
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
// @Router /auth/v1/send-verify-code [post].
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

	status, err := h.useCase.SendVerifyCode(c.Request.Context(), inp)
	if err != nil {
		c.JSON(status, domain.BadResponse{
			Status:  status,
			Message: err.Error(),
		})
		return
	}

	c.Status(status)
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
// @Router /auth/v1/check-verify-code [post].
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

	token, status, err := h.useCase.CheckVerifyCode(c.Request.Context(), inp)
	if err != nil {
		c.JSON(status, domain.BadResponse{
			Status:  status,
			Message: err.Error(),
		})
		return
	}

	c.SetCookie(
		viper.GetString("auth.token.name"),
		token,
		int(time.Hour.Seconds()*viper.GetDuration("auth.token.ttl").Seconds()),
		viper.GetString("auth.token.path"),
		viper.GetString("auth.token.domain"),
		viper.GetBool("auth.token.secure"),
		viper.GetBool("auth.token.http_only"),
	)

	c.Status(status)
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
// @Router /auth/v1/sign-in [post].
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

	status, err := h.useCase.SignIn(c.Request.Context(), inp)
	if err != nil {
		c.JSON(status, domain.BadResponse{
			Status:  status,
			Message: err.Error(),
		})

		return
	}

	c.Status(status)
}

// GetMyProfile
// @Tags user
// @Description get my profile
// @Success 200 {object} domain.ResponseUser
// @Failure 400 {object} domain.BadResponse
// @Failure 401 {object} domain.BadResponse
// @Failure 403 {object} domain.BadResponse
// @Failure 405 {object} domain.BadResponse
// @Failure 500 {object} domain.BadResponse
// @Router /auth/v1/get-my-profile [get].
func (h *Handler) GetMyProfile(c *gin.Context) {
	var myID string

	// c токена вытаскиваем
	if user, exist := c.Get(auth.CtxUserKey); exist {
		myID = user.(*domain.User).ID
	}

	user, status, err := h.useCase.GetProfile(c.Request.Context(), myID)
	if err != nil {
		c.JSON(status, domain.BadResponse{
			Status:  status,
			Message: err.Error(),
		})

		return
	}

	c.JSON(status, domain.ResponseUser{
		Status:  status,
		Payload: user,
	})
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
// @Router /auth/v1/get-profile/{id} [get].
func (h *Handler) GetProfile(c *gin.Context) {
	ID := c.Param("id")

	user, status, err := h.useCase.GetProfile(c.Request.Context(), ID)
	if err != nil {
		c.JSON(status, domain.BadResponse{
			Status:  status,
			Message: err.Error(),
		})

		return
	}

	c.JSON(status, domain.ResponseUser{
		Status:  status,
		Payload: user,
	})
}

// UpdateProfile
// @Tags user
// @Description update user profile
// @Success 200 {object} domain.Response
// @Failure 400 {object} domain.BadResponse
// @Failure 401 {object} domain.BadResponse
// @Failure 403 {object} domain.BadResponse
// @Failure 405 {object} domain.BadResponse
// @Failure 500 {object} domain.BadResponse
// @Router /auth/v1/update-profile [put].
func (h *Handler) UpdateProfile(c *gin.Context) {
	inp := new(auth.UpdateInput)

	// c токена вытаскиваем
	if user, exist := c.Get(auth.CtxUserKey); exist {
		inp.ID = user.(*domain.User).ID
	}

	inp.UserRoles = c.PostFormArray("userRoles")
	inp.Name = c.PostForm("name")
	inp.Surname = c.PostForm("surname")
	inp.Gender = c.PostForm("gender")
	inp.Birthday = c.PostForm("birthday")

	avatar, header, err := c.Request.FormFile("avatar")
	if err == nil {
		defer avatar.Close()

		fileSize := header.Size
		maxSize := int64(2 * 1024 * 1024) // 2MB
		if fileSize > maxSize {
			c.JSON(http.StatusNotAcceptable, domain.BadResponse{
				Status:  http.StatusNotAcceptable,
				Message: "Avatar size exceeds the limit of 2MB",
			})
			return
		}

		if !utils.IsValidAvatarImageExtension(header.Filename) {
			c.JSON(http.StatusUnsupportedMediaType, domain.BadResponse{
				Status:  http.StatusUnsupportedMediaType,
				Message: "Invalid avatar type (allowed: png, jpeg, jpg)",
			})
			return
		}

		inp.Avatar = avatar
		inp.AvatarName = header.Filename
	} else if err != http.ErrMissingFile {
		c.JSON(http.StatusUnprocessableEntity, domain.BadResponse{
			Status:  http.StatusUnprocessableEntity,
			Message: "Error uploading avatar",
		})
		return
	}

	if err := auth.ValidateUpdateInput(inp); err != nil {
		c.JSON(http.StatusNotAcceptable, domain.BadResponse{
			Status:  http.StatusNotAcceptable,
			Message: err.Error(),
		})

		return
	}

	// Вызовите метод вашей use case для обновления профиля.
	status, err := h.useCase.UpdateProfile(c.Request.Context(), inp)
	if err != nil {
		c.JSON(status, domain.BadResponse{
			Status:  status,
			Message: err.Error(),
		})
		return
	}

	c.Status(status)
}

// Logout
// @Tags user
// @Description logout
// @Success 200 {object} domain.Response
// @Failure 400 {object} domain.BadResponse
// @Failure 401 {object} domain.BadResponse
// @Failure 403 {object} domain.BadResponse
// @Failure 405 {object} domain.BadResponse
// @Failure 500 {object} domain.BadResponse
// @Router /auth/v1/logout [post].
func (h *Handler) Logout(c *gin.Context) {
	// Удаление куки
	c.SetCookie(
		viper.GetString("auth.token.name"),
		"",
		-1,
		viper.GetString("auth.token.path"),
		viper.GetString("auth.token.domain"),
		viper.GetBool("auth.token.secure"),
		viper.GetBool("auth.token.http_only"),
	)

	c.Status(http.StatusOK)
}