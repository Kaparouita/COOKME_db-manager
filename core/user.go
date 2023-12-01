package core

import (
	"db-manager/ports"
	"fmt"

	"github.com/Kaparouita/models/models"
)

type UserService struct {
	db ports.Db
}

func NewUserService(db ports.Db) *UserService {
	return &UserService{
		db,
	}
}

func (srv *UserService) Register(user *models.User) *models.User {
	err := srv.db.SaveUser(user)
	if err != nil {
		user.StatusCode = 400
		user.Message = fmt.Sprintf("Couldnt register USER : %v", err)
		return user
	}
	user.StatusCode = 200
	return user
}

func (srv *UserService) UpdateUser(user *models.User) *models.User {
	err := srv.db.UpdateUser(user)
	if err != nil {
		user.StatusCode = 400
		user.Message = fmt.Sprintf("Couldnt update USER : %v", err)
		return user
	}
	user.StatusCode = 201
	return user
}

func (srv *UserService) GetUser(user *models.User) *models.User {
	err := srv.db.GetUser(user)
	if err != nil {
		user.StatusCode = 400
		user.Message = fmt.Sprintf("Couldnt get USER : %v", err)
		return user
	}
	user.StatusCode = 200
	return user
}

func (srv *UserService) DeleteUser(user *models.User) *models.Response {
	resp := &models.Response{}
	err := srv.db.DeleteUser(user)
	if err != nil {
		resp.StatusCode = 400
		resp.Message = fmt.Sprintf("Couldnt delete USER : %v", err)
		return resp
	}
	resp.StatusCode = 200
	resp.Message = "Deleted user successfully"
	return resp
}

func (srv *UserService) GetUsers() ([]models.User, error) {
	users, err := srv.db.GetUsers()
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (srv *UserService) Login(login *models.LoginResp) *models.User {
	user, err := srv.db.Login(login)
	if err != nil || user.Username == "" {
		user.StatusCode = 400
		if err != nil {
			user.Message = fmt.Sprintf("Couldnt get USER : %v", err)
		} else {
			user.Message = fmt.Sprintf("Wrong username/password")
		}
		return user
	}
	user.StatusCode = 200
	return user
}
