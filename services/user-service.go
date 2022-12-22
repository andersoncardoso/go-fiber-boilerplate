package services

import (
	"fmt"
	"myapp/models"
)

func UserList() ([]models.User, error) {
	var users []models.User

	tx := db.Find(&users)
	if tx.Error != nil {
		return []models.User{}, tx.Error
	}
	return users, nil
}

func UserFind(id uint64) (models.User, error) {
	var user models.User

	tx := db.Where("id = ?", id).First(&user)
	if tx.Error != nil {
		return models.User{}, fmt.Errorf("Error: could not find id. " + tx.Error.Error())
	}
	return user, nil
}

func UserCreate(user *models.User) error {
	tx := db.Create(&user)
	return tx.Error
}

func UserUpdate(user *models.User) error {
	tx := db.Save(&user)
	return tx.Error
}

func UserDelete(id uint64) error {
	// Gorm does soft deletes, so we need Unscoped to do an actual deletion
	tx := db.Unscoped().Delete(&models.User{}, id)
	return tx.Error
}

func UserFindByEmail(email string) (models.User, error) {
	var user models.User
	tx := db.Where("email = ?", email).First(&user)
	return user, tx.Error
}
