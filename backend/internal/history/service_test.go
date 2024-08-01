package history

import (
	"backend/internal/domain"
	"context"
	"errors"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"testing"
)

type mockRepository struct {
	mock.Mock
}

func (mr *mockRepository) GetAll(ctx context.Context) ([]domain.History, error) {
	args := mr.Called(ctx)
	return args.Get(0).([]domain.History), args.Error(1)
}

func (mr *mockRepository) Save(ctx context.Context, history domain.History) (int, error) {
	args := mr.Called(ctx, history)
	return args.Int(0), args.Error(1)
}

func TestServiceGetAll_Successful(t *testing.T) {
	// Given
	expectedHistories := []domain.History{
		{
			UserMessage: "Test",
			BotResponse: "Test",
			Feedback:    true,
		},
		{
			UserMessage: "Test 2",
			BotResponse: "Test 2",
			Feedback:    false,
		},
	}

	mr := new(mockRepository)
	mr.On("GetAll", mock.Anything).Return(expectedHistories, nil)

	service := NewService(mr)

	// When
	histories, err := service.GetAll(context.Background())

	// Then
	require.NoError(t, err)
	require.Len(t, histories, len(expectedHistories))
	require.EqualValues(t, expectedHistories, histories)
}

func TestServiceGetAll_SuccessfulWithZeroUsers(t *testing.T) {
	// Given
	expectedHistories := make([]domain.History, 0)

	mr := new(mockRepository)
	mr.On("GetAll", mock.Anything).Return(expectedHistories, nil)

	service := NewService(mr)

	// When
	histories, err := service.GetAll(context.Background())

	// Then
	require.NoError(t, err)
	require.Len(t, histories, len(expectedHistories))
	require.EqualValues(t, expectedHistories, histories)
}

func TestServiceSave_Successful(t *testing.T) {
	// Given
	expectedHistory := domain.History{
		UserMessage: "Test",
		BotResponse: "Test",
		Feedback:    true,
	}

	mr := new(mockRepository)
	mr.On("Save", mock.Anything, expectedHistory).Return(expectedHistory.ID, nil)

	service := NewService(mr)

	// When
	history, err := service.Save(context.Background(), expectedHistory)

	// Then
	require.NoError(t, err)
	require.Equal(t, expectedHistory, history)
}

func TestServiceSave_FailsDueToRepositoryError(t *testing.T) {
	// Given
	expectedHistory := domain.History{}
	expectedError := errors.New("Error Code: 1054. Unknown column 'wrong' in 'field list'")

	mr := new(mockRepository)
	mr.On("Save", mock.Anything, expectedHistory).Return(0, expectedError)

	service := NewService(mr)

	// When
	history, err := service.Save(context.Background(), expectedHistory)

	// Then
	require.ErrorContains(t, err, "Error Code: 1054")
	require.ErrorContains(t, err, "Unknown column 'wrong' in 'field list'")
	require.Equal(t, expectedHistory, history)
}
