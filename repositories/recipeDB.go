package repositories

import "github.com/Kaparouita/models/models"

func (db *Db) SaveRecipe(Recipe *models.Recipe) error {
	err := db.Model(Recipe).Create(Recipe).Error
	if err != nil {
		return err
	}
	return nil
}

func (db *Db) UpdateRecipe(Recipe *models.Recipe) error {
	return db.Model(Recipe).Updates(Recipe).Error
}

func (db *Db) GetRecipe(Recipe *models.Recipe) error {
	return db.Model(Recipe).Preload("FinalRecipes").Find(&Recipe).Error
}

func (db *Db) DeleteRecipe(Recipe *models.Recipe) error {
	return db.Where("id = ?", Recipe.Id).Delete(&models.Recipe{}).Error
}

func (db *Db) GetRecipes() ([]models.Recipe, error) {
	var Recipe []models.Recipe
	err := db.Find(&Recipe).Error
	return Recipe, err
}

//---------------------------FinalRecipe-------------------------------//

func (db *Db) SaveFinalRecipe(FinalRecipe *models.FinalRecipe) error {
	err := db.Model(FinalRecipe).Create(FinalRecipe).Error
	if err != nil {
		return err
	}
	return nil
}

func (db *Db) UpdateFinalRecipe(FinalRecipe *models.FinalRecipe) error {
	return db.Model(FinalRecipe).Updates(FinalRecipe).Error
}

func (db *Db) GetFinalRecipe(FinalRecipe *models.FinalRecipe) error {
	return db.Model(FinalRecipe).Preload("MarketFinalRecipes").Find(&FinalRecipe).Error
}

func (db *Db) DeleteFinalRecipe(FinalRecipe *models.FinalRecipe) error {
	return db.Where("id = ?", FinalRecipe.Id).Delete(&models.FinalRecipe{}).Error
}

func (db *Db) GetFinalRecipes() ([]models.FinalRecipe, error) {
	var FinalRecipe []models.FinalRecipe
	err := db.Find(&FinalRecipe).Error
	return FinalRecipe, err
}
