package controller

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/tosha24/todo/config"
	"github.com/tosha24/todo/models"
	"github.com/tosha24/todo/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func Login(c *fiber.Ctx) error {
	// get the user data from the c body
	data := new(models.User)
	err := c.BodyParser(data)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid data",
		})
	}

	collection := config.GetCollection(config.DB, "users")

	filter := bson.D{{Key: "email", Value: data.Email}}

	var result models.User
	err = collection.FindOne(context.Background(), filter).Decode(&result)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "User not found",
		})
	}

	// compare the password with the hashed password
	if !utils.ComparePasswords(result.Password, data.Password) {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid password",
		})
	}

	// generate a jwt token for the user with that userId
	token, err := utils.GenerateJWTToken(result.ID)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Couldn't generate token",
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"token": token,
	})
}

func CreateUser(c *fiber.Ctx) error {
	user := new(models.User)
	err := c.BodyParser(user)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid data",
		})
	}

	collection := config.GetCollection(config.DB, "users")
	filter := bson.D{{Key: "email", Value: user.Email}}

	count, _ := collection.CountDocuments(context.Background(), filter)
	if count > 0 {
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{
			"error": "User already exists",
		})
	}

	user.Password = utils.HashPassword(user.Password)
	user.Todos = []models.TODO{}

	newUser, err := collection.InsertOne(context.Background(), user)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Couldn't create user",
		})
	}

	userId := newUser.InsertedID.(primitive.ObjectID).Hex()

	// generate a jwt token for the user with that userId
	token, err := utils.GenerateJWTToken(userId)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Couldn't generate token",
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"token": token,
	})
}