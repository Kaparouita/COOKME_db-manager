package domain

import (
	"time"
)

type Ingridient struct {
	Id            uint `json:"id" gorm:"primaryKey"`
	CreatedAt     time.Time
	ISBN          string    `json:"isbn" gorm:"uniqueIndex"` // ISBN (un)
	Library       string    `json:"library"`
	Title         string    `json:"title"`
	Writer        string    `json:"writer"`
	Type          string    `json:"type"`
	PublishDate   time.Time `json:"publish_date"` // Publish release
	PagesNumber   uint      `json:"pages_number"`
	Photo         string    `json:"photo"`
	IngridientURL string    `json:"Ingridient_url"`
	Reviews       []Review  `json:"reviews" gorm:"foreignKey:IngridientID;constraint:OnDelete:CASCADE;"`
}

type IngridientResponse struct {
	Ingridients []Ingridient `json:"Ingridients"`
	Total       int64        `json:"total"`
}

type Review struct {
	Id           uint   `json:"id" gorm:"primaryKey"`
	IngridientID uint   `json:"Ingridient_id"`
	UserName     string `json:"user_name"`
	Rating       int    `json:"rating"`
	Comment      string `json:"comment"`
}

type IngridientsPerUni struct {
	Group            string `json:"group"`
	TotalIngridients uint   `json:"total_Ingridients"`
}

type IngridientPagin struct {
	Pagin   *Pagination `json:"pagination"`
	MaxYear int         `json:"max_year"`
	MinYear int         `json:"min_year"`
}
