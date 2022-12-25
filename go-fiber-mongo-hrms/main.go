package main

import (
	"context"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoInstance struct {
	Client *mongo.Client
	Db     *mongo.Database
}

var mg MongoInstance

const dbName = "fiber-hrms"
const mongoURI = "mongodb://localhost:27017/" + dbName

type Employee struct {
	ID     string  `json:"id,omitempty" bson:"_id,omitempty"`
	Name   string  `json:"name"`
	Salary float64 `json:"salary"`
	Age    float64 `json:"age"`
}

func Connect() error {
	client, err := mongo.NewClient(options.Client().ApplyURI(mongoURI))
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	db := client.Database(dbName)
	if err != nil {
		return err
	}
	mg = MongoInstance{
		Client: client,
		Db:     db,
	}
	return nil
}

func main() {
	if err := Connect(); err != nil {
		log.Fatal(err)
	}
	app := fiber.New()
	app.Get("/emplyee", func(c *fiber.Ctx) error {
		query := bson.D{{}}

		cursor, err := mg.Db.Collection("employees").Find(c.Context(), query)
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}
		var employees []Employee = make([]Employee, 0)
		err = cursor.All(c.Context(), &employees)
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}
		return c.JSON(employees)
	})
	app.Post("/emplyee", func(c *fiber.Ctx) error {
		collection := mg.Db.Collection("employees")
		employee := new(Employee)
		if err := c.BodyParser(employee); err != nil {
			return c.Status(400).SendString(err.Error())
		}
		employee.ID = ""
		insertionResult, err := collection.InsertOne(c.Context(), employee)
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}
		filter := bson.D{{Key: "_id", Value: insertionResult.InsertedID}}
		createdRecord := collection.FindOne(c.Context(), filter)
		createdEmployee := &Employee{}
		createdRecord.Decode(createdEmployee)
		return c.Status(201).JSON(createdEmployee)
	})
	app.Put("/emplyee/:id", func(c *fiber.Ctx) error {
		idParam := c.Params("id")
		emplyeeId, err := primitive.ObjectIDFromHex(idParam)
		if err != nil {
			return c.Status(400).SendString(err.Error())
		}
		employee := new(Employee)
		if err = c.BodyParser(employee); err != nil {
			return c.Status(400).SendString(err.Error())
		}
		query := bson.D{{Key: "_id", Value: emplyeeId}}
		update := bson.D{
			{
				Key: "$set",
				Value: bson.D{
					{Key: "name", Value: employee.Name},
					{Key: "age", Value: employee.Age},
					{Key: "salary", Value: employee.Salary},
				},
			},
		}
		mg.Db.Collection("employees").FindOneAndUpdate(c.Context(), query, update).Err()
		//  if err != nil{
		// 	if err == mongo.ErrNoDocuments{
		// 		return c.SendStatus(400)
		// 	}
		// 	return c.SendStatus(500)
		//  }
		employee.ID = idParam
		return c.Status(200).JSON(200)
	})
	app.Delete("/emplyee/:id", func(c *fiber.Ctx) error {
		employeeID, err := primitive.ObjectIDFromHex(c.Params("id"))
		if err != nil {
			return c.SendStatus(400)
		}
		query := bson.D{{Key: "_id", Value: employeeID}}
		result, err := mg.Db.Collection("employees").DeleteOne(c.Context(), query)
		if err != nil {
			return c.SendStatus(500)
		}
		if result.DeletedCount < 1 {
			return c.SendStatus(404)
		}
		return c.Status(200).JSON("record deleted")
	})
	log.Fatal(app.Listen(":3000"))
}
