package main

import (
	"SendEmail/internal/campaign"
	"SendEmail/internal/endpoints"
	"SendEmail/internal/infrastructure/database"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
)

func main() {

	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	db := database.NewDb() // instanciando a conexao com o banco de dados

	service := campaign.ServiceImp{
		Repository: &database.CampaignRepository{Db: db},
	}

	handler := endpoints.Handler{
		CampaignService: &service,
	}
	// ROTAS AUTENTICADAS
	r.Route("/campaigns", func(r chi.Router) {
		r.Use(endpoints.Auth)
		//POST
		r.Post("/", endpoints.HandlerError(handler.CampaignsPost))
		//GET
		r.Get("/{id}", endpoints.HandlerError(handler.CampaignsGetById))
		// DELETE
		r.Delete("/delete/{id}", endpoints.HandlerError(handler.CampaignsDelete))
		// UPDATED
		r.Patch("/cancel/{id}", endpoints.HandlerError(handler.CampaignsCancelPath))
	})

	http.ListenAndServe(":3000", r)
}
