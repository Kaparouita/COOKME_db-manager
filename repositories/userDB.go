package repositories

import "github.com/Kaparouita/models/models"

func (db *Db) SaveUser(user *models.User) error {
	err := db.Model(user).Create(user).Error
	if err != nil {
		return err
	}
	return nil
}

func (db *Db) UpdateUser(user *models.User) error {
	return db.Model(user).Updates(user).Error
}

func (db *Db) GetUsers() ([]models.User, error) {
	var users []models.User
	err := db.Find(&users).Error
	return users, err
}

func (db *Db) GetUser(user *models.User) error {
	return db.Model(user).Preload("Address").Preload("Paso").Find(user).Error
}

func (db *Db) DeleteUser(user *models.User) error {
	return db.Delete(&models.User{Id: user.Id}).Error
}

func (db *Db) Login(login *models.LoginResp) (*models.User, error) {
	user := &models.User{}
	err := db.
		Model(&models.User{}).
		Preload("Address").
		Preload("Paso").
		Where("username = ?", login.Username).
		Where("password = ?", login.Password).
		Find(&user).Error
	return user, err
}
