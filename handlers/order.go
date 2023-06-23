package handlers

import (
	"game-boost/contracts"
	"game-boost/services"

	"github.com/gofiber/fiber/v2"
)

func CreateOrder(c *fiber.Ctx) error {
	json := new(contracts.CreateOrderRequest)

	if err := c.BodyParser(json); err != nil {
		return c.JSON(fiber.Map{
			"code":    400,
			"message": "Invalid JSON",
			"error":   err.Error(),
		})
	}

	userID := c.Locals("userID")
	order, err := services.CreateOrder(json, userID.(string))
	if err != nil {
		return c.JSON(fiber.Map{
			"code":    400,
			"message": "Fail",
			"error":   err,
		})
	}

	return c.JSON(fiber.Map{
		"code":    200,
		"message": "Success",
		"order":   order,
	})
}

func EditOrder(c *fiber.Ctx) error {
	json := new(contracts.EditOrderRequest)
	orderID := c.Params("id")
	if err := c.BodyParser(json); err != nil {
		return c.JSON(fiber.Map{
			"code":    400,
			"message": "Invalid JSON",
			"error":   err.Error(),
		})
	}

	userID := c.Locals("userID")
	order, err := services.EditOrder(json, userID.(string), orderID)
	if err != nil {
		return c.JSON(fiber.Map{
			"code":    400,
			"message": "Fail",
			"error":   err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"code":          200,
		"message":       "Success",
		"updated order": order,
	})
}

func DeleteOrder(c *fiber.Ctx) error {
	orderID := c.Params("id")
	userID := c.Locals("userID")
	err := services.DeleteOrder(orderID, userID.(string))
	if err != nil {
		return c.JSON(fiber.Map{
			"code":    400,
			"message": "Fail",
			"error":   err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"code":    200,
		"message": "Order successfully deleted",
	})
}

func GetAllOrders(c *fiber.Ctx) error {
	userID := c.Locals("userID")
	orders, err := services.GetAllOrders(userID.(string))

	if err != nil {
		return c.JSON(fiber.Map{
			"code":    400,
			"message": "Fail",
			"Error":   err,
		})
	}

	return c.JSON(fiber.Map{
		"code":    200,
		"message": "Success",
		"Orders":  orders,
	})
}

func GetOrder(c *fiber.Ctx) error {
	orderID := c.Params("id")
	userID := c.Locals("userID")
	order, err := services.GetOrderByID(orderID, userID.(string))
	if err != nil {
		return c.JSON(fiber.Map{
			"code":    400,
			"message": "Fail",
			"error":   err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"code":    200,
		"message": "Success",
		"order":   order,
	})
}
