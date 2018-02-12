package database

import (
	"context"
	"database/sql"
	"log"
	"os"

	"fmt"

	"encoding/json"

	"github.com/Sharykhin/gl-mail-api/entity"
	_ "github.com/go-sql-driver/mysql" // dependency of mysql
)

var db *sql.DB

func init() {
	var err error
	dbSource := os.Getenv("MYSQL_SOURCE")
	fmt.Println("MYSQL", dbSource)
	db, err = sql.Open("mysql", dbSource)
	if err != nil {
		log.Fatalf("could not connect to mysql, source: %s, error: %v", dbSource, err)
	}
}

// Create creates a new row in database
func Create(ctx context.Context, m entity.MessageRequest) (*entity.Message, error) {
	p, err := json.Marshal(m.Payload)
	if err != nil {
		return nil, fmt.Errorf("could not marshar payload: %s, err: %v", m, err)
	}
	res, err := db.ExecContext(ctx, "INSERT INTO failed_messages(`action`,`payload`) VALUES(?, ?)", m.Action, p)
	if err != nil {
		return nil, fmt.Errorf("could not create a new failed message: %v", err)
	}

	id, err := res.LastInsertId()
	if err != nil {
		return nil, fmt.Errorf("could not get last insert id: %v", err)
	}

	return &entity.Message{
		ID:      id,
		Action:  m.Action,
		Payload: m.Payload,
	}, nil
}
