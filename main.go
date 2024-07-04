package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Todo represents a task with an ID, completion status, and body content.
type Todo struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Completed bool               `json:"completed"`
	Body      string             `json:"body"`
}

var collection *mongo.Collection

func main() {
	fmt.Println("hello World")

	if os.Getenv("ENV") != "production" {
		// Load environment variables if not in production
		err := godotenv.Load(".env")
		if err != nil {
			log.Fatal("Error loading .env file:", err)
		}
	}

	// Get the MongoDB URI from the environment variables
	MONGODB_URI := os.Getenv("MONGODB_URI")

	// Set client options for connecting to MongoDB
	clientOptions := options.Client().ApplyURI(MONGODB_URI)

	// Connect to MongoDB
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Schedule the disconnection from MongoDB once the main function completes
	defer client.Disconnect(context.Background())

	// Ping the MongoDB server to verify the connection
	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}

	// Print a message indicating successful connection to MongoDB Atlas
	fmt.Println("Connected to MongoDB Atlas")

	// Get a handle for the 'todos' collection in the 'golang_db' database
	collection = client.Database("golang_db").Collection("todos")

	// Initialize a new Fiber app instance
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:5173",
		AllowHeaders: "Origin,Content-Type,Accept",
	}))

	// Define route handlers for the 'todos' API
	app.Get("/api/todos", getTodos)
	app.Post("/api/todos", createTodos)
	app.Patch("/api/todos/:id", updateTodos)
	app.Delete("/api/todos/:id", deleteTodos)

	// Get the port number from the environment variable, default to 5000 if not set
	port := os.Getenv("PORT")
	if port == "" {
		port = "5000"
	}

	if os.Getenv("ENV") == "production" {
		app.Static("/", "./client/dist")
	}
	// Start the Fiber app and listen on the specified port
	// Log and terminate the application if there is an error during startup
	log.Fatal(app.Listen("0.0.0.0:" + port))

}

// Handler function to retrieve todos
func getTodos(c *fiber.Ctx) error {
	var todos []Todo

	// Find all documents in the 'todos' collection
	cursor, err := collection.Find(context.Background(), bson.M{})
	if err != nil {
		return err
	}

	// Ensure the cursor is closed once all operations are done
	defer cursor.Close(context.Background())

	// Iterate over the cursor to decode each document into a Todo struct
	for cursor.Next(context.Background()) {
		var todo Todo
		// Decode the current document into a Todo struct
		if err := cursor.Decode(&todo); err != nil {
			return err
		}
		// Append the decoded Todo to the todos slice
		todos = append(todos, todo)
	}

	return c.JSON(todos)
}

// Handler function to create a new todo
func createTodos(c *fiber.Ctx) error {
	todo := new(Todo)
	// {id:0,completed:false,body:""}

	if err := c.BodyParser(todo); err != nil {
		return err
	}

	if todo.Body == "" {
		return c.Status(400).JSON(fiber.Map{"error": "Todo body cannot be empty"})
	}

	insertResult, err := collection.InsertOne(context.Background(), todo)
	if err != nil {
		return err
	}

	todo.ID = insertResult.InsertedID.(primitive.ObjectID)

	return c.Status(201).JSON(todo)
}

// Handler function to update an existing todo
func updateTodos(c *fiber.Ctx) error {
	id := c.Params("id")
	objectId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid todo id"})
	}

	filter := bson.M{"_id": objectId}
	update := bson.M{"$set": bson.M{"completed": true}}

	_, err = collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return err
	}

	return c.Status(200).JSON(fiber.Map{"success": "true"})
}

// Handler function to delete a todo
func deleteTodos(c *fiber.Ctx) error {
	id := c.Params("id")
	objectID, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid todo ID"})
	}

	filter := bson.M{"_id": objectID}
	_, err = collection.DeleteOne(context.Background(), filter)
	if err != nil {
		return err
	}

	return c.Status(200).JSON(fiber.Map{"success": true})
}
