package core

import (
	"db-manager/ports"
	"fmt"

	"github.com/Kaparouita/models/models"
)

type IngredientService struct {
	db ports.Db
}

func (srv *IngredientService) SaveIngredient(Ingredient *models.Ingredient) *models.Response {
	resp := &models.Response{}
	err := srv.db.SaveIngredient(Ingredient)
	if err != nil {
		resp.StatusCode = 400
		resp.Message = fmt.Sprintf("Couldnt create Ingredient : %v", err)
		return resp
	}
	resp.StatusCode = 200
	return resp
}

func (srv *IngredientService) UpdateIngredient(Ingredient *models.Ingredient) *models.Response {
	resp := &models.Response{}
	err := srv.db.UpdateIngredient(Ingredient)
	if err != nil {
		resp.StatusCode = 400
		resp.Message = fmt.Sprintf("Couldnt update Ingredient : %v", err)
		return resp
	}
	resp.StatusCode = 201
	return resp
}

func (srv *IngredientService) GetIngredient(Ingredient *models.Ingredient) (*models.Ingredient, *models.Response) {
	resp := &models.Response{}
	err := srv.db.GetIngredient(Ingredient)
	if err != nil {
		resp.StatusCode = 400
		resp.Message = fmt.Sprintf("Couldnt get Ingredient : %v", err)
		return Ingredient, resp
	}
	resp.StatusCode = 200
	return Ingredient, resp
}

func (srv *IngredientService) DeleteIngredient(Ingredient *models.Ingredient) *models.Response {
	resp := &models.Response{}
	err := srv.db.DeleteIngredient(Ingredient)
	if err != nil {
		resp.StatusCode = 400
		resp.Message = fmt.Sprintf("Couldnt delete Ingredient : %v", err)
		return resp
	}
	resp.StatusCode = 200
	resp.Message = "Deleted Ingredient successfully"
	return resp
}

func (srv *IngredientService) GetIngredients() ([]models.Ingredient, error) {
	Ingredients, err := srv.db.GetIngredients()
	if err != nil {
		return nil, err
	}
	return Ingredients, nil
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
