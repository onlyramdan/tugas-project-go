package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"tugas/campaign"
	"tugas/helper"
	"tugas/user"
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
	campaignDetail, err := h.campaignService.GetCampaignByID(input)
	if err != nil {
		response := helper.APIresponse("Failed to Get Campaign Detail ", http.StatusBadRequest, "Error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIresponse("Detail Of Campaign", http.StatusOK, "Success", campaign.FormatCampaignDetail(campaignDetail))
	c.JSON(http.StatusOK, response)
}

func (h *campagnHandler) CampaignCreate(c *gin.Context) {

	var input campaign.CreateCampaignInput

	err := c.ShouldBindJSON(&input)

	if err != nil {
		errors := helper.FormatError(err)
		msgError := gin.H{"errors": errors}
		response := helper.APIresponse("Failed to Create Campaign", http.StatusBadRequest, "Error", msgError)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	currentUser := c.MustGet("currentUser").(user.User)
	input.User = currentUser

	newCampaign, err := h.campaignService.CreateCampaign(input)

	if err != nil {
		response := helper.APIresponse("Failed to Create Campaign", http.StatusBadRequest, "Error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIresponse("Success Create Campaign", http.StatusOK, "Success", campaign.CampaignFormat(newCampaign))
	c.JSON(http.StatusOK, response)
}

func (h *campagnHandler) UpdateCampaign(c *gin.Context) {
	var ID campaign.DetailCampaignInput

	err := c.ShouldBindUri(&ID)

	if err != nil {
		errors := helper.FormatError(err)
		msgError := gin.H{"errors": errors}
		response := helper.APIresponse("Failed to Update Campaign", http.StatusBadRequest, "Error", msgError)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	var input campaign.CreateCampaignInput

	err = c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatError(err)
		msgError := gin.H{"errors": errors}
		response := helper.APIresponse("Failed to Update Campaign", http.StatusBadRequest, "Error", msgError)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	currentUser := c.MustGet("currentUser").(user.User)
	input.User = currentUser

	updateCampaign, err := h.campaignService.CampaignUpdate(ID, input)

	if err != nil {
		response := helper.APIresponse("Failed to Update Campaign", http.StatusBadRequest, "Error", err)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIresponse("Success Update Campaign", http.StatusOK, "Success", campaign.CampaignFormat(updateCampaign))
	c.JSON(http.StatusOK, response)
}
