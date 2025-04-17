package types

type RegisterDTO struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type UpdateStepsDTO struct {
	Steps int `json:"steps"`
}

type SettingsDTO struct {
	StepGoal int `json:"dailySteps"`
	ProfImg string `json:"profImg"`
}
