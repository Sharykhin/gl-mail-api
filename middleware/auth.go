package middleware

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/Sharykhin/gl-mail-api/util"
)


//JWTAuth middleware checks whether the jwt token was passed through Authorization header
func JWTAuth(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			util.SendResponse(util.Response{
				Success: false,
				Data:    nil,
				Error:   "Authorization header was not provided",
			}, w, http.StatusUnauthorized)
			return
		}

		isBearerAuth := strings.HasPrefix(authHeader, "Bearer ")

		if isBearerAuth {
			tokenString := authHeader[len("Bearer "):]
			publicKey, err := ioutil.ReadFile("public.pem")
			if err != nil {
				log.Fatalf("Could not read public.pem file: %v", err)
			}

			publicRSA, err := jwt.ParseRSAPublicKeyFromPEM(publicKey)
			if err != nil {
				log.Fatalf("could not parse public key: %s. %v", publicKey, err)
			}

			token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
				if _, ok := t.Method.(*jwt.SigningMethodRSA); !ok {
					return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
				}

				return publicRSA, err
			})

			if err != nil {
				util.SendResponse(util.Response{
					Success: false,
					Data:    nil,
					Error:   err.Error(),
				}, w, http.StatusUnauthorized)
				return
			}

			if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
				if role, ok := claims["role"]; ok {
					if role == "admin" {
						h.ServeHTTP(w, r)
						return
					} else {
						util.SendResponse(util.Response{
							Success: false,
							Data:    nil,
							Error:   http.StatusText(http.StatusForbidden),
						}, w, http.StatusForbidden)
						return
					}
				} else {
					util.SendResponse(util.Response{
						Success: false,
						Data:    nil,
						Error:   "Payload does not contain role attribute",
					}, w, http.StatusUnauthorized)
					return
				}
			} else {
				util.SendResponse(util.Response{
					Success: false,
					Data:    nil,
					Error:   http.StatusText(http.StatusUnauthorized),
				}, w, http.StatusUnauthorized)
				return
			}
		} else {
			util.SendResponse(util.Response{
				Success: false,
				Data:    nil,
				Error:   "Authorization header must have Bearer type",
			}, w, http.StatusUnauthorized)
			return
		}
	})
}
