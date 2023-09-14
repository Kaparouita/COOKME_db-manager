package domain

import "time"

type Order struct {
	Id              uint `json:"id" gorm:"primaryKey"`
	CreatedAt       time.Time
	IngridientRefer int        `json:"Ingridient_id"`
	Ingridient      Ingridient `gorm:"foreignKey:IngridientRefer"`
	UserRefer       int        `json:"user_id"`
	User            User       `gorm:"foreignKey:UserRefer"`
}
