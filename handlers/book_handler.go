package handlers

import (
	"encoding/json"
	"github.com/Kaparouita/db-manager/domain"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func (handler *Handler) GetIngridients(c *fiber.Ctx) error {
	page, err := strconv.Atoi(c.Query("page", "1"))
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	perPage, err := strconv.Atoi(c.Query("per_page", "10"))
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	maxYear, err := strconv.Atoi(c.Query("max_year", "2024"))
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	minYear, err := strconv.Atoi(c.Query("min_year", "0"))
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	sortBy := c.Query("sort_by", "id asc")
	searchBy := c.Query("search", "")
	searchBy = "%" + searchBy + "%"

	offset := 0
	if page > 0 {
		offset = (page - 1) * perPage
	}

	pagin := &domain.IngridientPagin{
		Pagin: &domain.Pagination{
			Offset: offset,
			Limit:  perPage,
			Search: searchBy,
			SortBy: sortBy},
		MaxYear: maxYear,
		MinYear: minYear,
	}

	Ingridients, err := handler.Srv.GetIngridients(pagin)
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.Status(200).JSON(Ingridients)
}

func (handler *Handler) SaveIngridient(c *fiber.Ctx) error {
	Ingridient := &domain.Ingridient{}
	err := json.Unmarshal(c.Body(), &Ingridient)
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	resp := handler.Srv.SaveIngridient(Ingridient)
	if err != nil {
		return c.Status(resp.StatusCode).JSON(resp.Message)
	}

	return c.SendStatus(resp.StatusCode)
}

func (handler *Handler) UpdateIngridient(c *fiber.Ctx) error {
	Ingridient := &domain.Ingridient{}

	id, err := c.ParamsInt("id")
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	err = json.Unmarshal(c.Body(), &Ingridient)
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	Ingridient.Id = uint(id)

	resp := handler.Srv.UpdateIngridient(Ingridient)
	if err != nil {
		return c.Status(resp.StatusCode).JSON(resp.Message)
	}

	return c.SendStatus(resp.StatusCode)
}

func (handler *Handler) GetIngridient(c *fiber.Ctx) error {
	Ingridient := &domain.Ingridient{}
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}
	Ingridient.Id = uint(id)

	Ingridient, resp := handler.Srv.GetIngridient(Ingridient)
	if resp.StatusCode == 400 {
		return c.Status(resp.StatusCode).JSON(resp.Message)
	}

	return c.Status(resp.StatusCode).JSON(Ingridient)
}

func (handler *Handler) DeleteIngridient(c *fiber.Ctx) error {
	Ingridient := &domain.Ingridient{}
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}
	Ingridient.Id = uint(id)

	resp := handler.Srv.DeleteIngridient(Ingridient)
	if resp.StatusCode == 400 {
		return c.Status(resp.StatusCode).JSON(resp.Message)
	}

	return c.Status(resp.StatusCode).JSON(resp.Message)
}

func (handler *Handler) GetIngridientsPerUni(c *fiber.Ctx) error {
	group := c.Query("group", "library")

	Ingridients, resp := handler.Srv.GetIngridientsPerUni(group)
	if resp.StatusCode == 400 {
		return c.Status(resp.StatusCode).JSON(resp.Message)
	}

	return c.Status(resp.StatusCode).JSON(Ingridients)
}
