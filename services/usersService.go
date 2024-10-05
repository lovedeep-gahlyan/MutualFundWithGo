package services 

import (
	"mutualfund/models"
	"mutualfund/repositories"
	"net/http"
)

type UsersService struct{
	usersRepository *repositories.UsersRepository
}

func NewUsersService(usersRepository *repositories.UsersRepository) *UsersService {
	return &UsersService{
		usersRepository: usersRepository,
	}
}

func (us UsersService) CreateUser(user *models.User) (*models.User, *models.ResponseError) {
	// responseErr := validateUser(user)
	// if responseErr != nil {
	// 	return nil, responseErr
	// }

	return us.usersRepository.CreateUser(user)
}

func (us UsersService) LoginUser(username, password string) (*models.User, *models.ResponseError) {
    user, err := us.usersRepository.FindUserByUsername(username)
    if err != nil || user.Password != password {
        return nil, &models.ResponseError{
			Message: "User Not Found",
			Status: http.StatusInternalServerError,
		}
    }
    return user, nil
}