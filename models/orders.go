package models

import "time"

type Order struct {
    OrderID     int64     `json:"order_id"`
    UserID      int64     `json:"user_id"`
    SchemeID    int64     `json:"scheme_id"`
    Units       float64   `json:"units"`
    OrderValue  float64   `json:"order_value"`
    Action      string    `json:"action"` // buy or sell
    OrderDate   time.Time `json:"order_date"` // Date and time of the order
}
