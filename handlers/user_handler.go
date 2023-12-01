package handlers

import (
	"db-manager/core"
	"encoding/json"
	"fmt"

	"github.com/Kaparouita/models/models"
	"github.com/Kaparouita/models/myrabbit"
	"github.com/Kaparouita/models/myrabbit/amqp"
)

//create a rabbitmq connection

type UserHandler struct {
	handler *amqp.AmqpHandler
	srv     *core.UserService
}

func NewUserHandler(srv *core.UserService, handler *amqp.AmqpHandler) *UserHandler {
	return &UserHandler{
		handler: handler,
		srv:     srv,
	}
}

func (handler *UserHandler) GetUser(msgs <-chan myrabbit.Delivery, pubCh myrabbit.Channel, subCh myrabbit.Channel) {
	for msg := range msgs {
		user := &models.User{}
		resp := &models.User{}
		err := json.Unmarshal(msg.Body(), user)
		if err != nil {
			resp.StatusCode = 400
			resp.Message = fmt.Sprintf("Couldnt unmarshal IUpdateIngredient : %v", err)
			pubCh.Respond(msg, resp)
			return
		}
		resp = handler.srv.GetUser(user)
		pubCh.Respond(msg, resp)
	}
}

func (handler *UserHandler) Register(msgs <-chan myrabbit.Delivery, pubCh myrabbit.Channel, subCh myrabbit.Channel) {
	for msg := range msgs {
		user := &models.User{}
		resp := &models.User{}
		err := json.Unmarshal(msg.Body(), user)
		if err != nil {
			resp.StatusCode = 400
			resp.Message = fmt.Sprintf("Couldnt unmarshal ISaveIngredient : %v", err)
			pubCh.Respond(msg, resp)
			return
		}
		resp = handler.srv.Register(user)
		pubCh.Respond(msg, resp)
	}
}

func (handler *UserHandler) UpdateUser(msgs <-chan myrabbit.Delivery, pubCh myrabbit.Channel, subCh myrabbit.Channel) {
	for msg := range msgs {
		user := &models.User{}
		resp := &models.User{}
		err := json.Unmarshal(msg.Body(), user)
		if err != nil {
			resp.StatusCode = 400
			resp.Message = fmt.Sprintf("Couldnt unmarshal ISaveIngredient : %v", err)
			pubCh.Respond(msg, resp)
			return
		}
		resp = handler.srv.UpdateUser(user)
		pubCh.Respond(msg, resp)
	}
}

func (handler *UserHandler) DeleteUser(msgs <-chan myrabbit.Delivery, pubCh myrabbit.Channel, subCh myrabbit.Channel) {
	for msg := range msgs {
		user := &models.User{}
		resp := &models.Response{}
		err := json.Unmarshal(msg.Body(), user)
		if err != nil {
			resp.StatusCode = 400
			resp.Message = fmt.Sprintf("Couldnt unmarshal ISaveIngredient : %v", err)
			pubCh.Respond(msg, resp)
			return
		}
		resp = handler.srv.DeleteUser(user)
		pubCh.Respond(msg, resp)
	}
}

func (handler *UserHandler) Login(msgs <-chan myrabbit.Delivery, pubCh myrabbit.Channel, subCh myrabbit.Channel) {
	for msg := range msgs {
		user := &models.User{}
		cred := &models.LoginResp{}
		resp := &models.Response{}

		err := json.Unmarshal(msg.Body(), cred)
		if err != nil {
			resp.StatusCode = 400
			resp.Message = fmt.Sprintf("Couldnt unmarshal ISaveIngredient : %v", err)
			pubCh.Respond(msg, resp)
			return
		}
		user = handler.srv.Login(cred)
		pubCh.Respond(msg, user)
	}
}

func (handler *UserHandler) GetUsers(msgs <-chan myrabbit.Delivery, pubCh myrabbit.Channel, subCh myrabbit.Channel) {
	for msg := range msgs {
		resp := &models.Response{}
		users, err := handler.srv.GetUsers()
		if err != nil {
			resp.StatusCode = 400
			resp.Message = fmt.Sprintf("Couldnt unmarshal ISaveIngredient : %v", err)
			pubCh.Respond(msg, resp)
			return
		}
		pubCh.Respond(msg, users)
	}
}

/*
func handleMessage(message Message) {
	switch message.Operation {
	case "GetRecipe":
		recipe, err := db.GetRecipe(message.Data.RecipeID)
		if err != nil {
			// Handle error
			return
		}
		// Send the recipe as a response, e.g., publish to another queue or send an HTTP response
	case "SaveRecipe":
		// Handle saving the recipe
	// Handle other operations similarly
	default:
		// Unknown operation, handle accordingly
	}
}
*/
