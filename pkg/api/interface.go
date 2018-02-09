package api

type Credentialer interface {
	GetCredentials(apiKey string, secretKey string) (*CredentialModel, error)
}
