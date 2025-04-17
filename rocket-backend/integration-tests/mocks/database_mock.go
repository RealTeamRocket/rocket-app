package mocks

import (
	"database/sql"
	"errors"

	"github.com/google/uuid"

	"rocket-backend/internal/types"
)

type MockDB struct {
	GetUserByIDFunc         func(userID uuid.UUID) (types.User, error)
	UpdateDailyStepsFunc    func(userID uuid.UUID, steps int) error
	HealthFunc              func() map[string]string
	GetUserByEmailFunc      func(email string) (types.Credentials, error)
	CheckEmailFunc          func(email string) error
	SaveCredentialsFunc     func(creds types.Credentials) error
	SaveUserProfileFunc     func(user types.User) error
	GetSettingsByUserIDFunc func(userID uuid.UUID) (*types.Settings, error)
	CreateSettingsFunc      func(settings types.Settings) error
	UpdateSettingsFunc 		func(userId uuid.UUID, settings types.SettingsDTO) error
}

func (m *MockDB) ExecuteRawSQL(query string) (sql.Result, error) {
	return nil, errors.New("not implemented")
}

func (m *MockDB) QueryRow(query string, args ...interface{}) *sql.Row {
	return nil
}

func (m *MockDB) Health() map[string]string {
	if m.HealthFunc != nil {
		return m.HealthFunc()
	}
	return map[string]string{"status": "up"}
}

func (m *MockDB) Close() error {
	return nil
}

func (m *MockDB) SaveCredentials(creds types.Credentials) error {
	if m.SaveCredentialsFunc != nil {
		return m.SaveCredentialsFunc(creds)
	}
	return nil
}

func (m *MockDB) GetUserByEmail(email string) (types.Credentials, error) {
	if m.GetUserByEmailFunc != nil {
		return m.GetUserByEmailFunc(email)
	}
	return types.Credentials{}, nil
}

func (m *MockDB) CheckEmail(email string) error {
	if m.CheckEmailFunc != nil {
		return m.CheckEmailFunc(email)
	}
	return nil
}

func (m *MockDB) SaveUserProfile(user types.User) error {
	if m.SaveUserProfileFunc != nil {
		return m.SaveUserProfileFunc(user)
	}
	return nil
}

func (m *MockDB) GetUserByID(userID uuid.UUID) (types.User, error) {
	if m.GetUserByIDFunc != nil {
		return m.GetUserByIDFunc(userID)
	}
	return types.User{}, nil
}

func (m *MockDB) UpdateDailySteps(userID uuid.UUID, steps int) error {
	if m.UpdateDailyStepsFunc != nil {
		return m.UpdateDailyStepsFunc(userID, steps)
	}
	return nil
}

func (m *MockDB) GetSettingsByUserID(userID uuid.UUID) (*types.Settings, error) {
	if m.GetSettingsByUserIDFunc != nil {
		return m.GetSettingsByUserIDFunc(userID)
	}
	return nil, nil
}

func (m *MockDB) CreateSettings(settings types.Settings) error {
	if m.CreateSettingsFunc != nil {
		return m.CreateSettingsFunc(settings)
	}
	return nil
}

func (m *MockDB) UpdateSettings(userId uuid.UUID, settings types.SettingsDTO) error {
	if m.UpdateSettingsFunc != nil {
		return m.UpdateSettingsFunc(userId, settings)
	}
	return nil
}
