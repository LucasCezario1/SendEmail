package main

import (
	"SendEmail/internal/campaign"
	"SendEmail/internal/infrastructure/database"
	"SendEmail/internal/infrastructure/mail"
)

func main() {
	println("Started worker")

	db := database.NewDb() // instanciando a conexao com o banco de dados
	repository := database.CampaignRepository{Db: db}
	// injesao de depedecias
	service := campaign.ServiceImp{
		Repository: &database.CampaignRepository{Db: db},
		SendMail:   mail.SendMail,
	}

	for {
		campaings, _ := repository.GetCampaignsToBeSent()

		for _, campaign := range campaings {
			service.SendEmailAndUpdatedStatus(&campaign)
		}
	}

}
