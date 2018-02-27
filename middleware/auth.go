package middleware

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/Sharykhin/gl-mail-api/util"
	"github.com/dgrijalva/jwt-go"
)

//JWTAuth middleware checks whether the jwt token was passed through Authorization header
func JWTAuth(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		env := os.Getenv("APP_ENV")
		if env == "test" {
			h.ServeHTTP(w, r)
			return
		}
		authHeader := r.Header.Get("Authorization")
		isBearerAuth := strings.HasPrefix(authHeader, "Bearer ")

		if isBearerAuth {
			tokenString := authHeader[len("Bearer "):]
			keyFile := os.Getenv("JWT_PUBLIC_KEY")

			publicKey, err := ioutil.ReadFile(keyFile)
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
					Error:   fmt.Errorf("token is invalid: %s", err),
				}, w, http.StatusUnauthorized)
				return
			}

			if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
				if role, ok := claims["role"]; ok {
					if role == "admin" {
						h.ServeHTTP(w, r)
						return
					}

					util.SendResponse(util.Response{
						Success: false,
						Data:    nil,
						Error:   fmt.Errorf(http.StatusText(http.StatusForbidden)),
					}, w, http.StatusForbidden)
					return

				}
				util.SendResponse(util.Response{
					Success: false,
					Data:    nil,
					Error:   fmt.Errorf(http.StatusText(http.StatusForbidden)),
				}, w, http.StatusForbidden)
				return
			}
			util.SendResponse(util.Response{
				Success: false,
				Data:    nil,
				Error:   fmt.Errorf(http.StatusText(http.StatusUnauthorized)),
			}, w, http.StatusUnauthorized)
			return
		}

		util.SendResponse(util.Response{
			Success: false,
			Data:    nil,
			Error:   fmt.Errorf(http.StatusText(http.StatusUnauthorized)),
		}, w, http.StatusUnauthorized)
	})
}
