package campaign

import (
	"github.com/jaswdr/faker/v2"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

// variavel global para as constats do teste
var (
	name     = "Camping x"
	content  = "body hi!"
	contacts = []string{"emai@1.com", "emai@2.com"}
)

// testing -> Ele ajuda testar no  go
func Test_NewCampaign_CreatedNewCampaign(t *testing.T) {

	// organiza eles para fazer os testes, atribuicao
	assert := assert.New(t)

	// teste aquie os atributos que foram mokcados
	campaign, _ := NewCampaign(name, content, contacts)

	//assert ou seja verificacao para fazer o teste
	assert.Equal(campaign.Name, name)
	assert.Equal(campaign.Content, content)
	assert.Equal(len(campaign.Contacts), len(contacts))

}

func Test_NewCampaign_Id_notNill(t *testing.T) {
	assert := assert.New(t)

	campaign, _ := NewCampaign(name, content, contacts)

	assert.NotNil(campaign.ID)
}

// data tem que se de agora
func Test_NewCampaign_createdOnMustBeNow(t *testing.T) {
	assert := assert.New(t)

	now := time.Now().Add(-time.Minute)

	campaign, _ := NewCampaign(name, content, contacts)

	assert.Greater(campaign.CreatedOn, now)
}

func Test_NewCampaign_MusValidateNameMin(t *testing.T) {
	assert := assert.New(t)

	_, err := NewCampaign("", content, contacts)

	assert.Equal("name is required with min 5", err.Error())
}

func Test_NewCampaign_MusValidateNameMax(t *testing.T) {
	assert := assert.New(t)
	fake := faker.New()
	_, err := NewCampaign(fake.Lorem().Text(30), content, contacts)

	assert.Equal("name is required with max 24", err.Error())
}

func Test_NewCampaign_MusValidateContent(t *testing.T) {
	assert := assert.New(t)
	fake := faker.New()
	_, err := NewCampaign(name, fake.Lorem().Text(2104), contacts)

	assert.Equal("content is required with max 1024", err.Error())
}

func Test_NewCampaign_MusValidateContacts(t *testing.T) {
	assert := assert.New(t)

	_, err := NewCampaign(name, content, []string{"email_invalid"})

	assert.Equal("emails is email_invalid", err.Error())
}

func Test_NewCampaign_MustSatusWithPending(t *testing.T) {
	assert := assert.New(t)

	campaign, _ := NewCampaign(name, content, contacts)

	assert.Equal(Pending, campaign.ID, campaign.Status)
}
