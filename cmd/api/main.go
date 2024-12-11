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

	//POST para campanha
	r.Post("/campaigns", endpoints.HandlerError(handler.CampaignsPost))

	//GET
	r.Get("/campaigns/{id}", endpoints.HandlerError(handler.CampaignsGetById))

	// UPDATED
	r.Patch("/campaigns/cancel/{id}", endpoints.HandlerError(handler.CampaignsCancelPath))

	// DELETE
	r.Delete("/campaigns/delete/{id}", endpoints.HandlerError(handler.CampaignsDelete))

	http.ListenAndServe(":8080", r)
}
