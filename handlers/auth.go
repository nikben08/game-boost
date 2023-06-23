package handlers

import (
	"game-boost/contracts"
	"game-boost/services"

	"github.com/gofiber/fiber/v2"
)

func Signin(c *fiber.Ctx) error {
	json := new(contracts.SigninRequest)

	if err := c.BodyParser(json); err != nil {
		return c.JSON(fiber.Map{
			"code":    400,
			"message": "Invalid JSON",
			"error":   err.Error(),
		})
	}

	if len(json.Email) < 5 {
		var response = contracts.AuthResponse{
			Code:    400,
			Message: "Incorrect email",
		}
		return c.JSON(response)
	}

	if len(json.Password) < 8 {
		var response = contracts.AuthResponse{
			Code:    400,
			Message: "Incorrect password",
		}
		return c.JSON(response)
	}

	token, err := services.Signin(json)

	if err != nil {
		var response = contracts.AuthResponse{
			Code:    400,
			Message: err.Error(),
		}
		return c.JSON(response)
	}

	var response = contracts.AuthResponse{
		Code:    200,
		Message: "User successfully logged in",
		Token:   token,
	}

	return c.JSON(response)

}

func Signup(c *fiber.Ctx) error {
	json := new(contracts.SignupRequest)

	if err := c.BodyParser(json); err != nil {
		return c.JSON(fiber.Map{
			"code":    400,
			"message": "Invalid JSON",
			"error":   err.Error(),
		})
	}

	if json.Password != json.PasswordRepeat {
		var response = contracts.AuthResponse{
			Code:    400,
			Message: "Passwords do not match",
		}
		return c.JSON(response)
	}

	if len(json.Email) < 5 {
		var response = contracts.AuthResponse{
			Code:    400,
			Message: "Incorrect email",
		}
		return c.JSON(response)
	}

	if len(json.Name) < 2 {
		var response = contracts.AuthResponse{
			Code:    400,
			Message: "Incorrect Name",
		}
		return c.JSON(response)
	}

	token, err := services.Signup(json)

	if err != nil {
		var response = contracts.AuthResponse{
			Code:    400,
			Message: err.Error(),
		}
		return c.JSON(response)
	}

	var response = contracts.AuthResponse{
		Code:    200,
		Message: "User successfully created",
		Token:   token,
	}

	return c.JSON(response)
}
