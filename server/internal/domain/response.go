package domain

type Response struct {
	Status  int         `json:"status"`
	Payload interface{} `json:"payload"`
}

type ResponseUser struct {
	Status  int   `json:"status"`
	Payload *User `json:"payload"`
}

type ResponseRoles struct {
	Status  int     `json:"status"`
	Payload []*Role `json:"payload"`
}

type ResponseApplication struct {
	Status  int          `json:"status"`
	Payload *Application `json:"payload"`
}

type ResponseApplications struct {
	Status  int            `json:"status"`
	Payload []*Application `json:"payload"`
}

type ResponseTicketResponse struct {
	Status  int             `json:"status"`
	Payload *TicketResponse `json:"payload"`
}

type ResponseTicketResponses struct {
	Status  int               `json:"status"`
	Payload []*TicketResponse `json:"payload"`
}

type BadResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}