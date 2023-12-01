package handlers

import (
	"db-manager/core"
	"encoding/json"
	"fmt"

	"github.com/Kaparouita/models/models"
	"github.com/Kaparouita/models/myrabbit"
	"github.com/Kaparouita/models/myrabbit/amqp"
)

type RecipeHandler struct {
	handler *amqp.AmqpHandler
	srv     *core.RecipeService
}

func NewRecipeHandler(srv *core.RecipeService, handler *amqp.AmqpHandler) *RecipeHandler {
	return &RecipeHandler{
		handler: handler,
		srv:     srv,
	}
}

func (handler *RecipeHandler) GetRecipe(msgs <-chan myrabbit.Delivery, pubCh myrabbit.Channel, subCh myrabbit.Channel) {
	for msg := range msgs {
		func() {
			defer msg.Ack(false)
			recipe := &models.Recipe{}
			resp := &models.Response{}
			err := json.Unmarshal(msg.Body(), recipe)
			if err != nil {
				resp.StatusCode = 400
				resp.Message = fmt.Sprintf("Couldnt unmarshal Recipe : %v", err)
				pubCh.Respond(msg, resp)
				return
			}
			_, resp = handler.srv.GetRecipe(recipe)
			pubCh.Respond(msg, resp)
		}()
	}
}

func (handler *RecipeHandler) SaveRecipe(msgs <-chan myrabbit.Delivery, pubCh myrabbit.Channel, subCh myrabbit.Channel) {
	for msg := range msgs {
		func() {
			defer msg.Ack(false)
			recipe := &models.Recipe{}
			resp := &models.Response{}
			err := json.Unmarshal(msg.Body(), recipe)
			if err != nil {
				resp.StatusCode = 400
				resp.Message = fmt.Sprintf("Couldnt unmarshal Recipe : %v", err)
				pubCh.Respond(msg, resp)
				return
			}
			resp = handler.srv.SaveRecipe(recipe)
			pubCh.Respond(msg, resp)
		}()
	}
}

func (handler *RecipeHandler) UpdateRecipe(msgs <-chan myrabbit.Delivery, pubCh myrabbit.Channel, subCh myrabbit.Channel) {
	for msg := range msgs {
		func() {
			defer msg.Ack(false)
			recipe := &models.Recipe{}
			resp := &models.Response{}
			err := json.Unmarshal(msg.Body(), recipe)
			if err != nil {
				resp.StatusCode = 400
				resp.Message = fmt.Sprintf("Couldnt unmarshal Recipe : %v", err)
				pubCh.Respond(msg, resp)
				return
			}
			resp = handler.srv.UpdateRecipe(recipe)
			pubCh.Respond(msg, resp)
		}()
	}
}

func (handler *RecipeHandler) DeleteRecipe(msgs <-chan myrabbit.Delivery, pubCh myrabbit.Channel, subCh myrabbit.Channel) {
	for msg := range msgs {
		func() {
			defer msg.Ack(false)
			recipe := &models.Recipe{}
			resp := &models.Response{}
			err := json.Unmarshal(msg.Body(), recipe)
			if err != nil {
				resp.StatusCode = 400
				resp.Message = fmt.Sprintf("Couldnt unmarshal Recipe : %v", err)
				pubCh.Respond(msg, resp)
				return
			}
			resp = handler.srv.DeleteRecipe(recipe)
			pubCh.Respond(msg, resp)
		}()
	}
}

func (handler *RecipeHandler) GetRecipes(msgs <-chan myrabbit.Delivery, pubCh myrabbit.Channel, subCh myrabbit.Channel) {
	for msg := range msgs {
		func() {
			defer msg.Ack(false)
			resp := &models.Response{}
			recipes, err := handler.srv.GetRecipes()
			if err != nil {
				resp.StatusCode = 400
				resp.Message = fmt.Sprintf("Couldnt get Recipes : %v", err)
				pubCh.Respond(msg, resp)
				return
			}
			pubCh.Respond(msg, recipes)
		}()
	}
}

func (handler *RecipeHandler) GetFinalRecipe(msgs <-chan myrabbit.Delivery, pubCh myrabbit.Channel, subCh myrabbit.Channel) {
	for msg := range msgs {
		func() {
			finalRecipe := &models.FinalRecipe{}
			resp := &models.Response{}
			err := json.Unmarshal(msg.Body(), finalRecipe)
			if err != nil {
				resp.StatusCode = 400
				resp.Message = fmt.Sprintf("Couldnt unmarshal FinalRecipe : %v", err)
				pubCh.Respond(msg, resp)
				return
			}
			_, resp = handler.srv.GetFinalRecipe(finalRecipe)
			pubCh.Respond(msg, resp)
		}()
	}
}

func (handler *RecipeHandler) SaveFinalRecipe(msgs <-chan myrabbit.Delivery, pubCh myrabbit.Channel, subCh myrabbit.Channel) {
	for msg := range msgs {
		finalRecipe := &models.FinalRecipe{}
		resp := &models.Response{}
		err := json.Unmarshal(msg.Body(), finalRecipe)
		if err != nil {
			resp.StatusCode = 400
			resp.Message = fmt.Sprintf("Couldnt unmarshal FinalRecipe : %v", err)
			pubCh.Respond(msg, resp)
			return
		}
		resp = handler.srv.SaveFinalRecipe(finalRecipe)
		pubCh.Respond(msg, resp)
	}
}

func (handler *RecipeHandler) UpdateFinalRecipe(msgs <-chan myrabbit.Delivery, pubCh myrabbit.Channel, subCh myrabbit.Channel) {
	for msg := range msgs {
		finalRecipe := &models.FinalRecipe{}
		resp := &models.Response{}
		err := json.Unmarshal(msg.Body(), finalRecipe)
		if err != nil {
			resp.StatusCode = 400
			resp.Message = fmt.Sprintf("Couldnt unmarshal FinalRecipe : %v", err)
			pubCh.Respond(msg, resp)
			return
		}
		resp = handler.srv.UpdateFinalRecipe(finalRecipe)
		pubCh.Respond(msg, resp)
	}
}

func (handler *RecipeHandler) DeleteFinalRecipe(msgs <-chan myrabbit.Delivery, pubCh myrabbit.Channel, subCh myrabbit.Channel) {
	for msg := range msgs {
		finalRecipe := &models.FinalRecipe{}
		resp := &models.Response{}
		err := json.Unmarshal(msg.Body(), finalRecipe)
		if err != nil {
			resp.StatusCode = 400
			resp.Message = fmt.Sprintf("Couldnt unmarshal FinalRecipe : %v", err)
			pubCh.Respond(msg, resp)
			return
		}
		resp = handler.srv.DeleteFinalRecipe(finalRecipe)
		pubCh.Respond(msg, resp)
	}
}

func (handler *RecipeHandler) GetFinalRecipes(msgs <-chan myrabbit.Delivery, pubCh myrabbit.Channel, subCh myrabbit.Channel) {
	for msg := range msgs {
		resp := &models.Response{}
		finalRecipes, err := handler.srv.GetFinalRecipes()
		if err != nil {
			resp.StatusCode = 400
			resp.Message = fmt.Sprintf("Couldnt get FinalRecipes : %v", err)
			pubCh.Respond(msg, resp)
			return
		}
		pubCh.Respond(msg, finalRecipes)
	}
}
