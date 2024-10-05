package repositories

import (
    "database/sql"
    "mutualfund/models"
	"net/http"
	"strconv"
	"log"
)

type FundSchemeRepository struct {
    dbHandler *sql.DB
}

func NewFundSchemeRepository(dbHandler *sql.DB) *FundSchemeRepository {
	return &FundSchemeRepository{
		dbHandler: dbHandler,
	}
}

// CreateFundScheme inserts a new fund scheme into the database
func (r FundSchemeRepository) CreateFundScheme(scheme *models.FundScheme) (*models.FundScheme, *models.ResponseError)  {
    query := `INSERT INTO fund_schemes (fundhouse_name, fund_manager_name, fund_category, returns, current_value, min_investment)
              VALUES (?, ?, ?, ?, ?, ?)`
	
	res,err := r.dbHandler.Exec(query,scheme.FundHouse, scheme.FundManager, scheme.Category, scheme.ReturnsPA, scheme.CurrentValue, scheme.MinInvestment)
	if err != nil {
		return nil, &models.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}
	schemeId, err := res.LastInsertId()
	if err != nil {
		return nil, &models.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	return &models.FundScheme{
		SchemeID : strconv.FormatInt(schemeId,10),
		FundHouse : scheme.FundHouse,
		FundManager : scheme.FundManager,
		Category : scheme.Category,
		ReturnsPA : scheme.ReturnsPA,
		CurrentValue : scheme.CurrentValue,
		MinInvestment : scheme.MinInvestment,
	},nil
}

// GetFundSchemes retrieves all fund schemes
func (r FundSchemeRepository) GetFundSchemes() ([]models.FundScheme, *models.ResponseError) {
    query := `SELECT * FROM fund_schemes`
    rows, err := r.dbHandler.Query(query)
	log.Println("rows : ", rows)
    if err != nil {
        return nil, &models.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
    }
    defer rows.Close()

    var schemes []models.FundScheme
    for rows.Next() {
        var scheme models.FundScheme
        if err := rows.Scan(&scheme.SchemeID, &scheme.FundHouse, &scheme.FundManager, &scheme.Category, &scheme.ReturnsPA, &scheme.CurrentValue, &scheme.MinInvestment); err != nil {
            return nil, &models.ResponseError{
				Message: err.Error(),
				Status:  http.StatusInternalServerError,
			}
        }
		log.Println("fundschemes : ", scheme)
        schemes = append(schemes, scheme)
    }
    return schemes, nil
}

// GetFundSchemeByID retrieves a single fund scheme by its ID
func (r FundSchemeRepository) GetFundSchemeByID(id int) (*models.FundScheme, *models.ResponseError) {
    query := `SELECT * FROM fund_schemes WHERE fund_id = ?`
    row := r.dbHandler.QueryRow(query, id)

    var scheme models.FundScheme
    err := row.Scan(&scheme.SchemeID, &scheme.FundHouse, &scheme.FundManager, &scheme.Category, &scheme.ReturnsPA, &scheme.CurrentValue, &scheme.MinInvestment)
    if err != nil {
        return nil, &models.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
    }
    return &scheme, nil
}

// UpdateFundScheme updates an existing fund scheme by its ID
func (r FundSchemeRepository) UpdateFundScheme(scheme *models.FundScheme) *models.ResponseError {
    query := `UPDATE fund_schemes SET fundhouse_name=?, fund_manager_name=?, fund_category=?, returns=?, current_value=?, min_investment=?
              WHERE fund_id=?`
    stmt, err := r.dbHandler.Prepare(query)
    if err != nil {
        return &models.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
    }
    defer stmt.Close()

    _, err = stmt.Exec(scheme.FundHouse, scheme.FundManager, scheme.Category, scheme.ReturnsPA, scheme.CurrentValue, scheme.MinInvestment, scheme.SchemeID)
    if err != nil {
        return &models.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
    }

    return nil  // Added this to indicate a successful operation
}


// DeleteFundScheme deletes a fund scheme by its ID
func (r FundSchemeRepository) DeleteFundScheme(id int) *models.ResponseError {

    query := `DELETE FROM fund_schemes WHERE fund_id = ?`
    stmt, err := r.dbHandler.Prepare(query)
    if err != nil {
        return &models.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
    }
	
    defer stmt.Close()

    _, err = stmt.Exec(id)
	if err != nil {
		log.Println("error while deleting: ", err.Error())
        return &models.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
    }

    return nil  // Added this to indicate a successful deletion
}

// GetFundSchemesWithFilters retrieves fund schemes based on the provided filters
func (r FundSchemeRepository) GetFundSchemesWithFilters(fundHouse, category string, minReturns, minInvestment float64) ([]models.FundScheme, *models.ResponseError) {
    // Base query
    query := `SELECT * FROM fund_schemes WHERE 1=1`
    var params []interface{}

    // Add filters dynamically
    if fundHouse != "" {
        query += ` AND fundhouse_name = ?`
        params = append(params, fundHouse)
    }
    if category != "" {
        query += ` AND fund_category = ?`
        params = append(params, category)
    }
    if minReturns > 0 {
        query += ` AND returns >= ?`
        params = append(params, minReturns)
    }
    if minInvestment > 0 {
        query += ` AND min_investment >= ?`
        params = append(params, minInvestment)
    }

    // Execute query with params
    rows, err := r.dbHandler.Query(query, params...)
    if err != nil {
        return nil, &models.ResponseError{
            Message: err.Error(),
            Status:  http.StatusInternalServerError,
        }
    }
    defer rows.Close()

    // Parse result
    var schemes []models.FundScheme
    for rows.Next() {
        var scheme models.FundScheme
        if err := rows.Scan(&scheme.SchemeID, &scheme.FundHouse, &scheme.FundManager, &scheme.Category, &scheme.ReturnsPA, &scheme.CurrentValue, &scheme.MinInvestment); err != nil {
            return nil, &models.ResponseError{
                Message: err.Error(),
                Status:  http.StatusInternalServerError,
            }
        }
        schemes = append(schemes, scheme)
    }
    
    return schemes, nil
}
