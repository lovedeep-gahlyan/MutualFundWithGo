package repositories

import (
    "database/sql"
    "mutualfund/models"
    "net/http"
	"time"
	"log"
)

type OrderRepository struct {
    dbHandler *sql.DB
}

func NewOrderRepository(dbHandler *sql.DB) *OrderRepository {
    return &OrderRepository{dbHandler: dbHandler}
}

// CreateOrder inserts a new order into the database
func (r OrderRepository) CreateOrder(order *models.Order) (*models.Order, *models.ResponseError) {
    query := `INSERT INTO orders (user_id, fund_id, quantity, order_value, order_type, order_date)
              VALUES (?, ?, ?, ?, ?, ?)`
    
    order.OrderDate = time.Now() // Set the order date to current time

    res, err := r.dbHandler.Exec(query, order.UserID, order.SchemeID, order.Units, order.OrderValue, order.Action, order.OrderDate)
    if err != nil {
        return nil, &models.ResponseError{
            Message: err.Error(),
            Status:  http.StatusInternalServerError,
        }
    }
    
    orderID, err := res.LastInsertId()
    if err != nil {
        return nil, &models.ResponseError{
            Message: err.Error(),
            Status:  http.StatusInternalServerError,
        }
    }
    
    order.OrderID = orderID
    return order, nil
}

// GetOrders retrieves all orders by a specific user ID
func (r OrderRepository) GetOrdersByUserID(userID int64) ([]models.Order, *models.ResponseError) {
    query := `SELECT order_id, user_id, fund_id, quantity, order_value, order_type FROM orders WHERE user_id = ?`
    rows, err := r.dbHandler.Query(query, userID)
	log.Println("data: ", rows)
    if err != nil {
        return nil, &models.ResponseError{
            Message: err.Error(),
            Status:  http.StatusInternalServerError,
        }
    }
    defer rows.Close()

    var orders []models.Order
    for rows.Next() {
        var order models.Order
        if err := rows.Scan(&order.OrderID, &order.UserID, &order.SchemeID, &order.Units, &order.OrderValue, &order.Action); err != nil {
            log.Println("error is ", err.Error())
			return nil, &models.ResponseError{
                Message: err.Error(),
                Status:  http.StatusInternalServerError,
            }
        }
		log.Println("order is ", order)
        orders = append(orders, order)
    }
    return orders, nil
}

// GetOrdersByFundID retrieves all orders for a specific fund ID
func (r OrderRepository) GetOrdersByFundID(fundID int) ([]models.Order, *models.ResponseError) {
    query := `SELECT order_id, user_id, fund_id, quantity, order_value, order_type FROM orders WHERE fund_id = ?`
    rows, err := r.dbHandler.Query(query, fundID)
    if err != nil {
        return nil, &models.ResponseError{
            Message: err.Error(),
            Status:  http.StatusInternalServerError,
        }
    }
    defer rows.Close()

    var orders []models.Order
    for rows.Next() {
        var order models.Order
        if err := rows.Scan(&order.OrderID, &order.UserID, &order.SchemeID, &order.Units, &order.OrderValue, &order.Action); err != nil {
            return nil, &models.ResponseError{
                Message: err.Error(),
                Status:  http.StatusInternalServerError,
            }
        }
        orders = append(orders, order)
    }
    return orders, nil
}
