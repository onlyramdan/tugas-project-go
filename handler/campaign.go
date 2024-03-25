package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"tugas/campaign"
	"tugas/helper"
)

type campagnHandler struct {
	campaignService campaign.Service
}

func NewCampagnHandler(campaignService campaign.Service) *campagnHandler {
	return &campagnHandler{campaignService}
}

func (h *campagnHandler) GetCampaigns(c *gin.Context) {
	userID, _ := strconv.Atoi(c.Query("user_id"))

	campaigns, err := h.campaignService.GetCampaigns(userID)

	if err != nil {
		response := helper.APIresponse("Failed to Get Campaigns", http.StatusBadRequest, "Failed", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.APIresponse("List Of Campaign", http.StatusOK, "Success", campaign.FormatCampaign(campaigns))
	c.JSON(http.StatusOK, response)

}

func (h *campagnHandler) GetCampaign(c *gin.Context) {
	var input campaign.DetailCampaignInput

	err := c.ShouldBindUri(&input)
	if err != nil {
		response := helper.APIresponse("Failed to Get Campaign Detail ", http.StatusBadRequest, "Error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	campaign, err := h.campaignService.GetCampaignByID(input)
	if err != nil {
		response := helper.APIresponse("Failed to Get Campaign Detail ", http.StatusBadRequest, "Error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIresponse("Detail Of Campaign", http.StatusOK, "Success", campaign)
	c.JSON(http.StatusOK, response)
}
