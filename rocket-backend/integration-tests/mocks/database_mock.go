package mocks

import (
	"database/sql"
	"errors"

	"github.com/google/uuid"

	"rocket-backend/internal/types"
)

type MockDB struct {
	ExecuteRawSQLFunc          func(query string) (sql.Result, error)
	QueryRowFunc               func(query string, args ...interface{}) *sql.Row
	HealthFunc                 func() map[string]string
	CloseFunc                  func() error
	SaveCredentialsFunc        func(creds types.Credentials) error
	GetUserByEmailFunc         func(email string) (types.Credentials, error)
	CheckEmailFunc             func(email string) error
	SaveUserProfileFunc        func(user types.User) error
	GetUserByIDFunc            func(userID uuid.UUID) (types.User, error)
	UpdateRocketPointsFunc     func(userID uuid.UUID, rocketPoints int) error
	GetUserIDByNameFunc        func(name string) (uuid.UUID, error)
	GetTopUsersFunc            func(limit int) ([]types.User, error)
	GetAllUsersFunc            func() ([]types.User, error) // Missing method
	UpdateDailyStepsFunc       func(userID uuid.UUID, steps int) error
	GetUserStatisticsFunc      func(userID uuid.UUID) ([]types.StepStatistic, error)
	GetSettingsByUserIDFunc    func(userID uuid.UUID) (*types.Settings, error)
	CreateSettingsFunc         func(settings types.Settings) error
	UpdateSettingsStepGoalFunc func(userID uuid.UUID, stepGoal int) error
	UpdateSettingsImageFunc    func(userID uuid.UUID, imageID uuid.UUID) error
	UpdateStepGoalFunc         func(userID uuid.UUID, stepGoal int) error
	UpdateImageFunc            func(userID uuid.UUID, imageID uuid.UUID) error
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
	AddFriendFunc              func(userID, friendID uuid.UUID) error
	GetFriendsFunc             func(userID uuid.UUID) ([]types.User, error)
	GetFriendsRankedByPointsFunc func(userID uuid.UUID) ([]types.User, error)
	DeleteFriendFunc             func(userID, friendID uuid.UUID) error
}

func (m *MockDB) ExecuteRawSQL(query string) (sql.Result, error) {
	if m.ExecuteRawSQLFunc != nil {
		return m.ExecuteRawSQLFunc(query)
	}
	return nil, errors.New("not implemented")
}

func (m *MockDB) QueryRow(query string, args ...interface{}) *sql.Row {
	if m.QueryRowFunc != nil {
		return m.QueryRowFunc(query, args...)
	}
	return nil
}

func (m *MockDB) Health() map[string]string {
	if m.HealthFunc != nil {
		return m.HealthFunc()
	}
	return map[string]string{"status": "up"}
}

func (m *MockDB) Close() error {
	if m.CloseFunc != nil {
		return m.CloseFunc()
	}
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

func (m *MockDB) UpdateRocketPoints(userID uuid.UUID, rocketPoints int) error {
	if m.UpdateRocketPointsFunc != nil {
		return m.UpdateRocketPointsFunc(userID, rocketPoints)
	}
	return nil
}

func (m *MockDB) GetUserIDByName(name string) (uuid.UUID, error) {
	if m.GetUserIDByNameFunc != nil {
		return m.GetUserIDByNameFunc(name)
	}
	return uuid.Nil, nil
}

func (m *MockDB) GetTopUsers(limit int) ([]types.User, error) {
	if m.GetTopUsersFunc != nil {
		return m.GetTopUsersFunc(limit)
	}
	return nil, nil
}

func (m *MockDB) UpdateDailySteps(userID uuid.UUID, steps int) error {
	if m.UpdateDailyStepsFunc != nil {
		return m.UpdateDailyStepsFunc(userID, steps)
	}
	return nil
}

func (m *MockDB) GetUserStatistics(userID uuid.UUID) ([]types.StepStatistic, error) {
	if m.GetUserStatisticsFunc != nil {
		return m.GetUserStatisticsFunc(userID)
	}
	return nil, nil
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

func (m *MockDB) AddFriend(userID, friendID uuid.UUID) error {
	if m.AddFriendFunc != nil {
		return m.AddFriendFunc(userID, friendID)
	}
	return nil
}

func (m *MockDB) GetFriends(userID uuid.UUID) ([]types.User, error) {
	if m.GetFriendsFunc != nil {
		return m.GetFriendsFunc(userID)
	}
	return nil, nil
}

func (m *MockDB) GetFriendsRankedByPoints(userID uuid.UUID) ([]types.User, error) {
	if m.GetFriendsRankedByPointsFunc != nil {
		return m.GetFriendsRankedByPointsFunc(userID)
	}
	return nil, nil
}

func (m *MockDB) DeleteFriend(userID, friendID uuid.UUID) error {
	if m.DeleteFriendFunc != nil {
		return m.DeleteFriendFunc(userID, friendID)
	}
	return nil
}

func (m *MockDB) UpdateSettingsStepGoal(userID uuid.UUID, stepGoal int) error {
	if m.UpdateSettingsStepGoalFunc != nil {
		return m.UpdateSettingsStepGoalFunc(userID, stepGoal)
	}
	return nil
}

func (m *MockDB) UpdateSettingsImage(userID uuid.UUID, imageID uuid.UUID) error {
	if m.UpdateSettingsImageFunc != nil {
		return m.UpdateSettingsImageFunc(userID, imageID)
	}
	return nil
}

func (m *MockDB) UpdateImage(userID uuid.UUID, imageID uuid.UUID) error {
	if m.UpdateImageFunc != nil {
		return m.UpdateImageFunc(userID, imageID)
	}
	return nil
}

func (m *MockDB) UpdateStepGoal(userID uuid.UUID, stepGoal int) error {
	if m.UpdateStepGoalFunc != nil {
		return m.UpdateStepGoalFunc(userID, stepGoal)
	}
	return nil
}

func (m *MockDB) GetAllUsers() ([]types.User, error) {
	if m.GetAllUsersFunc != nil {
		return m.GetAllUsersFunc()
	}
	return nil, nil
}
