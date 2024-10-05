package controllers

import (
    "mutualfund/models"
    "mutualfund/services"
    "github.com/gin-gonic/gin"
    "net/http"
    "strconv"
)

type FundSchemeController struct {
    fundSchemeService *services.FundSchemeService
}

func NewFundSchemeController(fundSchemeService *services.FundSchemeService) *FundSchemeController{
	return &FundSchemeController{
		fundSchemeService : fundSchemeService,
	}
}

// CreateFundScheme handles creating a new fund scheme
func (ctrl FundSchemeController) CreateFundScheme(c *gin.Context) {
    var scheme models.FundScheme
    if err := c.ShouldBindJSON(&scheme); err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
        return
    }
    
    response,responseErr := ctrl.fundSchemeService.CreateFundScheme(&scheme)
    if responseErr != nil {
        c.AbortWithStatusJSON(responseErr.Status, responseErr)
        return
    }
    c.JSON(http.StatusOK, response)
}

// GetFundSchemes handles retrieving all fund schemes
func (ctrl FundSchemeController) GetFundSchemes(c *gin.Context) {
    schemes, err := ctrl.fundSchemeService.GetFundSchemes()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch fund schemes"})
        return
    }
    c.JSON(http.StatusOK, gin.H{"data": schemes})
}

// GetFundSchemeByID handles retrieving a fund scheme by ID
func (ctrl FundSchemeController) GetFundSchemeByID(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "invalid scheme ID"})
        return
    }
    response, responseErr := ctrl.fundSchemeService.GetFundSchemeByID(id)
    if responseErr != nil {
        c.AbortWithStatusJSON(responseErr.Status, responseErr)
        return
    }
    c.JSON(http.StatusOK, response)
}



// UpdateFundScheme handles updating an existing fund scheme
func (ctrl FundSchemeController) UpdateFundScheme(c *gin.Context) {
    _, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "invalid scheme ID"})
        return
    }
    var scheme models.FundScheme
    if err := c.ShouldBindJSON(&scheme); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    scheme.SchemeID = c.Param("id")
    if err := ctrl.fundSchemeService.UpdateFundScheme(&scheme); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update fund scheme"})
        return
    }
    c.JSON(http.StatusOK, gin.H{"message": "fund scheme updated successfully"})
}

// DeleteFundScheme handles deleting a fund scheme
func (ctrl FundSchemeController) DeleteFundScheme(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "invalid scheme ID"})
        return
    }
    if err := ctrl.fundSchemeService.DeleteFundScheme(id); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Message})
        return
    }
    c.JSON(http.StatusOK, gin.H{"message": "fund scheme deleted successfully"})
}

// GetFilteredFundSchemes retrieves fund schemes based on filters
func (ctrl *FundSchemeController) GetFilteredFundSchemes(c *gin.Context) {
    // Retrieve query parameters
    fundHouse := c.Query("fund_house")
    category := c.Query("category")
    
    minReturns, err := strconv.ParseFloat(c.Query("returns_pa"), 64)
    if err != nil {
        minReturns = 0 // Default value
    }

    minInvestment, err := strconv.ParseFloat(c.Query("min_investment"), 64)
    if err != nil {
        minInvestment = 0 // Default value
    }

    // Call service method
    schemes, errResp := ctrl.fundSchemeService.GetFundSchemesWithFilters(fundHouse, category, minReturns, minInvestment)
    if errResp != nil {
        c.JSON(errResp.Status, errResp)
        return
    }

    // Return response
    c.JSON(http.StatusOK, schemes)
}
