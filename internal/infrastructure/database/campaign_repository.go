package database

import (
	"SendEmail/internal/campaign"
	"gorm.io/gorm"
)

type CampaignRepository struct {
	Db *gorm.DB
}

// criacao e updated usando o mesmo metodo =O
func (c *CampaignRepository) Create(campaign *campaign.Campaign) error {
	tx := c.Db.Create(campaign) //criacao no banco de dados
	return tx.Error
}

func (c *CampaignRepository) Update(campaign *campaign.Campaign) error {
	tx := c.Db.Save(campaign) //criacao no banco de dados
	return tx.Error
}

func (c *CampaignRepository) Get() ([]campaign.Campaign, error) {
	var campaigns []campaign.Campaign
	tx := c.Db.Find(&campaigns)
	return campaigns, tx.Error
}

func (c *CampaignRepository) GetBy(id string) (*campaign.Campaign, error) {
	var campaign campaign.Campaign
	tx := c.Db.Preload("Contacts").First(&campaign, "id = ?", id) // a quantidade de numero de contattos usando o Preload com join
	return &campaign, tx.Error
}

func (c *CampaignRepository) Delete(campaign *campaign.Campaign) error {

	tx := c.Db.Select("Contacts").Delete(campaign) //criacao no banco de dados
	return tx.Error
}

// testando com a relacao com a data de consulta
func (c *CampaignRepository) GetCampaignsToBeSent() ([]campaign.Campaign, error) {
	var campaigns []campaign.Campaign
	tx := c.Db.Preload("Contacts").Find(
		&campaigns,
		"status = ? and date_part('minute', now()::timestamp - updated_on::timestamp) > ?",
		campaign.Started,
		1)
	return campaigns, tx.Error
}
