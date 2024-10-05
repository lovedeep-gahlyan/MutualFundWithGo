package repositories

import (
	"database/sql"
	"net/http"
	"mutualfund/models"
	"strconv"
)

type UsersRepository struct {
	dbHandler *sql.DB
}

func NewUsersRepository(dbHandler *sql.DB) *UsersRepository {
	return &UsersRepository{
		dbHandler: dbHandler,
	}
}


func (ur UsersRepository) CreateUser(user *models.User) (*models.User, *models.ResponseError) {
	query := `
		INSERT INTO users(user_name, password, role)
		VALUES (?, ?, ?)`

	res, err := ur.dbHandler.Exec(query, user.UserName, user.Password, user.Role)
	if err != nil {
		return nil, &models.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	userId, err := res.LastInsertId()
	if err != nil {
		return nil, &models.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
	}

	return &models.User{
		ID:        strconv.FormatInt(userId, 10),
		UserName: user.UserName,
		Password:  user.Password,
		Role:       user.Role,
	}, nil
}

// FindUserByUsername fetches a user by their username
func (ur UsersRepository) FindUserByUsername(username string) (*models.User, *models.ResponseError) {
    query := "SELECT user_id, user_name, password, role FROM users WHERE user_name = ?"
    row := ur.dbHandler.QueryRow(query, username)

    var user models.User
    err := row.Scan(&user.ID, &user.UserName, &user.Password, &user.Role)
    if err != nil {
		return nil, &models.ResponseError{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		}
    }
    return &user, nil
}
