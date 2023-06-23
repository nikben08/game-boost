package services

import (
	"game-boost/contracts"
	"game-boost/models"
	"game-boost/repositories"

	"github.com/google/uuid"
)

func CreateOrder(json *contracts.CreateOrderRequest, userID string) (models.Order, error) {
	order := models.Order{
		UserId:          userID,
		OrderType:       json.OrderType,
		CurrentRank:     json.CurrentRank,
		CurrentLp:       json.CurrentLp,
		Server:          json.Server,
		DesiredRank:     json.DesiredRank,
		QType:           json.QType,
		ChampionsRoles:  json.ChampionsRoles,
		PriorityBoost:   json.PriorityBoost,
		StreamGames:     json.StreamGames,
		SoloOnly:        json.SoloOnly,
		BonusWin:        json.BonusWin,
		TotalPrice:      json.TotalPrice,
		WinsAmount:      json.WinsAmount,
		PlayWithBooster: json.PlaysWithBooster,
		PreviousRank:    json.PreviousRank,
		GameMode:        json.GameMode,
		GamesAmount:     json.GamesAmount,
	}
	createdOrder, err := repositories.CreateOrder(order)
	if err != nil {
		return models.Order{}, err
	}
	return createdOrder, nil
}

func GetAllOrders(userID string) ([]models.Order, error) {
	orders, err := repositories.GetAllOrders(userID)
	if err != nil {
		return []models.Order{}, err
	}

	return orders, nil
}

func GetOrderByID(orderID string, userID string) (models.Order, error) {
	order, err := repositories.GetOrderByID(orderID, userID)
	if err != nil {
		return models.Order{}, err
	}
	return order, nil
}

func DeleteOrder(orderID string, userID string) error {
	err := repositories.DeleteOrder(orderID, userID)
	if err != nil {
		return err
	}
	return nil
}

func EditOrder(json *contracts.EditOrderRequest, userID string, orderID string) (models.Order, error) {
	order := models.Order{
		ID:              uuid.MustParse(orderID),
		UserId:          userID,
		OrderType:       json.OrderType,
		CurrentRank:     json.CurrentRank,
		CurrentLp:       json.CurrentLp,
		Server:          json.Server,
		DesiredRank:     json.DesiredRank,
		QType:           json.QType,
		ChampionsRoles:  json.ChampionsRoles,
		PriorityBoost:   json.PriorityBoost,
		StreamGames:     json.StreamGames,
		SoloOnly:        json.SoloOnly,
		BonusWin:        json.BonusWin,
		TotalPrice:      json.TotalPrice,
		WinsAmount:      json.WinsAmount,
		PlayWithBooster: json.PlaysWithBooster,
		PreviousRank:    json.PreviousRank,
		GameMode:        json.GameMode,
		GamesAmount:     json.GamesAmount,
	}

	err := repositories.EditOrder(&order, orderID, userID)
	if err != nil {
		return models.Order{}, err
	}
	return order, nil
}
