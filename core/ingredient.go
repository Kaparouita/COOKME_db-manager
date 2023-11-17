package core

import (
	"db-manager/ports"
	"fmt"

	"github.com/Kaparouita/models/models"
)

type IngredientService struct {
	db ports.Db
}

//-----------------------------------MarketIngredient---------------------------------------//

func (srv *IngredientService) SaveMarketIngredient(MarketIngredient *models.MarketIngredient) *models.Response {
	resp := &models.Response{}
	err := srv.db.SaveMarketIngredient(MarketIngredient)
	if err != nil {
		resp.StatusCode = 400
		resp.Message = fmt.Sprintf("Couldnt create MarketIngredient : %v", err)
		return resp
	}
	resp.StatusCode = 200
	return resp
}

func (srv *IngredientService) UpdateMarketIngredient(MarketIngredient *models.MarketIngredient) *models.Response {
	resp := &models.Response{}
	err := srv.db.UpdateMarketIngredient(MarketIngredient)
	if err != nil {
		resp.StatusCode = 400
		resp.Message = fmt.Sprintf("Couldnt update MarketIngredient : %v", err)
		return resp
	}
	resp.StatusCode = 201
	return resp
}

func (srv *IngredientService) GetMarketIngredient(MarketIngredient *models.MarketIngredient) (*models.MarketIngredient, *models.Response) {
	resp := &models.Response{}
	err := srv.db.GetMarketIngredient(MarketIngredient)
	if err != nil {
		resp.StatusCode = 400
		resp.Message = fmt.Sprintf("Couldnt get MarketIngredient : %v", err)
		return MarketIngredient, resp
	}
	resp.StatusCode = 200
	return MarketIngredient, resp
}

func (srv *IngredientService) DeleteMarketIngredient(MarketIngredient *models.MarketIngredient) *models.Response {
	resp := &models.Response{}
	err := srv.db.DeleteMarketIngredient(MarketIngredient)
	if err != nil {
		resp.StatusCode = 400
		resp.Message = fmt.Sprintf("Couldnt delete MarketIngredient : %v", err)
		return resp
	}
	resp.StatusCode = 200
	resp.Message = "Deleted MarketIngredient successfully"
	return resp
}

func (srv *IngredientService) GetMarketIngredients() ([]models.MarketIngredient, error) {
	MarketIngredients, err := srv.db.GetMarketIngredients()
	if err != nil {
		return nil, err
	}
	return MarketIngredients, nil
}
