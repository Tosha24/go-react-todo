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

func AddTodo(c *fiber.Ctx) error {
	// parse the token from the request header
	token := string(c.Request().Header.Peek("Authorization"))

	// get userId from the token
	userId, err := utils.AuthenticateJWTToken(token)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "User not Authenticated",
		})
	}

	// take todo data from the body and keep in new todo
	newTodo := new(models.TODO)
	err = c.BodyParser(newTodo)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid JSON data",
		})
	}

	// create new id for todo
	newTodo.ID = primitive.NewObjectID().Hex()

	// convert userId string to object id
	userIdObj, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Invalid user Id",
		})
	}

	// get the user collection from the database
	collection := config.GetCollection(config.DB, "users")

	// create a filter to check if the user exists
	filter := bson.M{"_id": userIdObj}
	update := bson.M{"$push": bson.M{"todos": newTodo}}

	// update the user with the new todo
	_, err1 := collection.UpdateOne(context.Background(), filter, update)
	if err1 != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Cannot Add Todo to Database",
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"message": "Added todo to database",
		"todo":    newTodo,
	})
}

func GetAllTodos(c *fiber.Ctx) error {
	// parse the token from the request header
	token := string(c.Request().Header.Peek("Authorization"))

	// get userId from the token
	userId, err := utils.AuthenticateJWTToken(token)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "User not Authenticated",
		})
	}

	// convert userId string to object id
	userIdObj, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Invalid user Id",
		})
	}

	// get the user collection from the database
	collection := config.GetCollection(config.DB, "users")

	// create a filter to check if the user exists
	filter := bson.M{"_id": userIdObj}
	var user models.User

	// find the user with the userId
	err = collection.FindOne(context.Background(), filter).Decode(&user)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "User not found",
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"todos": user.Todos,
	})
}

func MarkAsCompleted(c *fiber.Ctx) error {
	// parse the token from the request header
	token := string(c.Request().Header.Peek("Authorization"))
	userId, err := utils.AuthenticateJWTToken(token)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "User not Authenticated",
		})
	}

	// convert userId string to object id
	userIdObj, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Invalid user Id",
		})
	}

	// get id from the params
	todoId := c.Params("id")

	collection := config.GetCollection(config.DB, "users")

	var user models.User
	filter := bson.M{"_id": userIdObj, "todos._id": todoId}
	err = collection.FindOne(context.Background(), filter).Decode(&user)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "User not found",
		})
	}

	var currentTodo *models.TODO

	for _, todo := range user.Todos {
		if todo.ID == todoId {
			currentTodo = &todo
			break
		}
	}

	if currentTodo == nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Todo not found",
		})
	}

	update := bson.M{"$set": bson.M{"todos.$.completed": !currentTodo.Completed}}

	_, err = collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Could not update todo",
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"message": "Todo updated",
	})
}

func UpdateTodo(c *fiber.Ctx) error {
	// get the token from the request header
	token := string(c.Request().Header.Peek("Authorization"))
	userId, err := utils.AuthenticateJWTToken(token)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "User not Authenticated",
		})
	}

	// convert userId string to object id
	userIdObj, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Invalid user Id",
		})
	}

	// get id from the params
	todoId := c.Params("id")

	collection := config.GetCollection(config.DB, "users")

	var user models.User
	filter := bson.M{"_id": userIdObj, "todos._id": todoId}
	err = collection.FindOne(context.Background(), filter).Decode(&user)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "User not found",
		})
	}

	var currentTodo *models.TODO

	for _, todo := range user.Todos {
		if todo.ID == todoId {
			currentTodo = &todo
			break
		}
	}

	if currentTodo == nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Todo not found",
		})
	}

	updatedTodo := new(models.TODO)
	err = c.BodyParser(updatedTodo)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid JSON data",
		})
	}

	update := bson.M{"$set": bson.M{"todos.$.title": updatedTodo.Title, "todos.$.completed": currentTodo.Completed}}

	_, err = collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Could not update todo",
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"message": "Todo updated",
		"todo":    updatedTodo,
	})
}

func DeleteTodo(c *fiber.Ctx) error {
	// get the token from the request header
	token := string(c.Request().Header.Peek("Authorization"))
	userId, err := utils.AuthenticateJWTToken(token)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "User not Authenticated",
		})
	}

	// convert userId string to object id
	userIdObj, err := primitive.ObjectIDFromHex(userId)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Invalid user Id",
		})
	}

	// get id from the params
	todoId := c.Params("id")

	collection := config.GetCollection(config.DB, "users")

	var user models.User
	filter := bson.M{"_id": userIdObj, "todos._id": todoId}
	err = collection.FindOne(context.Background(), filter).Decode(&user)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "User not found",
		})
	}

	var currentTodo *models.TODO

	for _, todo := range user.Todos {
		if todo.ID == todoId {
			currentTodo = &todo
			break
		}
	}

	if currentTodo == nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Todo not found",
		})
	}

	update := bson.M{"$pull": bson.M{"todos": bson.M{"_id": todoId}}}

	_, err = collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Could not delete todo",
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"message": "Todo deleted",
	})
}