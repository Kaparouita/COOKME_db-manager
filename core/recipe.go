package core

import (
	"fmt"

	"github.com/Kaparouita/models/models"
)

func (srv *Service) SaveRecipe(Recipe *models.Recipe) *models.Response {
	resp := &models.Response{}
	err := srv.db.SaveRecipe(Recipe)
	if err != nil {
		resp.StatusCode = 400
		resp.Message = fmt.Sprintf("Couldnt create Recipe : %v", err)
		return resp
	}
	resp.StatusCode = 200
	return resp
}

func (srv *Service) UpdateRecipe(Recipe *models.Recipe) *models.Response {
	resp := &models.Response{}
	err := srv.db.UpdateRecipe(Recipe)
	if err != nil {
		resp.StatusCode = 400
		resp.Message = fmt.Sprintf("Couldnt update Recipe : %v", err)
		return resp
	}
	resp.StatusCode = 201
	return resp
}

func (srv *Service) GetRecipe(Recipe *models.Recipe) (*models.Recipe, *models.Response) {
	resp := &models.Response{}
	err := srv.db.GetRecipe(Recipe)
	if err != nil {
		resp.StatusCode = 400
		resp.Message = fmt.Sprintf("Couldnt get Recipe : %v", err)
		return Recipe, resp
	}
	resp.StatusCode = 200
	return Recipe, resp
}

func (srv *Service) DeleteRecipe(Recipe *models.Recipe) *models.Response {
	resp := &models.Response{}
	err := srv.db.DeleteRecipe(Recipe)
	if err != nil {
		resp.StatusCode = 400
		resp.Message = fmt.Sprintf("Couldnt delete Recipe : %v", err)
		return resp
	}
	resp.StatusCode = 200
	resp.Message = "Deleted Recipe successfully"
	return resp
}

func (srv *Service) GetRecipes() ([]models.Recipe, error) {
	Recipes, err := srv.db.GetRecipes()
	if err != nil {
		return nil, err
	}
	return Recipes, nil
}

//-----------------------------------FinalRecipe---------------------------------------//

func (srv *Service) SaveFinalRecipe(FinalRecipe *models.FinalRecipe) *models.Response {
	resp := &models.Response{}
	err := srv.db.SaveFinalRecipe(FinalRecipe)
	if err != nil {
		resp.StatusCode = 400
		resp.Message = fmt.Sprintf("Couldnt create FinalRecipe : %v", err)
		return resp
	}
	resp.StatusCode = 200
	return resp
}

func (srv *Service) UpdateFinalRecipe(FinalRecipe *models.FinalRecipe) *models.Response {
	resp := &models.Response{}
	err := srv.db.UpdateFinalRecipe(FinalRecipe)
	if err != nil {
		resp.StatusCode = 400
		resp.Message = fmt.Sprintf("Couldnt update FinalRecipe : %v", err)
		return resp
	}
	resp.StatusCode = 201
	return resp
}

func (srv *Service) GetFinalRecipe(FinalRecipe *models.FinalRecipe) (*models.FinalRecipe, *models.Response) {
	resp := &models.Response{}
	err := srv.db.GetFinalRecipe(FinalRecipe)
	if err != nil {
		resp.StatusCode = 400
		resp.Message = fmt.Sprintf("Couldnt get FinalRecipe : %v", err)
		return FinalRecipe, resp
	}
	resp.StatusCode = 200
	return FinalRecipe, resp
}

func (srv *Service) DeleteFinalRecipe(FinalRecipe *models.FinalRecipe) *models.Response {
	resp := &models.Response{}
	err := srv.db.DeleteFinalRecipe(FinalRecipe)
	if err != nil {
		resp.StatusCode = 400
		resp.Message = fmt.Sprintf("Couldnt delete FinalRecipe : %v", err)
		return resp
	}
	resp.StatusCode = 200
	resp.Message = "Deleted FinalRecipe successfully"
	return resp
}

func (srv *Service) GetFinalRecipes() ([]models.FinalRecipe, error) {
	FinalRecipes, err := srv.db.GetFinalRecipes()
	if err != nil {
		return nil, err
	}
	return FinalRecipes, nil
}
