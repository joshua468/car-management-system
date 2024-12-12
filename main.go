package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/joshua468/car-management-system/carZone/driver"
	carHandler "github.com/joshua468/car-management-system/carZone/handler/car"
	engineHandler "github.com/joshua468/car-management-system/carZone/handler/engine"
	carService "github.com/joshua468/car-management-system/carZone/service/car"
	engineService "github.com/joshua468/car-management-system/carZone/service/engine"
	carStore "github.com/joshua468/car-management-system/carZone/store/car"
	engineStore "github.com/joshua468/car-management-system/carZone/store/engine"
)

func main() {
	// Load environment variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file. Please ensure it exists and contains the required variables.")
	}

	// Initialize database connection
	driver.InitDB()
	defer driver.CloseDB()

	db := driver.GetDB()

	// Set up stores and services
	carStore := carStore.New(db)
	carService := carService.NewCarService(carStore)

	engineStore := engineStore.New(db)
	engineService := engineService.NewEngineService(engineStore)

	// Set up handlers
	carHandlerInstance := carHandler.NewCarHandler(carService)
	engineHandlerInstance := engineHandler.NewEngineHandler(engineService)

	// Set up router and execute schema
	router := mux.NewRouter()
	// Adjust this to the absolute path of your schema.sql file
	schemaFile := "store/schema.sql"
	if err := executiveSchemaFile(db, schemaFile); err != nil {
		log.Println("Error while executing the schema file", err)
	}

	// Set up routes
	router.HandleFunc("/cars/{id}", carHandlerInstance.GetCarByID).Methods("GET")
	router.HandleFunc("/cars", carHandlerInstance.GetCarByBrand).Methods("GET")
	router.HandleFunc("/cars", carHandlerInstance.CreateCar).Methods("POST")
	router.HandleFunc("/cars/{id}", carHandlerInstance.UpdateCar).Methods("PUT")
	router.HandleFunc("/cars/{id}", carHandlerInstance.DeleteCar).Methods("DELETE")

	router.HandleFunc("/engine/{id}", engineHandlerInstance.GetEngineByID).Methods("GET")
	router.HandleFunc("/engine", engineHandlerInstance.CreateEngine).Methods("POST")
	router.HandleFunc("/engine/{id}", engineHandlerInstance.UpdateEngine).Methods("PUT")
	router.HandleFunc("/engine/{id}", engineHandlerInstance.DeleteEngine).Methods("DELETE")

	// Start the server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	addr := fmt.Sprintf(":%s", port)
	log.Printf("server is listening on %s:", addr)
	log.Fatal(http.ListenAndServe(addr, router))
}

func executiveSchemaFile(db *sql.DB, fileName string) error {
	sqlFile, err := os.ReadFile(fileName)
	if err != nil {
		return err
	}

	log.Println("Executing schema file:", fileName)
	_, err = db.Exec(string(sqlFile))
	if err != nil {
		return err
	}
	return nil
}
