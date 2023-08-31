package handler

import (
	"net/http"
	"reqwizard/internal/domain"
	"reqwizard/internal/routes/auth"
	"reqwizard/internal/routes/ticketResponse"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	useCase ticketResponse.UseCase
}

func NewHandler(useCase ticketResponse.UseCase) *Handler {
	return &Handler{
		useCase: useCase,
	}
}

// CreateTicketResponse
// @Tags ticketResponses
// @Summary Create TicketResponses
// @Description Create a specific ticketResponse
// @Success 200 {object} domain.Response
// @Failure 400 {object} domain.BadResponse
// @Failure 401 {object} domain.BadResponse
// @Failure 403 {object} domain.BadResponse
// @Failure 404 {object} domain.BadResponse
// @Failure 500 {object} domain.BadResponse
// @Router /api/ticketResponse/v1/create [post].
func (h *Handler) CreateTicketResponse(c *gin.Context) {
	inp := new(ticketResponse.CreateTicketResponseInput)

	// c токена вытаскиваем
	if user, exist := c.Get(auth.CtxUserKey); exist {
		inp.ID = user.(*domain.User).ID
		inp.Email = user.(*domain.User).Email
	}

	if err := c.BindJSON(inp); err != nil {
		c.JSON(http.StatusBadRequest, domain.BadResponse{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		})
		return
	}

	if err := ticketResponse.ValidateCreateTicketResponseInput(inp); err != nil {
		c.JSON(http.StatusNotAcceptable, domain.BadResponse{
			Status:  http.StatusNotAcceptable,
			Message: err.Error(),
		})

		return
	}

	err := h.useCase.CreateTicketResponse(c.Request.Context(), inp)
	if err != nil {
		c.JSON(http.StatusNotAcceptable, domain.BadResponse{
			Status:  http.StatusNotAcceptable,
			Message: err.Error(),
		})
		return
	}

	c.Status(http.StatusOK)
}

// GetTicketResponses
// @Tags ticketResponses
// @Summary receiving ticketResponses
// @Description get all ticketResponses
// @Param search query string false "9999990000"
// @Param createdAtFrom query string false "2019-01-25T10:30:00.000Z"
// @Param createdAtTo query string false "2019-02-25T10:30:00.000Z"
// @Param updatedAtFrom query string false "2019-01-25T10:30:00.000Z"
// @Param updatedAtTo query string false "2019-02-25T10:30:00.000Z"
// @Success 200 {object} domain.ResponseTicketResponses
// @Failure 400 {object} domain.BadResponse
// @Failure 401 {object} domain.BadResponse
// @Failure 403 {object} domain.BadResponse
// @Failure 405 {object} domain.BadResponse
// @Failure 500 {object} domain.BadResponse
// @Router /api/ticket-response/v1/list [get].
func (h *Handler) GetTicketResponsesByManagerID(c *gin.Context) {
	inp := new(ticketResponse.TicketResponseListInput)

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

	if err := ticketResponse.ValidateTicketResponseListInput(inp); err != nil {
		c.JSON(http.StatusNotAcceptable, domain.BadResponse{
			Status:  http.StatusNotAcceptable,
			Message: err.Error(),
		})

		return
	}

	ticketResponses, err := h.useCase.GetTicketResponsesByManagerID(c.Request.Context(), inp)
	if err != nil {
		c.JSON(http.StatusNotAcceptable, domain.BadResponse{
			Status:  http.StatusNotAcceptable,
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, domain.ResponseTicketResponses{
		Status:  http.StatusOK,
		Payload: ticketResponses,
	})
}

// GetTicketResponseByID
// @Tags ticketResponses
// @Summary Get TicketResponse by ID
// @Description Get a specific ticketResponse by ID
// @Param id path string true "TicketResponse ID"
// @Success 200 {object} domain.ResponseTicketResponse
// @Failure 400 {object} domain.BadResponse
// @Failure 401 {object} domain.BadResponse
// @Failure 403 {object} domain.BadResponse
// @Failure 404 {object} domain.BadResponse
// @Failure 500 {object} domain.BadResponse
// @Router /api/ticket-response/v1/{id} [get].
func (h *Handler) GetTicketResponseByID(c *gin.Context) {
	ticketResponseID := c.Param("id")

	ticketResponse, err := h.useCase.GetTicketResponseByID(c.Request.Context(), ticketResponseID)
	if err != nil {
		c.JSON(http.StatusNotAcceptable, domain.BadResponse{
			Status:  http.StatusNotAcceptable,
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, domain.ResponseTicketResponse{
		Status:  http.StatusOK,
		Payload: ticketResponse,
	})
}

// GetTicketResponsesByUserID
// @Tags ticketResponses
// @Summary receiving ticketResponses
// @Description get all ticketResponses for user id
// @Param search query string false "9999990000"
// @Param createdAtFrom query string false "2019-01-25T10:30:00.000Z"
// @Param createdAtTo query string false "2019-02-25T10:30:00.000Z"
// @Param updatedAtFrom query string false "2019-01-25T10:30:00.000Z"
// @Param updatedAtTo query string false "2019-02-25T10:30:00.000Z"
// @Success 200 {object} domain.ResponseTicketResponses
// @Failure 400 {object} domain.BadResponse
// @Failure 401 {object} domain.BadResponse
// @Failure 403 {object} domain.BadResponse
// @Failure 405 {object} domain.BadResponse
// @Failure 500 {object} domain.BadResponse
// @Router /api/ticket-response/v1/my-list [get].
func (h *Handler) GetTicketResponsesByUserID(c *gin.Context) {
	inp := new(ticketResponse.TicketResponseListInput)

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

	if err := ticketResponse.ValidateTicketResponseListInput(inp); err != nil {
		c.JSON(http.StatusNotAcceptable, domain.BadResponse{
			Status:  http.StatusNotAcceptable,
			Message: err.Error(),
		})

		return
	}

	ticketResponses, err := h.useCase.GetTicketResponsesByUserID(c.Request.Context(), inp)
	if err != nil {
		c.JSON(http.StatusNotAcceptable, domain.BadResponse{
			Status:  http.StatusNotAcceptable,
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, domain.ResponseTicketResponses{
		Status:  http.StatusOK,
		Payload: ticketResponses,
	})
}