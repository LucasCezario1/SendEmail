package endpoints

import (
	"context"
	"github.com/coreos/go-oidc/v3/oidc"
	jwtgo "github.com/dgrijalva/jwt-go/v4"
	"github.com/go-chi/render"
	"net/http"
	"strings"
)

// analisa todos os endpoiuts globomaneete
func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		tokenString := r.Header.Get("Authorization")
		if tokenString == "" {
			render.JSON(w, r, map[string]string{"ERROR": "request does  not contain an authorization token"})
		}

		// tirar o Bearer do token
		tokenString = strings.Replace(tokenString, "Bearer", "", 1)

		provider, err := oidc.NewProvider(r.Context(), "http://localhost:8080/realms/provider")
		if err != nil {
			render.Status(r, http.StatusInternalServerError)
			render.JSON(w, r, map[string]string{"ERROR": "error to connect to the provider"})
			return
		}

		verifier := provider.Verifier(&oidc.Config{ClientID: "email"})
		_, err = verifier.Verify(r.Context(), tokenString)
		if err != nil {
			render.Status(r, http.StatusUnauthorized)
			render.JSON(w, r, map[string]string{"ERROR": "error to verify the token"})
			return
		}

		token, _ := jwtgo.Parse(tokenString, nil)
		claims := token.Claims.(jwtgo.MapClaims)
		email := claims["email"]

		// cotexnto  da vailidacao
		ctx := context.WithValue(r.Context(), "email", email)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
