package domain

type LoginBody struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginResult struct {
	Status int     `json:"status" binding:"required"`
	Data   *string `json:"data"`
}
