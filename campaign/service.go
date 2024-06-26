package campaign

import (
	"errors"
	"fmt"
	"github.com/gosimple/slug"
)

type Service interface {
	GetCampaigns(userID int) ([]Campaign, error)
	GetCampaignByID(input DetailCampaignInput) (Campaign, error)
	CreateCampaign(input CreateCampaignInput) (Campaign, error)
	CampaignUpdate(ID DetailCampaignInput, input CreateCampaignInput) (Campaign, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) GetCampaigns(userID int) ([]Campaign, error) {
	if userID != 0 {
		campaigns, err := s.repository.GetByUserID(userID)
		if err != nil {
			return campaigns, err
		}

		return campaigns, nil
	}

	campaigns, err := s.repository.FindAll()
	if err != nil {
		return campaigns, err
	}
	return campaigns, nil
}

func (s *service) GetCampaignByID(input DetailCampaignInput) (Campaign, error) {
	campaign, err := s.repository.GetByID(input.ID)

	if err != nil {
		return campaign, err
	}
	return campaign, nil
}

func (s *service) CreateCampaign(input CreateCampaignInput) (Campaign, error) {
	campaign := Campaign{}
	campaign.Name = input.Name
	campaign.ShortDescription = input.ShortDescription
	campaign.Description = input.Description
	campaign.GoalAmount = input.GoalAmount
	campaign.Perks = input.Perks
	campaign.UserID = input.User.ID
	stringSlug := fmt.Sprintf("%s %d", input.Name, input.User.ID)
	campaign.Slug = slug.Make(stringSlug)
	newCampaign, err := s.repository.Save(campaign)
	if err != nil {
		return newCampaign, err
	}

	return newCampaign, nil
}

func (s *service) CampaignUpdate(inputID DetailCampaignInput, input CreateCampaignInput) (Campaign, error) {

	campaign, err := s.repository.GetByID(inputID.ID)

	if err != nil {
		return campaign, err
	}

	if campaign.UserID != input.User.ID {
		return campaign, errors.New("Not owner of thi campaign")
	}

	campaign.Name = input.Name
	campaign.ShortDescription = input.ShortDescription
	campaign.Description = input.Description
	campaign.Perks = input.Perks
	campaign.GoalAmount = input.GoalAmount

	updateCampaign, err := s.repository.Update(campaign)

	if err != nil {
		return campaign, err
	}

	return updateCampaign, nil
}
