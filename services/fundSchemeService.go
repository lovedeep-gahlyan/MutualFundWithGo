package services

import (
    "mutualfund/models"
    "net/http"
    "mutualfund/repositories"
)

type FundSchemeService struct {
    fundSchemeRepository *repositories.FundSchemeRepository
    orderRepository *repositories.OrderRepository
}

func NewFundSchemeService(fundSchemeRepository *repositories.FundSchemeRepository, orderRepository *repositories.OrderRepository) *FundSchemeService {
	return &FundSchemeService{
		fundSchemeRepository: fundSchemeRepository,
        orderRepository: orderRepository,
	}
}

// CreateFundScheme creates a new fund scheme
func (s FundSchemeService) CreateFundScheme(scheme *models.FundScheme) (*models.FundScheme, *models.ResponseError) {
    return s.fundSchemeRepository.CreateFundScheme(scheme)
}


// GetFundSchemes fetches all fund schemes
func (s FundSchemeService) GetFundSchemes() ([]models.FundScheme, *models.ResponseError) {
    return s.fundSchemeRepository.GetFundSchemes()
}

// GetFundSchemeByID fetches a fund scheme by its ID
func (s FundSchemeService) GetFundSchemeByID(id int) (*models.FundScheme, *models.ResponseError) {
    return s.fundSchemeRepository.GetFundSchemeByID(id)
}

// UpdateFundScheme updates an existing fund scheme
func (s FundSchemeService) UpdateFundScheme(scheme *models.FundScheme) *models.ResponseError {
    return s.fundSchemeRepository.UpdateFundScheme(scheme)
}

// DeleteFundScheme deletes a fund scheme by its ID
func (s FundSchemeService) DeleteFundScheme(id int) *models.ResponseError {
    // Check for existing orders
    orders, err := s.orderRepository.GetOrdersByFundID(id)
    if err != nil {
        return  &models.ResponseError{
            Message: err.Message,
            Status:  http.StatusConflict,
        }
    }
    if len(orders) > 0 {
        return &models.ResponseError{
            Message: "Cannot delete fund scheme with existing orders.",
            Status:  http.StatusConflict,
        }
    }
    //proceed for deletion
    return s.fundSchemeRepository.DeleteFundScheme(id)
}

func (s FundSchemeService) GetFundSchemesWithFilters(fundHouse, category string, minReturns, minInvestment float64) ([]models.FundScheme, *models.ResponseError) {
    return s.fundSchemeRepository.GetFundSchemesWithFilters(fundHouse, category, minReturns, minInvestment)
}
