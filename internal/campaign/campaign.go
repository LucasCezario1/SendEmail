package campaign

import (
	"SendEmail/internalErrors"
	"github.com/rs/xid"
	"time"
)

// Constates  para o status do usuario
const (
	Pending  string = "pending"
	Started  string = "started"
	Canceled string = "canceled"
	Deleted  string = "deleted"
	Done     string = "done"
)

type Contact struct {
	ID         string `gorm:"size:50" `
	Email      string `validate:"email" gorm:"size:50"`
	CampaignID string `gorm:"size:50"`
}

type Campaign struct {
	ID        string    `validate:"required" gorm:"size:50"`
	Name      string    `validate:"min=5,max=24" gorm:"size:100"`
	CreatedOn time.Time `validate:"required"`
	Content   string    `validate:"min=5,max=1024" gorm:"size:1024"`
	Contacts  []Contact `validate:"min=1,dive"`
	Status    string    `gorm:"size:20"`
}

func (c *Campaign) Cancel() {
	c.Status = "canceled"
}

func (c *Campaign) Delete() {
	c.Status = "deleted"
}

// feita para criar uma nova campanha
func NewCampaign(name string, content string, emails []string) (*Campaign, error) {

	//contacts ele vira o emails
	contacts := make([]Contact, len(emails))
	for index, email := range emails {
		contacts[index].Email = email
		contacts[index].ID = xid.New().String()
	}
	campaign := &Campaign{
		ID:        xid.New().String(),
		Name:      name,
		Content:   content,
		CreatedOn: time.Now(),
		Contacts:  contacts,
		Status:    Pending,
	}
	err := internalErrors.ValidateStruct(campaign)
	if err != nil {
		return nil, err
	}
	return campaign, err

}
