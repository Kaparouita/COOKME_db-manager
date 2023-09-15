package models

import (
	"time"

	"github.com/Kaparouita/models/models"
)

type Order struct {
	Id              uint `json:"id" gorm:"primaryKey"`
	CreatedAt       time.Time
	IngredientRefer int               `json:"Ingredient_id"`
	Ingredient      models.Ingredient `gorm:"foreignKey:IngredientRefer"`
	UserRefer       int               `json:"user_id"`
	User            models.User       `gorm:"foreignKey:UserRefer"`
}
