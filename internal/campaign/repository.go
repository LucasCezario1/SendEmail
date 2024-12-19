package campaign

type Repository interface {
	Create(campaign *Campaign) error
	Get() ([]Campaign, error)
	GetBy(id string) (*Campaign, error)
	Delete(campaing *Campaign) error
	Update(campaing *Campaign) error
	GetCampaignsToBeSent() ([]Campaign, error)
}
