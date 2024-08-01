package history

import (
	"backend/internal/domain"
	"context"
	"github.com/jmoiron/sqlx"
	"log"
)

// Queries.
const (
	_getAllHistoriesStmt = `SELECT id, user_message, bot_response, feedback FROM histories;`
	_saveHistoryStmt     = `INSERT INTO 
    						histories (user_message, bot_response, feedback)
							VALUES ($1, $2, $3)
							RETURNING id;`
)

type Repository interface {
	// GetAll obtain all users from the database.
	GetAll(ctx context.Context) ([]domain.History, error)

	// Save a history to the database.
	Save(ctx context.Context, history domain.History) (int, error)
}

type repository struct {
	conn *sqlx.DB
}

func NewRepository(conn *sqlx.DB) Repository {
	return &repository{conn: conn}
}

func (r *repository) GetAll(ctx context.Context) ([]domain.History, error) {
	histories := make([]domain.History, 0)

	if err := r.conn.SelectContext(ctx, &histories, _getAllHistoriesStmt); err != nil {
		return make([]domain.History, 0), err
	}

	return histories, nil
}

func (r *repository) Save(ctx context.Context, history domain.History) (int, error) {
	tx, err := r.conn.Beginx()
	if err != nil {
		return 0, err
	}

	stmt, err := tx.PreparexContext(ctx, _saveHistoryStmt)
	if err != nil {
		return 0, err
	}

	defer func() {
		err = stmt.Close()
		if err != nil {
			log.Println(err)
		}
	}()

	var id int
	err = stmt.QueryRowxContext(ctx, history.UserMessage, history.BotResponse, history.Feedback).Scan(&id)
	if err != nil {
		if rollbackErr := tx.Rollback(); rollbackErr != nil {
			return 0, rollbackErr
		}

		return 0, err
	}

	if err = tx.Commit(); err != nil {
		return 0, err
	}

	return id, nil
}
