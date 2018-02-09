package middleware

import (
	"fmt"
	"io/ioutil"
	"log"
	myjwt "github.com/Sharykhin/gl-mail-api/pkg/jwt"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/Sharykhin/gl-mail-api/service"
)

const ADMIN_ROLE = "admin"
const PUBLICKEY = "public.pem"

func JWTAuth(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if r := recover(); r != nil {
				log.Printf("Catch panic. Error: %s", r)
				service.SendResponse(service.Response{
					Success: false,
					Data:    nil,
					Error:   http.StatusText(http.StatusUnauthorized),
				}, w, http.StatusUnauthorized)
				return
			}
		}()

		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			service.SendResponse(service.Response{
				Success: false,
				Data:    nil,
				Error:   http.StatusText(http.StatusUnauthorized),
			}, w, http.StatusUnauthorized)
			return
		}

		isMailAuth := strings.HasPrefix(authHeader, "Mail ")
		isBearerAuth := strings.HasPrefix(authHeader, "Bearer ")

		if isMailAuth {
			token := authHeader[len("Mail "):]

			payload, err := myjwt.Decode(token, myjwt.SECRET)
			if err != nil {
				service.SendResponse(service.Response{
					Success: false,
					Data:    nil,
					Error:   http.StatusText(http.StatusUnauthorized),
				}, w, http.StatusUnauthorized)
				return
			}

			id := payload.(myjwt.Payload).Public.(map[string]interface{})["id"]
			if id != nil && id != "" {
				h.ServeHTTP(w, r)
				return
			} else {
				service.SendResponse(service.Response{
					Success: false,
					Data:    nil,
					Error:   http.StatusText(http.StatusUnauthorized),
				}, w, http.StatusUnauthorized)
				return
			}
		}

		if isBearerAuth {
			tokenString := authHeader[len("Bearer "):]
			publicKey, err := ioutil.ReadFile(PUBLICKEY)
			if err != nil {
				log.Panic(err)
			}

			publicRSA, err := jwt.ParseRSAPublicKeyFromPEM(publicKey)
			if err != nil {
				log.Println(err)
				service.SendResponse(service.Response{
					Success: false,
					Data:    nil,
					Error:   http.StatusText(http.StatusInternalServerError),
				}, w, http.StatusInternalServerError)
				return
			}

			token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
				if _, ok := t.Method.(*jwt.SigningMethodRSA); !ok {
					return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
				}

				return publicRSA, err
			})

			if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
				log.Println(claims)
				role := claims["role"]
				if role == ADMIN_ROLE {
					h.ServeHTTP(w, r)
					return
				} else {
					service.SendResponse(service.Response{
						Success: false,
						Data:    nil,
						Error:   http.StatusText(http.StatusForbidden),
					}, w, http.StatusForbidden)
					return
				}
			} else {
				service.SendResponse(service.Response{
					Success: false,
					Data:    nil,
					Error:   http.StatusText(http.StatusUnauthorized),
				}, w, http.StatusUnauthorized)
				return
			}
		}
	})
}
