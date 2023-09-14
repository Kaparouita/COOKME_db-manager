package core

import (
	"fmt"
	"github.com/Kaparouita/db-manager/domain"
)

func (srv *Service) GetIngridients(pagin *domain.IngridientPagin) (*domain.IngridientResponse, error) {
	return srv.db.GetIngridients(pagin)
}

func (srv *Service) SaveIngridient(Ingridient *domain.Ingridient) *domain.Response {
	resp := &domain.Response{}
	err := srv.db.SaveIngridient(Ingridient)
	if err != nil {
		resp.StatusCode = 400
		resp.Message = fmt.Sprintf("Couldnt create Ingridient : %v", err)
		return resp
	}
	resp.StatusCode = 200
	return resp
}

func (srv *Service) UpdateIngridient(Ingridient *domain.Ingridient) *domain.Response {
	resp := &domain.Response{}
	err := srv.db.UpdateIngridient(Ingridient)
	if err != nil {
		resp.StatusCode = 400
		resp.Message = fmt.Sprintf("Couldnt update Ingridient : %v", err)
		return resp
	}
	resp.StatusCode = 201
	return resp
}

func (srv *Service) GetIngridient(Ingridient *domain.Ingridient) (*domain.Ingridient, *domain.Response) {
	resp := &domain.Response{}
	err := srv.db.GetIngridient(Ingridient)
	if err != nil {
		resp.StatusCode = 400
		resp.Message = fmt.Sprintf("Couldnt get Ingridient : %v", err)
		return Ingridient, resp
	}
	resp.StatusCode = 200
	return Ingridient, resp
}

func (srv *Service) DeleteIngridient(Ingridient *domain.Ingridient) *domain.Response {
	resp := &domain.Response{}
	err := srv.db.DeleteIngridient(Ingridient)
	if err != nil {
		resp.StatusCode = 400
		resp.Message = fmt.Sprintf("Couldnt delete Ingridient : %v", err)
		return resp
	}
	resp.StatusCode = 200
	resp.Message = "Deleted Ingridient successfully"
	return resp
}

func (srv *Service) GetIngridientsPerUni(group string) ([]domain.IngridientsPerUni, *domain.Response) {
	resp := &domain.Response{}

	Ingridients, err := srv.db.GetIngridientsPerUni(group)
	if err != nil {
		resp.StatusCode = 400
		resp.Message = fmt.Sprintf("Couldn't get Ingridient: %v", err)
		return Ingridients, resp
	}

	resp.StatusCode = 200
	return Ingridients, resp
}
