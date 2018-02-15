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
