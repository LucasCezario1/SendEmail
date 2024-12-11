package endpoints

import (
	"github.com/go-chi/chi/v5"
	"net/http"
)

func (h *Handler) CampaignsGetById(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {

	id := chi.URLParam(r, "id")
	campaigns, err := h.CampaignService.GetBy(id)

	return campaigns, 200, err // novo retorno
}
