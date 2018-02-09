package api

type ApiService struct {
	Credential Credentialer
}

func NewApiService() *ApiService {
	var cs = NewMysqlCredentialService(getStorage())
	var ms = &ApiService{Credential:cs}
	return ms
}
