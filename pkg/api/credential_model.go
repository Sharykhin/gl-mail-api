package api

type CredentialModel struct {
	Id int				`json:"id"`
	ApiKet string		`json:"api_key"`
	SecretKey string	`json:"secret_key"`
	CreatedAt string	`json:"created_at"`
}
