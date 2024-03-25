package campaign

type CampaignFormatter struct {
	ID               int    `json:"id"`
	UserID           int    `json:"user_id"`
	Name             string `json:"name"`
	ShortDescription string `json:"short_description"`
	ImageURL         string `json:"image_url"`
	GoalAmount       int    `json:"goal_amount"`
	CurrentAmount    int    `json:"current_amount"`
	Slug             string `json:"slug"`
}

func CampaignFormat(campaign Campaign) CampaignFormatter {
	campaignsFormatter := CampaignFormatter{}
	campaignsFormatter.ID = campaign.ID
	campaignsFormatter.UserID = campaign.UserID
	campaignsFormatter.Name = campaign.Name
	campaignsFormatter.ShortDescription = campaign.ShortDescription
	campaignsFormatter.GoalAmount = campaign.GoalAmount
	campaignsFormatter.CurrentAmount = campaign.CurrentAmount
	campaignsFormatter.Slug = campaign.Slug
	campaignsFormatter.ImageURL = ""

	if (len(campaign.CampaignImages)) > 0 {
		campaignsFormatter.ImageURL = campaign.CampaignImages[0].FileName
	}

	return campaignsFormatter
}
func FormatCampaign(campaigns []Campaign) []CampaignFormatter {

	campaignsFormatter := []CampaignFormatter{}

	for _, campaign := range campaigns {
		campaignFormatter := CampaignFormat(campaign)
		campaignsFormatter = append(campaignsFormatter, campaignFormatter)
	}

	return campaignsFormatter
}
