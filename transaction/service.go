package transaction

import (
	"crowd-funding/campaign"
	"errors"
)

type Service interface {
	GetTransactionByCampaignID(input GetCampaignTransactionsInput) ([]Transaction, error)
}

type service struct {
	repository Repository
	campaignRepositoy campaign.Repository
}

func NewService (repository Repository, campaignRepository campaign.Repository) *service {
	return &service{repository, campaignRepository}
}

func (s *service) GetTransactionByCampaignID(input GetCampaignTransactionsInput) ([]Transaction, error)  {
	// get campaign
	// check campaign.userID != user_id yg melakukan request

	campaign, err := s.campaignRepositoy.FindByID(input.ID)
	if err != nil {
		return []Transaction{}, err
	}

	if campaign.UserID != input.User.ID {
		return []Transaction{}, errors.New("Not owner of the campaign")
	}


	transactions, err := s.repository.GetByCampaignID(input.ID)

	if err != nil {
		return transactions, err
	}

	return transactions, nil
}