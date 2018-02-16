package database

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/Sharykhin/gl-mail-api/entity"
	_ "github.com/go-sql-driver/mysql" // dependency of mysql
)

var db *sql.DB

// Storage keeps reference to a database source
var Storage storage

// TODO: experimental case how to implement it
type storage struct {
	db *sql.DB
}

func (s storage) Create(ctx context.Context, mr entity.MessageRequest) (*entity.Message, error) {
	return Create(ctx, mr)
}

func (s storage) GetList(ctx context.Context, limit, offset int) ([]entity.Message, error) {
	return GetList(ctx, limit, offset)
}

func (s storage) Count(ctx context.Context) (int, error) {
	return Count(ctx)
}

func init() {
	var err error
	dbSource := os.Getenv("MYSQL_SOURCE")
	db, err = sql.Open("mysql", dbSource)
	if err != nil {
		log.Fatalf("could not connect to mysql, source: %s, error: %v", dbSource, err)
	}

	Storage = storage{db: db}
}

// Create creates a new record of failed mail
func Create(ctx context.Context, mr entity.MessageRequest) (*entity.Message, error) {
	p, err := json.Marshal(mr.Payload)
	if err != nil {
		return nil, fmt.Errorf("could not marshar payload: %s, err: %v", mr, err)
	}

	res, err := db.ExecContext(ctx, "INSERT INTO failed_mails(`action`, `payload`, `reason`, `created_at`) VALUES(?, ?, ?, NOW())", mr.Action, p, mr.Reason)
	if err != nil {
		return nil, fmt.Errorf("could not create a new failed message: %v", err)
	}

	id, err := res.LastInsertId()
	if err != nil {
		return nil, fmt.Errorf("could not get last insert id: %v", err)
	}

	return &entity.Message{
		ID:        id,
		Action:    mr.Action,
		Payload:   mr.Payload,
		Reason:    mr.Reason,
		CreatedAt: time.Now(),
	}, nil
}

// GetList returns limited number of rows
func GetList(ctx context.Context, limit, offset int) ([]entity.Message, error) {
	rows, err := db.QueryContext(ctx, "SELECT `id`, `action`, `reason` FROM failed_mails LIMIT ? OFFSET ?", limit, offset)
	if err != nil {
		return nil, fmt.Errorf("could not make a select statement: %v", err)
	}
	defer rows.Close() // nolint: errcheck

	var messages []entity.Message
	for rows.Next() {
		var m entity.Message
		err := rows.Scan(&m.ID, &m.Action, &m.Reason)
		if err != nil {
			return nil, fmt.Errorf("could not scan a row to struct %v: %v", m, err)
		}
		messages = append(messages, m)
	}

	return messages, rows.Err()
}

// Count returns number of all rows
func Count(ctx context.Context) (int, error) {
	var count int
	row := db.QueryRowContext(ctx, "SELECT COUND(id) AS `total` FROM failed_mails")
	err := row.Scan(&count)
	if err != nil {
		return 0, fmt.Errorf("could not make select statement: %v", err)
	}
	return count, nil
}
