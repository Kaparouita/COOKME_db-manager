package repositories

import (
	"github.com/Kaparouita/models/models"
)

func (db *Db) SaveIngredient(Ingredient *models.Ingredient) error {
	err := db.Model(Ingredient).Create(Ingredient).Error
	if err != nil {
		return err
	}
	return nil
}

func (db *Db) UpdateIngredient(Ingredient *models.Ingredient) error {
	return db.Model(Ingredient).Updates(Ingredient).Error
}

func (db *Db) GetIngredient(Ingredient *models.Ingredient) error {
	return db.Model(Ingredient).Preload("Reviews").Find(&Ingredient).Error
}

func (db *Db) DeleteIngredient(Ingredient *models.Ingredient) error {
	return db.Where("id = ?", Ingredient.Id).Delete(&models.Ingredient{}).Error
}

func (db *Db) GetIngredients() ([]models.Ingredient, error) {
	var Ingredient []models.Ingredient
	err := db.Find(&Ingredient).Error
	return Ingredient, err
}

//------------------------------Market Ingredient---------------------------//

func (db *Db) SaveMarketIngredient(MarketIngredient *models.MarketIngredient) error {
	err := db.Model(MarketIngredient).Create(MarketIngredient).Error
	if err != nil {
		return err
	}
	return nil
}

func (db *Db) UpdateMarketIngredient(MarketIngredient *models.MarketIngredient) error {
	return db.Model(MarketIngredient).Updates(MarketIngredient).Error
}

func (db *Db) GetMarketIngredient(MarketIngredient *models.MarketIngredient) error {
	return db.Model(MarketIngredient).Preload("MarketMarketIngredients").Find(&MarketIngredient).Error
}

func (db *Db) DeleteMarketIngredient(MarketIngredient *models.MarketIngredient) error {
	return db.Where("id = ?", MarketIngredient.Id).Delete(&models.MarketIngredient{}).Error
}

func (db *Db) GetMarketIngredients() ([]models.MarketIngredient, error) {
	var MarketIngredient []models.MarketIngredient
	err := db.Find(&MarketIngredient).Error
	return MarketIngredient, err
}
