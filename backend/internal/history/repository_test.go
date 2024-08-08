package history

import (
	"backend/internal/domain"
	"context"
	"errors"
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/require"
	"regexp"
	"testing"
)

func TestRepositoryGetAll_Successful(t *testing.T) {
	// Given
	db, mock, err := sqlmock.New()
	require.NoError(t, err)

	defer db.Close()

	dbx := sqlx.NewDb(db, "sqlmock")

	ctx := context.Background()

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

	columns := []string{"user_message", "bot_response", "feedback"}
	rows := sqlmock.NewRows(columns)
	rows.AddRow(expectedHistories[0].UserMessage, expectedHistories[0].BotResponse, expectedHistories[0].Feedback)
	rows.AddRow(expectedHistories[1].UserMessage, expectedHistories[1].BotResponse, expectedHistories[1].Feedback)
	mock.ExpectQuery("SELECT .*").WillReturnRows(rows)

	repository := NewRepository(dbx)

	// When
	histories, err := repository.GetAll(ctx)

	// Then
	require.NoError(t, err)
	require.NotNil(t, histories)
	require.Equal(t, expectedHistories, histories)
	require.NoError(t, mock.ExpectationsWereMet())
}

func TestRepositoryGetAll_FailsDueToInvalidSelect(t *testing.T) {
	// Given
	db, mock, err := sqlmock.New()
	require.NoError(t, err)

	defer db.Close()

	dbx := sqlx.NewDb(db, "sqlmock")

	ctx := context.Background()

	wrongQuery := regexp.QuoteMeta("SELECT wrong FROM histories;")
	expectedError := errors.New(`Query: could not match actual sql: \"SELECT user_message, bot_response,
										feedback FROM histories;\" with expected regexp \"SELECT wrong FROM histories;\"`)

	expectedHistories := make([]domain.History, 0)
	mock.ExpectQuery(wrongQuery).WillReturnError(expectedError)

	repository := NewRepository(dbx)

	// When
	histories, err := repository.GetAll(ctx)

	// Then
	require.Equal(t, expectedHistories, histories)
	require.ErrorContains(t, err, "Query")
	require.ErrorContains(t, err, "could not match actual sql")
	require.ErrorContains(t, err, "with expected regexp")
}

func TestRepositorySave_Successful(t *testing.T) {
	// Given
	db, mock, err := sqlmock.New()
	require.NoError(t, err)

	defer db.Close()

	dbx := sqlx.NewDb(db, "sqlmock")

	ctx := context.Background()

	expectedHistoryID := 1
	history := domain.History{
		UserMessage: "Test",
		BotResponse: "Test",
		Feedback:    true,
	}

	addRow := sqlmock.NewRows([]string{"id"}).AddRow("1")
	mock.ExpectBegin()
	mock.ExpectPrepare(`INSERT INTO histories`)
	mock.ExpectQuery(`INSERT INTO histories`).WillReturnRows(addRow)
	mock.ExpectCommit()

	repository := NewRepository(dbx)

	// When
	historyID, err := repository.Save(ctx, history)

	// Then
	require.NoError(t, err)
	require.NotNil(t, historyID)
	require.Equal(t, expectedHistoryID, historyID)
	require.NoError(t, mock.ExpectationsWereMet())
}

func TestRepositorySave_FailsDueToInvalidBeginTransaction(t *testing.T) {
	// Given
	db, mock, err := sqlmock.New()
	require.NoError(t, err)

	defer db.Close()

	dbx := sqlx.NewDb(db, "sqlmock")

	ctx := context.Background()

	expectedHistoryID := 0
	history := domain.History{
		UserMessage: "Test",
		BotResponse: "Test",
		Feedback:    true,
	}

	expectedError := errors.New("You have an error in your SQL syntax")

	mock.ExpectBegin().WillReturnError(expectedError)

	repository := NewRepository(dbx)

	// When
	historyID, err := repository.Save(ctx, history)

	// Then
	require.Equal(t, expectedHistoryID, historyID)
	require.ErrorContains(t, err, "You have an error in your SQL syntax")
}

func TestRepositorySave_FailsDueToInvalidPreparation(t *testing.T) {
	// Given
	db, mock, err := sqlmock.New()
	require.NoError(t, err)

	defer db.Close()

	dbx := sqlx.NewDb(db, "sqlmock")

	ctx := context.Background()

	expectedHistoryID := 0
	history := domain.History{
		UserMessage: "Test",
		BotResponse: "Test",
		Feedback:    true,
	}
	wrongQuery := regexp.QuoteMeta(`INSERT INTO histories (user_message, bot_response, feedback) VALUES ();`)
	expectedError := errors.New(`Prepare: could not match actual sql: \"INSERT INTO histories (user_message, 
										bot_response, feedback) VALUES (?, ?, ?, ?);\" with expected 
										regexp \"INSERT INTO histories \\(user_message, bot_response, feedback\\) 
										VALUES \\(\\);\"`)

	mock.ExpectBegin()
	mock.ExpectPrepare(wrongQuery).WillReturnError(expectedError)

	repository := NewRepository(dbx)

	// When
	historyID, err := repository.Save(ctx, history)

	// Then
	require.Equal(t, expectedHistoryID, historyID)
	require.ErrorContains(t, err, "Prepare: could not match actual sql")
	require.ErrorContains(t, err, "with expected regexp")
}

func TestRepositorySave_FailsDueToFailingExec(t *testing.T) {
	// Given
	db, mock, err := sqlmock.New()
	require.NoError(t, err)

	defer db.Close()

	dbx := sqlx.NewDb(db, "sqlmock")

	ctx := context.Background()

	expectedHistoryID := 0
	history := domain.History{
		UserMessage: "Test",
		BotResponse: "Test",
		Feedback:    true,
	}

	expectedError := errors.New("Error Code: 1136. Column count doesn't match value count at row 1")

	mock.ExpectBegin()
	mock.ExpectPrepare(`INSERT INTO histories`)
	mock.ExpectExec(`INSERT INTO histories`).WillReturnError(expectedError)
	mock.ExpectRollback()

	repository := NewRepository(dbx)

	// When
	historyID, err := repository.Save(ctx, history)

	// Then
	require.Equal(t, expectedHistoryID, historyID)
	require.ErrorContains(t, err, "Error Code: 1136")
	require.ErrorContains(t, err, "Column count doesn't match value count at row 1")
}

func TestRepositorySave_FailsDueToFailingExecWithFailingRollback(t *testing.T) {
	// Given
	db, mock, err := sqlmock.New()
	require.NoError(t, err)

	defer db.Close()

	dbx := sqlx.NewDb(db, "sqlmock")

	ctx := context.Background()

	expectedHistoryID := 0
	history := domain.History{
		UserMessage: "Test",
		BotResponse: "Test",
		Feedback:    true,
	}

	expectedExecError := errors.New("Error Code: 1136. Column count doesn't match value count at row 1")
	expectedRollbackError := fmt.Errorf("insert failed: %v, unable to back: %v",
		expectedExecError, "Rollack error")

	mock.ExpectBegin()
	mock.ExpectPrepare(`INSERT INTO histories`)
	mock.ExpectQuery(`INSERT INTO histories`).WillReturnError(expectedExecError)
	mock.ExpectRollback().WillReturnError(expectedRollbackError)

	repository := NewRepository(dbx)

	// When
	historyID, err := repository.Save(ctx, history)

	// Then
	require.Equal(t, expectedHistoryID, historyID)
	require.ErrorContains(t, err, "insert failed")
	require.ErrorContains(t, err, "Error Code: 1136")
	require.ErrorContains(t, err, "Column count doesn't match value count at row 1")
	require.ErrorContains(t, err, "unable to back")
	require.ErrorContains(t, err, "Rollack error")
}

func TestRepositorySave_FailsDueToFailingCommit(t *testing.T) {
	// Given
	db, mock, err := sqlmock.New()
	require.NoError(t, err)

	defer db.Close()

	dbx := sqlx.NewDb(db, "sqlmock")

	ctx := context.Background()

	expectedHistoryID := 0
	history := domain.History{
		UserMessage: "Test",
		BotResponse: "Test",
		Feedback:    true,
	}
	expectedError := errors.New("sql: transaction has already been committed or rolled back")

	addRow := sqlmock.NewRows([]string{"id"}).AddRow("1")
	mock.ExpectBegin()
	mock.ExpectPrepare(`INSERT INTO histories`)
	mock.ExpectQuery(`INSERT INTO histories`).WillReturnRows(addRow)
	mock.ExpectCommit().WillReturnError(expectedError)

	repository := NewRepository(dbx)

	// When
	historyID, err := repository.Save(ctx, history)

	// Then
	require.Equal(t, expectedHistoryID, historyID)
	require.ErrorContains(t, err, "sql")
	require.ErrorContains(t, err, "transaction has already been committed or rolled back")
}
