package campaign

import (
	"SendEmail/internal/contract"
	"SendEmail/internalErrors"
	"errors"
	"gorm.io/gorm"
)

type ServiceImp struct {
	Repository Repository
	SendMail   func(campaign *Campaign) error
}

// izolando a capana de post
type Service interface {
	Create(newCampaign contract.NewCampaignDto) (string, error)
	GetBy(id string) (*contract.CampaignResponse, error)
	Cancel(id string) error
	Delete(id string) error
	Start(id string) error
}

func (s *ServiceImp) Create(dto contract.NewCampaignDto) (string, error) {
	campaing, err := NewCampaign(dto.Name, dto.Content, dto.Emails)
	if err != nil {
		return "", err
	}

	err = s.Repository.Create(campaing)
	if err != nil {
		return "", internalErrors.ErrInternal
	}

	return campaing.ID, nil
}

func (s *ServiceImp) GetBy(id string) (*contract.CampaignResponse, error) {
	campaing, err := s.Repository.GetBy(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, internalErrors.ProcessErrorToReturn(err)
		}
		return nil, err
	}

	return &contract.CampaignResponse{
		ID:                   campaing.ID,
		Name:                 campaing.Name,
		Content:              campaing.Content,
		Status:               campaing.Status,
		AmountOfEmailsToSend: len(campaing.Contacts),
	}, nil

}

func (s *ServiceImp) Cancel(id string) error {
	campaing, err := s.Repository.GetBy(id)
	if err != nil {
		return internalErrors.ProcessErrorToReturn(err)
	}
	if campaing.Status != Pending {
		return errors.New("Campaing status invalid ")
	}
	campaing.Cancel()

	err = s.Repository.Update(campaing) // inteligente ao sufucinete para fazer o update na campanha
	if err != nil {
		return internalErrors.ErrInternal
	}

	return nil
}

func (s *ServiceImp) Delete(id string) error {
	campaignSaved, err := s.getAndValidateStatusIsPending(id)
	if err != nil {
		return err
	}

	campaignSaved.Delete()

	err = s.Repository.Delete(campaignSaved)
	if err != nil {
		return internalErrors.ErrInternal
	}

	return nil
}

func (s *ServiceImp) Start(id string) error {
	campaignSaved, err := s.getAndValidateStatusIsPending(id)
	if err != nil {
		return err
	}

	if campaignSaved.Status != Pending {
		return errors.New("Campaing status invalid ")
	}

	// go routines
	go func() {
		err := s.SendMail(campaignSaved)
		if err != nil {
			campaignSaved.Fail()
		} else {
			campaignSaved.Done()
		}
		err = s.Repository.Update(campaignSaved)
	}()

	campaignSaved.Started()
	err = s.Repository.Update(campaignSaved)
	if err != nil {
		return internalErrors.ErrInternal
	}
	return nil
}

func (s *ServiceImp) getAndValidateStatusIsPending(id string) (*Campaign, error) {
	campaign, err := s.Repository.GetBy(id)

	if err != nil {
		return nil, internalErrors.ProcessErrorToReturn(err)
	}

	if campaign.Status != Pending {
		return nil, errors.New("Campaign status invalid")
	}
	return campaign, nil
}
