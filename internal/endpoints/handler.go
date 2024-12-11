package endpoints

import "SendEmail/internal/campaign"

type Handler struct {
	CampaignService campaign.Service
}
