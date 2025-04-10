package types

type RegisterDTO struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type UpdateStepsDTO struct {
	Steps int `json:"steps"`
}
