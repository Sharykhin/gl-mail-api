package api

type CredentialBody struct {
	ApiKey string 			`json:"api_key"`
	SecretKey string		`json:"secret_key"`
}
