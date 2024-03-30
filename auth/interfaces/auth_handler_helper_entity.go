package rest

type LoginResult struct {
	Status int     `json:"status" binding:"required"`
	Token  *string `json:"refresh_token"`
}
