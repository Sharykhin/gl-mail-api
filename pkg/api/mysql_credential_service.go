package api

import (
	"database/sql"
	"log"
)

type mysqlCredentialService struct {
	db *sql.DB
}

func (ss *mysqlCredentialService) GetCredentials(apiKey string, secretKey string) (*CredentialModel, error) {
	var credential CredentialModel
	row := ss.db.QueryRow("SELECT id, api_key, secret_key, created_at FROM credentials WHERE api_key=? AND secret_key=? LIMIT 1", apiKey, secretKey)
	switch err := row.Scan(&credential.Id, &credential.ApiKet, &credential.SecretKey, &credential.CreatedAt); err {
	case sql.ErrNoRows:
		log.Println(err)
		return &credential, err
	case nil:
	default:
		panic(err)
	}

	return &credential, nil
}

func NewMysqlCredentialService(db *sql.DB) Credentialer {
	return &mysqlCredentialService{db: db}
}