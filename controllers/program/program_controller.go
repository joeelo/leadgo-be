package controllers

import (
	"context"
	"fiber-mongo-api/configs"
	"fiber-mongo-api/models"
	"fiber-mongo-api/responses"
	"net/http"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var programCollection *mongo.Collection = configs.GetCollection(configs.DB, "programs")
var validate = validator.New()

func CreateProgram(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var program models.Program
	defer cancel()

	//validate the request body
	if err := c.BodyParser(&program); err != nil {
		return c.Status(http.StatusBadRequest).JSON(responses.ProgramResponse{Status: http.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	//use the validator library to validate required fields
	if validationErr := validate.Struct(&program); validationErr != nil {
		return c.Status(http.StatusBadRequest).JSON(responses.ProgramResponse{Status: http.StatusBadRequest, Message: "error", Data: &fiber.Map{"data": validationErr.Error()}})
	}

	newProgram := models.Program{
		Id:   primitive.NewObjectID(),
		Name: program.Name,
	}

	result, err := programCollection.InsertOne(ctx, newProgram)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.ProgramResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	return c.Status(http.StatusCreated).JSON(responses.ProgramResponse{Status: http.StatusCreated, Message: "success", Data: &fiber.Map{"data": result}})
}

func GetAProgram(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	q := c.Queries()
	programId := q["programId"]
	var program models.Program
	defer cancel()

	objId, _ := primitive.ObjectIDFromHex(programId)

	err := programCollection.FindOne(ctx, bson.M{"_id": objId}).Decode(&program)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.ProgramResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	return c.Status(http.StatusOK).JSON(responses.ProgramResponse{Status: http.StatusOK, Message: "success", Data: &fiber.Map{"data": program}})
}

func GetAllPrograms(c *fiber.Ctx) error {
	log.Print("hello world from get all programs")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var programs []models.Program
	defer cancel()

	results, err := programCollection.Find(ctx, bson.M{})

	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.ProgramResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	//reading from the db in an optimal way
	defer results.Close(ctx)

	for results.Next(ctx) {
		var singleProgram models.Program
		if err = results.Decode(&singleProgram); err != nil {
			return c.Status(http.StatusInternalServerError).JSON(responses.ProgramResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"data": err.Error()}})
		}

		programs = append(programs, singleProgram)
	}

	return c.Status(http.StatusOK).JSON(
		responses.ProgramResponse{Status: http.StatusOK, Message: "success", Data: &fiber.Map{"data": programs}},
	)
}
