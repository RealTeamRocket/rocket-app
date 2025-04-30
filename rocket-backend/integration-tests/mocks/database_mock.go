package mocks

import (
	"database/sql"
	"errors"

	"github.com/google/uuid"

	"rocket-backend/internal/types"
)

type MockDB struct {
    GetUserByIDFunc            func(userID uuid.UUID) (types.User, error)
    UpdateDailyStepsFunc       func(userID uuid.UUID, steps int) error
    UpdateRocketPointsFunc     func(userID uuid.UUID, rocketPoints int) error
    HealthFunc                 func() map[string]string
    GetUserByEmailFunc         func(email string) (types.Credentials, error)
    CheckEmailFunc             func(email string) error
    SaveCredentialsFunc        func(creds types.Credentials) error
    SaveUserProfileFunc        func(user types.User) error
    GetSettingsByUserIDFunc    func(userID uuid.UUID) (*types.Settings, error)
    CreateSettingsFunc         func(settings types.Settings) error
    UpdateSettingsFunc         func(userId uuid.UUID, settings types.SettingsDTO, imageID uuid.UUID) error
    SaveImageFunc              func(filename string, data []byte) (uuid.UUID, error)
    GetUserImageFunc           func(userID uuid.UUID) (*types.UserImage, error)
    GetAllChallengesFunc       func() ([]types.Challenge, error)
    AssignChallengesToUserFunc func(userID uuid.UUID, challenges []types.Challenge) error
    GetUserDailyChallengesFunc func(userID uuid.UUID) ([]types.Challenge, error)
    ResetDailyChallengesFunc   func() error
    InsertChallengeFunc        func(challenge types.Challenge) error
    CompleteChallengeFunc      func(userID uuid.UUID, dto types.CompleteChallengesDTO) error
    IsNewDayForUserFunc        func(userID uuid.UUID) (bool, error)
    CleanUpChallengesForUserFunc func(userID uuid.UUID) error
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

func (m *MockDB) UpdateRocketPoints(userID uuid.UUID, rocketPoints int) error {
	if m.UpdateRocketPointsFunc != nil {
		return m.UpdateRocketPointsFunc(userID, rocketPoints)
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

func (m *MockDB) UpdateSettings(userId uuid.UUID, settings types.SettingsDTO, imageID uuid.UUID) error {
	if m.UpdateSettingsFunc != nil {
		return m.UpdateSettingsFunc(userId, settings, imageID)
	}
	return nil
}

func (m *MockDB) SaveImage(filename string, data []byte) (uuid.UUID, error) {
	if m.SaveImageFunc != nil {
		return m.SaveImageFunc(filename, data)
	}
	return uuid.Nil, nil
}

func (m *MockDB) GetUserImage(userID uuid.UUID) (*types.UserImage, error) {
	if m.GetUserImageFunc != nil {
		return m.GetUserImageFunc(userID)
	}
	return nil, nil
}

func (m *MockDB) GetAllChallenges() ([]types.Challenge, error) {
	if m.GetAllChallengesFunc != nil {
		return m.GetAllChallengesFunc()
	}
	return nil, nil
}

func (m *MockDB) AssignChallengesToUser(userID uuid.UUID, challenges []types.Challenge) error {
	if m.AssignChallengesToUserFunc != nil {
		return m.AssignChallengesToUserFunc(userID, challenges)
	}
	return nil
}

func (m *MockDB) GetUserDailyChallenges(userID uuid.UUID) ([]types.Challenge, error) {
	if m.GetUserDailyChallengesFunc != nil {
		return m.GetUserDailyChallengesFunc(userID)
	}
	return nil, nil
}

func (m *MockDB) ResetDailyChallenges() error {
	if m.ResetDailyChallengesFunc != nil {
		return m.ResetDailyChallengesFunc()
	}
	return nil
}

func (m *MockDB) InsertChallenge(challenge types.Challenge) error {
	if m.InsertChallengeFunc != nil {
		return m.InsertChallengeFunc(challenge)
	}
	return nil
}

func (m *MockDB) CompleteChallenge(userID uuid.UUID, dto types.CompleteChallengesDTO) error {
	if m.CompleteChallengeFunc != nil {
		return m.CompleteChallengeFunc(userID, dto)
	}
	return nil
}

func (m *MockDB) IsNewDayForUser(userID uuid.UUID) (bool, error) {
    if m.IsNewDayForUserFunc != nil {
        return m.IsNewDayForUserFunc(userID)
    }
    return false, nil
}

func (m *MockDB) CleanUpChallengesForUser(userID uuid.UUID) error {
    if m.CleanUpChallengesForUserFunc != nil {
        return m.CleanUpChallengesForUserFunc(userID)
    }
    return nil
}
