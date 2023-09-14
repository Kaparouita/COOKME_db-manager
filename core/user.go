package core

import (
	"fmt"
	"github.com/Kaparouita/db-manager/domain"
)

func (srv *Service) Register(user *domain.User) *domain.User {
	err := srv.db.SaveUser(user)
	if err != nil {
		user.StatusCode = 400
		user.Message = fmt.Sprintf("Couldnt register USER : %v", err)
		return user
	}
	user.StatusCode = 200
	return user
}

func (srv *Service) UpdateUser(user *domain.User) *domain.User {
	err := srv.db.UpdateUser(user)
	if err != nil {
		user.StatusCode = 400
		user.Message = fmt.Sprintf("Couldnt update USER : %v", err)
		return user
	}
	user.StatusCode = 201
	return user
}

func (srv *Service) GetUser(user *domain.User) *domain.User {
	err := srv.db.GetUser(user)
	if err != nil {
		user.StatusCode = 400
		user.Message = fmt.Sprintf("Couldnt get USER : %v", err)
		return user
	}
	user.StatusCode = 200
	return user
}

func (srv *Service) DeleteUser(user *domain.User) *domain.Response {
	resp := &domain.Response{}
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

func (srv *Service) GetUsers() ([]domain.User, error) {
	users, err := srv.db.GetUsers()
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (srv *Service) Login(login *domain.LoginResp) *domain.User {
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

func (srv *Service) ReviewIngridient(review *domain.Review) *domain.Response {
	resp := &domain.Response{}
	Ingridient := &domain.Ingridient{Id: review.IngridientID}

	if err := srv.db.GetIngridient(Ingridient); err != nil {
		resp.StatusCode = 400
		resp.Message = fmt.Sprintf("Couldnt get Ingridient : %v", err)
		return resp
	}
	Ingridient.Reviews = append(Ingridient.Reviews, *review)
	if err := srv.db.UpdateIngridient(Ingridient); err != nil {
		resp.StatusCode = 400
		resp.Message = fmt.Sprintf("Couldnt update Ingridient : %v", err)
		return resp
	}
	resp.StatusCode = 200
	return resp
}

func (srv *Service) GetUserPerType(group string) ([]domain.UserPerType, *domain.Response) {
	resp := &domain.Response{}

	users, err := srv.db.GetUserPerUserType(group)
	if err != nil {
		resp.StatusCode = 400
		resp.Message = fmt.Sprintf("Couldn't get users: %v", err)
		return users, resp
	}

	resp.StatusCode = 200
	return users, resp
}
