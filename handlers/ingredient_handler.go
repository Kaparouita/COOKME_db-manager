package handlers

import (
	"db-manager/core"
	"encoding/json"
	"fmt"

	"github.com/Kaparouita/models/models"
	"github.com/Kaparouita/models/myrabbit"
	"github.com/Kaparouita/models/myrabbit/amqp"
)

type IngredientHandler struct {
	handler *amqp.AmqpHandler
	srv     *core.IngredientService
}

func NewIngredientHandler(srv *core.IngredientService, handler *amqp.AmqpHandler) *IngredientHandler {
	return &IngredientHandler{
		handler: handler,
		srv:     srv,
	}
}

func (handler *IngredientHandler) UpdateMarketIngredient(msgs <-chan myrabbit.Delivery, pubCh myrabbit.Channel, subCh myrabbit.Channel) {
	for msg := range msgs {
		marketIngredient := &models.MarketIngredient{}
		resp := &models.Response{}
		err := json.Unmarshal(msg.Body(), marketIngredient)
		if err != nil {
			resp.StatusCode = 400
			resp.Message = fmt.Sprintf("Couldnt unmarshal IUpdateIngredient : %v", err)
			pubCh.Respond(msg, resp)
			return
		}
		resp = handler.srv.UpdateMarketIngredient(marketIngredient)
		pubCh.Respond(msg, resp)
	}
}

func (handler *IngredientHandler) SaveMarketIngredient(msgs <-chan myrabbit.Delivery, pubCh myrabbit.Channel, subCh myrabbit.Channel) {
	for msg := range msgs {
		marketIngredient := &models.MarketIngredient{}
		resp := &models.Response{}
		err := json.Unmarshal(msg.Body(), marketIngredient)
		if err != nil {
			resp.StatusCode = 400
			resp.Message = fmt.Sprintf("Couldnt unmarshal ISaveIngredient : %v", err)
			pubCh.Respond(msg, resp)
			return
		}
		resp = handler.srv.SaveMarketIngredient(marketIngredient)
		pubCh.Respond(msg, resp)
	}
}

func (handler *IngredientHandler) GetMarketIngredient(msgs <-chan myrabbit.Delivery, pubCh myrabbit.Channel, subCh myrabbit.Channel) {
	for msg := range msgs {
		marketIngredient := &models.MarketIngredient{}
		resp := &models.Response{}
		err := json.Unmarshal(msg.Body(), marketIngredient)
		if err != nil {
			resp.StatusCode = 400
			resp.Message = fmt.Sprintf("Couldnt unmarshal IGetIngredient : %v", err)
			pubCh.Respond(msg, resp)
			return
		}
		_, resp = handler.srv.GetMarketIngredient(marketIngredient)
		pubCh.Respond(msg, resp)
	}
}

func (handler *IngredientHandler) DeleteMarketIngredient(msgs <-chan myrabbit.Delivery, pubCh myrabbit.Channel, subCh myrabbit.Channel) {
	for msg := range msgs {
		marketIngredient := &models.MarketIngredient{}
		resp := &models.Response{}
		err := json.Unmarshal(msg.Body(), marketIngredient)
		if err != nil {
			resp.StatusCode = 400
			resp.Message = fmt.Sprintf("Couldnt unmarshal IDeleteIngredient : %v", err)
			pubCh.Respond(msg, resp)
			return
		}
		resp = handler.srv.DeleteMarketIngredient(marketIngredient)
		pubCh.Respond(msg, resp)
	}
}

// FIXXXXX
func (handler *IngredientHandler) GetMarketIngredients(msgs <-chan myrabbit.Delivery, pubCh myrabbit.Channel, subCh myrabbit.Channel) {
	for msg := range msgs {
		marketIngredient := &models.MarketIngredient{}
		resp := &models.Response{}
		err := json.Unmarshal(msg.Body(), marketIngredient)
		if err != nil {
			resp.StatusCode = 400
			resp.Message = fmt.Sprintf("Couldnt unmarshal IGetIngredient : %v", err)
			pubCh.Respond(msg, resp)
			return
		}
		_, resp = handler.srv.GetMarketIngredient(marketIngredient)
		pubCh.Respond(msg, resp)
	}
}
