package endpoints

import (
	"SendEmail/internal/contract"
	"github.com/go-chi/render"
	"net/http"
)

func (h *Handler) CampaignsPost(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {

	var request contract.NewCampaignRequest
	render.DecodeJSON(r.Body, &request)
	email := r.Context().Value("email").(string)
	request.CreatedBy = email
	id, err := h.CampaignService.Create(request)
	return map[string]string{"id": id}, 201, err // novo retorno, com um mapa de string
}
