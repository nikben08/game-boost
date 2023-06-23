package services

import (
	"errors"
	"game-boost/contracts"
	"game-boost/models"
	"game-boost/repositories"
	encryption "game-boost/utils/encryption"
	jwt "game-boost/utils/jwt"
)

func Signup(json *contracts.SignupRequest) (string, error) {

	whetherUserExists := repositories.WhetherUserExists(json.Email)

	if whetherUserExists {
		return "", errors.New("Email has already used")
	}

	var newUser = models.User{
		Email:    json.Email,
		Password: encryption.GenerateHash(json.Password),
		Name:     json.Name,
	}

	user, err := repositories.CreateUser(newUser)

	if err != nil {
		return "", err
	}

	token, err := jwt.GenerateJwtToken(user.ID, user.Email, user.Name)

	if err != nil {
		return "", err
	}

	return token, nil
}

func Signin(json *contracts.SigninRequest) (string, error) {

	user, err := repositories.GetUserByEmail(json.Email)

	if err != nil || encryption.GenerateHash(json.Password) != user.Password {
		return "", errors.New("Wrong e-mail or password")
	}

	token, err := jwt.GenerateJwtToken(user.ID, user.Email, user.Name)

	if err != nil {
		return "", err
	}

	return token, nil
}
