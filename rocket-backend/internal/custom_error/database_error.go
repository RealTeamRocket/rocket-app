package custom_error

import "errors"

var (
	ErrEmailAlreadyExists   = errors.New("email already exists")
	ErrDatabaseQuery        = errors.New("database query error")
	ErrFailedToSave         = errors.New("failed to save data")
	ErrFailedToRetrieveData = errors.New("failed to retrieve data")
	ErrUserNotFound         = errors.New("user not found")
	ErrFailedToUpdate       = errors.New("failed to update data")
	ErrChallengeNotFound    = errors.New("challenge not found")
	ErrImageNotFound        = errors.New("image not found")
	ErrSettingsNotFound     = errors.New("settings not found")
	ErrFailedToDelete       = errors.New("failed to delete data")
	ErrFailedToLoad         = errors.New("failed to load data")
)
