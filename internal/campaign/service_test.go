package campaign

import (
	"SendEmail/internal/contract"
	"SendEmail/internalErrors"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)


type repositoryMock struct {
	mock.Mock
}

func (r *repositoryMock) Save(campaign *Campaign) error {
	args := r.Called(campaign)
	return args.Error(0)
}
func (r *repositoryMock) Get() ([]Campaign, error) {
	//args := r.Called(campaign)
	return nil, nil
}

func (r *repositoryMock) GetBy(id string) (*Campaign, error) {
	args, := r.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*Campaign), args.Error(1) //cast
}

var (
	newCampaing = contract.NewCampaignRequest{
		Name:    "teste y",
		Content: "body hi!",
		Emails:  []string{"lucasdsc96@gmail.com"},
	}

	service = ServiceImp{}
)

func Test_Create_Campaing(t *testing.T) {
	assert := assert.New(t)

	repositoryMock := new(repositoryMock)
	repositoryMock.On("Save", mock.Anything).Return(nil)
	service.Repository = repositoryMock
	id, err := service.Create(newCampaing)

	assert.NotNil(id)
	assert.Nil(err)
}

func Test_Create_ValidateDomainError(t *testing.T) {
	assert := assert.New(t)

	_, err := service.Create(newCampaing)

	assert.False(errors.Is(internalErrors.ErrInternal, err))
}

func Test_Save_Campaing(t *testing.T) {

	repositoryMock := new(repositoryMock)
	repositoryMock.On("Save", mock.MatchedBy(func(campaign *Campaign) bool {
		if campaign.Name != newCampaing.Name ||
			campaign.Content != newCampaing.Content ||
			len(campaign.Contacts) != len(newCampaing.Emails) {
			return false
		}
		return true
	})).Return(nil)
	service.Repository = repositoryMock{}
	service.Create(newCampaing)

	repositoryMock.AssertExpectations(t)
}

func Test_Save_ValidateRepositorySave(t *testing.T) {
	assert := assert.New(t)
	repositoryMock := new(repositoryMock)
	repositoryMock.On("Save", mock.Anything).Return(errors.New("Error to save database"))
	service.Repository = repositoryMock
	_, err := service.Create(newCampaing)

	assert.True(errors.Is(internalErrors.ErrInternal, err))

}

func GetById_ReturnCamping(t *testing.T) {
	assert := assert.New(t)
	campaign, _ := NewCampaign(newCampaing.Name, newCampaing.Content, newCampaing.Emails)
	repositoryMock := new(repositoryMock)
	service.Repository = repositoryMock
	repositoryMock.On("GetById", mock.Anything).Return("Error to save on databasee", nil)

	campaignReturned, _ := service.GetBy(campaign.ID)
	assert.Equal(campaign.ID, campaignReturned.ID)
}
