package controllers

import (
    "github.com/gin-gonic/gin"
    "mutualfund/models"
	"mutualfund/services"
    "net/http"
    "strconv"
    "log"
)

type OrderController struct {
    orderService *services.OrderService
}

func NewOrderController(orderService *services.OrderService) *OrderController {
    return &OrderController{orderService: orderService}
}

// CreateOrder handles the POST request to create an order
func (ctrl OrderController) CreateOrder(c *gin.Context) {
    var order models.Order
    if err := c.ShouldBindJSON(&order); err != nil {
        c.JSON(http.StatusBadRequest, &models.ResponseError{
            Message: "Invalid input",
            Status:  http.StatusBadRequest,
        })
        return
    }

    createdOrder, err := ctrl.orderService.CreateOrder(&order)
    if err != nil {
        c.JSON(err.Status, err)
        return
    }

    c.JSON(http.StatusOK, createdOrder)
}

// GetOrdersByUserID handles GET requests to retrieve orders for a user
func (ctrl OrderController) GetOrdersByUserID(c *gin.Context) {
    userID := c.Param("id")
    log.Println("userID: ", userID)
    id, err := strconv.ParseInt(userID, 10, 64)
    if err != nil {
        c.AbortWithStatusJSON(http.StatusBadRequest, models.ResponseError{
            Message: "Invalid user ID",
            Status:  http.StatusBadRequest,
        })
        return
    }
    response, responseErr := ctrl.orderService.GetOrdersByUserID(id)
    if err != nil {
		c.AbortWithStatusJSON(responseErr.Status, responseErr)
        return
    }
    c.JSON(http.StatusOK, response)
}
