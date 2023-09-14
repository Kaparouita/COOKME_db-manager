package repositories

import "github.com/Kaparouita/db-manager/domain"

func (db *Db) SaveIngridient(Ingridient *domain.Ingridient) error {
	err := db.Model(Ingridient).Create(Ingridient).Error
	if err != nil {
		return err
	}
	return nil
}

func (db *Db) GetIngridients(pagin *domain.IngridientPagin) (*domain.IngridientResponse, error) {
	Ingridients := &domain.IngridientResponse{}
	err := db.
		Model(&domain.Ingridient{}).
		Preload("Reviews").
		Order(pagin.Pagin.SortBy).
		Offset(pagin.Pagin.Offset).
		Limit(pagin.Pagin.Limit).
		Where("type ilike ?", pagin.Pagin.Search).
		Where("extract(year from publish_date) <= ?", pagin.MaxYear).
		Where("extract(year from publish_date) >= ?", pagin.MinYear).
		Find(&Ingridients.Ingridients).Error
	if err != nil {
		return nil, err
	}

	err = db.
		Model(&domain.Ingridient{}).
		Where("type ilike ?", pagin.Pagin.Search).
		Where("extract(year from publish_date) <= ?", pagin.MaxYear).
		Where("extract(year from publish_date) >= ?", pagin.MinYear).
		Count(&Ingridients.Total).Error
	if err != nil {
		return nil, err
	}

	return Ingridients, nil
}

func (db *Db) UpdateIngridient(Ingridient *domain.Ingridient) error {
	return db.Model(Ingridient).Updates(Ingridient).Error
}

func (db *Db) GetIngridient(Ingridient *domain.Ingridient) error {
	return db.Model(Ingridient).Preload("Reviews").Find(&Ingridient).Error
}

func (db *Db) DeleteIngridient(Ingridient *domain.Ingridient) error {
	return db.Where("id = ?", Ingridient.Id).Delete(&domain.Ingridient{}).Error
}

func (db *Db) GetIngridientsPerUni(groupField string) ([]domain.IngridientsPerUni, error) {
	var results []domain.IngridientsPerUni

	err := db.
		Model(&domain.Ingridient{}).
		Select(groupField+" as group", "COUNT(*) as total_Ingridients").
		Group(groupField).
		Scan(&results).
		Error

	if err != nil {
		return nil, err
	}

	return results, nil
}
