package main

import (
	"Loan_Tracker_API/database"
	"Loan_Tracker_API/delivery/controller"
	"Loan_Tracker_API/delivery/router"
	"Loan_Tracker_API/repository"
	"Loan_Tracker_API/usecase"
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	// MongoDB URI
	mongoURI := "mongodb://localhost:27017"

	// Connect to MongoDB
	clientOptions := options.Client().ApplyURI(mongoURI)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}

	// Create a collection interface
	collection := &database.MongoCollection{Collection: client.Database("loan_tracker").Collection("loans")}

	// Create repositories
	loanRepository := repository.NewLoanRepository(collection)

	// Create usecases
	loanUsecase := usecase.NewLoanUseCase(loanRepository)

	// Create controllers
	loanController := controller.NewLoanController(loanUsecase, nil) // Assuming userUsecase is not needed

	// Setup router
	r := router.SetupRouter(loanController)

	// Start the server
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
