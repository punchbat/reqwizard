package handler

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"path/filepath"
	"reqwizard/internal/domain"
	"reqwizard/internal/routes/application"
	"reqwizard/internal/routes/auth"
	"reqwizard/internal/shared/utils"
	"strings"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	useCase application.UseCase
}

func NewHandler(useCase application.UseCase) *Handler {
	return &Handler{
		useCase: useCase,
	}
}

// CreateApplication
// @Tags applications
// @Summary Create Application
// @Description Create a specific application
// @Param application body application.CreateApplicationInput true "Application body"
// @Success 200 {object} domain.Response
// @Failure 400 {object} domain.BadResponse
// @Failure 401 {object} domain.BadResponse
// @Failure 403 {object} domain.BadResponse
// @Failure 404 {object} domain.BadResponse
// @Failure 500 {object} domain.BadResponse
// @Router /api/application/v1/create [post].
func (h *Handler) CreateApplication(c *gin.Context) {
	if !strings.HasPrefix(c.GetHeader("Content-Type"), "multipart/form-data") {
		c.JSON(http.StatusBadRequest, domain.BadResponse{
			Status:  http.StatusBadRequest,
			Message: "Request Content-Type must be multipart/form-data",
		})
		return
	}

	inp := new(application.CreateApplicationInput)

	// c токена вытаскиваем
	if user, exist := c.Get(auth.CtxUserKey); exist {
		inp.ID = user.(*domain.User).ID
		inp.Email = user.(*domain.User).Email
	}

	inp.Type = c.PostForm("type")
	inp.SubType = c.PostForm("subType")
	inp.Title = c.PostForm("title")
	inp.Description = c.PostForm("description")

	file, header, err := c.Request.FormFile("file")
	if err == nil {
		defer file.Close()

		fileBytes, err := ioutil.ReadAll(file)
		if err != nil {
			c.JSON(http.StatusInternalServerError, domain.BadResponse{
				Status:  http.StatusInternalServerError,
				Message: "Error reading file",
			})
			return
		}

		// * maximum 3MB
		if len(fileBytes) > 3*1024*1024 {
			c.JSON(http.StatusBadRequest, domain.BadResponse{
				Status:  http.StatusBadRequest,
				Message: "File size exceeds the limit of 5MB",
			})
			return
		}

		if !utils.IsValidFileExtension(header.Filename) {
			c.JSON(http.StatusBadRequest, domain.BadResponse{
				Status:  http.StatusBadRequest,
				Message: "Invalid file type (allowed: txt, json)",
			})
			return
		}

		inp.File = fileBytes
		inp.FileExtension = filepath.Ext(header.Filename)
	} else if err != http.ErrMissingFile {
		c.JSON(http.StatusBadRequest, domain.BadResponse{
			Status:  http.StatusBadRequest,
			Message: "Error uploading file",
		})
		return
	}

	if err := application.ValidateCreateApplicationInput(inp); err != nil {
		c.JSON(http.StatusNotAcceptable, domain.BadResponse{
			Status:  http.StatusNotAcceptable,
			Message: err.Error(),
		})

		return
	}

	err = h.useCase.CreateApplication(c.Request.Context(), inp)
	if err != nil {
		c.JSON(http.StatusNotAcceptable, domain.BadResponse{
			Status:  http.StatusNotAcceptable,
			Message: err.Error(),
		})
		return
	}

	c.Status(http.StatusOK)
}

// GetFile
// @Tags applications
// @Summary Download file
// @Description download file
// @Param filename path string true "File name"
// @Success 200 {object} domain.Response
// @Failure 400 {object} domain.BadResponse
// @Failure 401 {object} domain.BadResponse
// @Failure 403 {object} domain.BadResponse
// @Failure 404 {object} domain.BadResponse
// @Failure 500 {object} domain.BadResponse
// @Router /api/application/v1/download-file/{:fileName} [get].
func (h *Handler) GetFile(c *gin.Context) {
	fileName := c.Param("fileName")

	var userID string
	if user, exist := c.Get(auth.CtxUserKey); exist {
		userID = user.(*domain.User).ID
	}

	fileContents, mimeType, err := h.useCase.GetFile(c.Request.Context(), userID, fileName)
	if err != nil {
		c.JSON(http.StatusNotAcceptable, domain.BadResponse{
			Status:  http.StatusNotAcceptable,
			Message: err.Error(),
		})
		return
	}

	c.Header("Content-Disposition", fmt.Sprintf("attachment; filename=%s", fileName))
	c.Data(http.StatusOK, mimeType, fileContents)
}

// GetApplicationByID
// @Tags applications
// @Summary Get Application by ID
// @Description Get a specific application by ID
// @Param id path string true "Application ID"
// @Success 200 {object} domain.ResponseApplication
// @Failure 400 {object} domain.BadResponse
// @Failure 401 {object} domain.BadResponse
// @Failure 403 {object} domain.BadResponse
// @Failure 404 {object} domain.BadResponse
// @Failure 500 {object} domain.BadResponse
// @Router /api/application/v1/{id} [get].
func (h *Handler) GetApplicationByID(c *gin.Context) {
	applicationID := c.Param("id")

	application, err := h.useCase.GetApplicationByID(c.Request.Context(), applicationID)
	if err != nil {
		c.JSON(http.StatusNotAcceptable, domain.BadResponse{
			Status:  http.StatusNotAcceptable,
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, domain.ResponseApplication{
		Status:  http.StatusOK,
		Payload: application,
	})
}

// GetApplicationsByUserID
// @Tags applications
// @Summary receiving applications
// @Description get all applications for user id
// @Param search query string false "9999990000"
// @Param status query []string false "string enums" Enums(canceled, waiting, working, done)
// @Param type query []string false "string enums" Enums(general, financial)
// @Param subType query []string false "string enums" Enums(information, account_help, refunds, payment)
// @Param createdAtFrom query string false "2019-01-25T10:30:00.000Z"
// @Param createdAtTo query string false "2019-02-25T10:30:00.000Z"
// @Param updatedAtFrom query string false "2019-01-25T10:30:00.000Z"
// @Param updatedAtTo query string false "2019-02-25T10:30:00.000Z"
// @Success 200 {object} domain.ResponseApplications
// @Failure 400 {object} domain.BadResponse
// @Failure 401 {object} domain.BadResponse
// @Failure 403 {object} domain.BadResponse
// @Failure 405 {object} domain.BadResponse
// @Failure 500 {object} domain.BadResponse
// @Router /api/application/v1/my-list [get].
func (h *Handler) GetApplicationsByUserID(c *gin.Context) {
	inp := new(application.ApplicationListInput)

	// c токена вытаскиваем
	if user, exist := c.Get(auth.CtxUserKey); exist {
		inp.ID = user.(*domain.User).ID
		inp.Email = user.(*domain.User).Email
	}

	if err := c.ShouldBindQuery(inp); err != nil {
		c.JSON(http.StatusBadRequest, domain.BadResponse{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	statuses := utils.RemoveEmptyStrings(strings.Split(c.Query("statuses"), ","))
	if len(statuses) > 0 {
		inp.Statuses = statuses
	}
	types := utils.RemoveEmptyStrings(strings.Split(c.Query("types"), ","))
	if len(types) > 0 {
		inp.Types = types
	}
	subTypes := utils.RemoveEmptyStrings(strings.Split(c.Query("subTypes"), ","))
	if len(subTypes) > 0 {
		inp.SubTypes = subTypes
	}

	if err := application.ValidateApplicationListInput(inp); err != nil {
		c.JSON(http.StatusNotAcceptable, domain.BadResponse{
			Status:  http.StatusNotAcceptable,
			Message: err.Error(),
		})

		return
	}

	applications, err := h.useCase.GetApplicationsByUserID(c.Request.Context(), inp)
	if err != nil {
		c.JSON(http.StatusNotAcceptable, domain.BadResponse{
			Status:  http.StatusNotAcceptable,
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, domain.ResponseApplications{
		Status:  http.StatusOK,
		Payload: applications,
	})
}

// GetApplications
// @Tags applications
// @Summary receiving applications
// @Description get all applications
// @Param search query string false "9999990000"
// @Param status query []string false "string enums" Enums(canceled, waiting, working, done)
// @Param type query []string false "string enums" Enums(general, financial)
// @Param subType query []string false "string enums" Enums(information, account_help, refunds, payment)
// @Param createdAtFrom query string false "2019-01-25T10:30:00.000Z"
// @Param createdAtTo query string false "2019-02-25T10:30:00.000Z"
// @Param updatedAtFrom query string false "2019-01-25T10:30:00.000Z"
// @Param updatedAtTo query string false "2019-02-25T10:30:00.000Z"
// @Success 200 {object} domain.ResponseApplications
// @Failure 400 {object} domain.BadResponse
// @Failure 401 {object} domain.BadResponse
// @Failure 403 {object} domain.BadResponse
// @Failure 405 {object} domain.BadResponse
// @Failure 500 {object} domain.BadResponse
// @Router /api/application/v1/list [get].
func (h *Handler) GetApplications(c *gin.Context) {
	inp := new(application.ApplicationListInput)

	// c токена вытаскиваем
	if user, exist := c.Get(auth.CtxUserKey); exist {
		inp.ID = user.(*domain.User).ID
		inp.Email = user.(*domain.User).Email
	}

	if err := c.ShouldBindQuery(inp); err != nil {
		c.JSON(http.StatusBadRequest, domain.BadResponse{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	statuses := utils.RemoveEmptyStrings(strings.Split(c.Query("statuses"), ","))
	if len(statuses) > 0 {
		inp.Statuses = statuses
	}
	types := utils.RemoveEmptyStrings(strings.Split(c.Query("types"), ","))
	if len(types) > 0 {
		inp.Types = types
	}
	subTypes := utils.RemoveEmptyStrings(strings.Split(c.Query("subTypes"), ","))
	if len(subTypes) > 0 {
		inp.SubTypes = subTypes
	}

	if err := application.ValidateApplicationListInput(inp); err != nil {
		c.JSON(http.StatusNotAcceptable, domain.BadResponse{
			Status:  http.StatusNotAcceptable,
			Message: err.Error(),
		})

		return
	}

	applications, err := h.useCase.GetApplications(c.Request.Context(), inp)
	if err != nil {
		c.JSON(http.StatusNotAcceptable, domain.BadResponse{
			Status:  http.StatusNotAcceptable,
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, domain.ResponseApplications{
		Status:  http.StatusOK,
		Payload: applications,
	})
}