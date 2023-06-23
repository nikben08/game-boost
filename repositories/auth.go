package repositories

import (
	"game-boost/models"
)

func WhetherUserExists(email string) bool {
	var user models.User
	if result := DB.Where("email = ?", email).First(&user); result.Error != nil {
		return false
	}
	return true
}

func CreateUser(user models.User) (models.User, error) {
	if result := DB.Create(&user); result.Error != nil {
		return models.User{}, result.Error
	}
	return user, nil
}

func GetUserByEmail(email string) (models.User, error) {
	var user models.User
	if result := DB.Where("email = ?", email).First(&user); result.Error != nil {
		return user, result.Error
	}
	return user, nil
}
