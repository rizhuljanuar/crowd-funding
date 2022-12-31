package transaction

import "crowd-funding/user"

type GetCampaignTransactionsInput struct {
	ID int `uri:"id" binding:"required"`
	User user.User
}