package models

type FundScheme struct {
    SchemeID      string     `json:"fund_id"`
    FundHouse     string  `json:"fund_house"`
    FundManager   string  `json:"fund_manager"`
    Category      string  `json:"category"`       // (equity, debt, hybrid)
    ReturnsPA     float64 `json:"returns_pa"`     // Returns per annum
    CurrentValue  float64 `json:"current_value"`  // Current value of the scheme
    MinInvestment float64 `json:"min_investment"` // Minimum investment amount
}

