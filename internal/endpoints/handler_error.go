package endpoints

import (
	"SendEmail/internalErrors"
	"errors"
	"github.com/go-chi/render"
	"gorm.io/gorm"
	"net/http"
)

// ASSINATURA DO METODO
type EndpointFunc func(w http.ResponseWriter, r *http.Request) (interface{}, int, error)

// analisa todos os endpoiuts globomanete
func HandlerError(handler EndpointFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		obj, statusCode, err := handler(w, r)
		if err != nil {
			if errors.Is(err, internalErrors.ErrInternal) {
				render.Status(r, http.StatusInternalServerError)
			} else if errors.Is(err, gorm.ErrRecordNotFound) {
				render.Status(r, http.StatusNotFound)
			} else {
				render.Status(r, http.StatusBadRequest)
			}
			render.JSON(w, r, map[string]string{"error": err.Error()})
			return
		}
		render.Status(r, statusCode)
		// Sempre renderize JSON, mesmo que obj seja nil
		render.JSON(w, r, obj)
	})
}
