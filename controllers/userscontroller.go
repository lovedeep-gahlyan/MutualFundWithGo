package controllers

import(
	"encoding/json"
	"io"
	"log"
	"net/http"
	"github.com/gin-gonic/gin"
	"mutualfund/models"
	"mutualfund/services"
)

type UsersController struct{
	usersService *services.UsersService
}

func NewUsersController(usersService *services.UsersService) *UsersController{
	return &UsersController{
		usersService: usersService,
	}
}

func (uc UsersController) CreateUser(ctx *gin.Context){
	body, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		log.Println("Error while reading create USER request body", err)
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	var user models.User
	err = json.Unmarshal(body, &user)
	if err != nil {
		log.Println("Error while unmarshaling create user request body", err)
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	response, responseErr := uc.usersService.CreateUser(&user)
	if responseErr != nil {
		ctx.AbortWithStatusJSON(responseErr.Status, responseErr)
		return
	}

	ctx.JSON(http.StatusOK, response)
}

// LoginUser handles user login
func (uc UsersController) LoginUser(ctx *gin.Context) {
    var input struct {
        UserName string `json:"user_name"`
        Password string `json:"password"`
    }
    if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
        return
    }
    _, err := uc.usersService.LoginUser(input.UserName, input.Password)
    if err != nil {
		ctx.AbortWithStatusJSON(err.Status, err)
        return
    }
    ctx.JSON(http.StatusOK, gin.H{"message": "login successful"})
}
