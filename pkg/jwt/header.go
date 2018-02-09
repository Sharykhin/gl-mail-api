package jwt

type Header struct {
	Alg string `json:"alg"`
	Typ string `json:"typ"`
}
