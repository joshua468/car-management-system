package models

import (
	"errors"
	"strconv"
	"time"

	"github.com/google/uuid"
)

type Car struct {
	ID 		uuid.UUID       `json:"id"`
	Name	        string		`json:"name"`
	Year 	        string 		`json:"year"`
	Brand 	        string 		`json:"brand"`
	FuelType 	string 		`json:"fuel_type"`
	Engine 		Engine 		`json:"engine"`
	Price		float64		`json:"price"`
	CreatedAt	time.Time	`json:"created_at"`
	UpdatedAt	time.Time	`json:"updated_at"`
}

type CarRequest struct {
	Name	    string		`json:"name"`
	Year 	    string 		`json:"year"`
	Brand 	    string 		`json:"brand"`
	FuelType    string 		`json:"fuel_type"`
	Engine 	    Engine 		`json:"engine"`
	Price	   float64		`json:"price"`
}

func validate(name string)error {
if name ==  ""{
	return errors.New("Name is Required")
}
return nil
}
func validateYear(year string)error {
if year == ""{
	return errors.New("Year is Required")
}
_,err:=strconv.Atoi(year)
if err!=nil{
	return errors.New("year must be a valid number")
}
currentYear:= time.Now().Year()
yearInt,_:= strconv.Atoi(year)
if yearInt < 1886 || yearInt > currentYear {
	return errors.New("Year must be between 1886 and current year")
}
return nil
} 

func validatedBrand(brand string)error {
if brand == "" {
	return errors.New("Brand is required")
}
return nil
}

func  validateFuelType(fuelType string)error {
	validateFuelTypes:= []string{"Petrol","Diesel","Electric","Hybrid"}
	for _,validType := range validateFuelTypes {
		if fuelType == validType {
			return nil
		}
	}
	return errors.New("fuelType must be one of :Petrol,Diesel,Elctric,Hybrid")
}

func validateEngine(engine Engine)error {
	if engine.EngineID == uuid.Nil {
		return errors.New("EngineID is required")
	}
	if engine.Displacement <=0 {
		return errors.New("displacement must be greater than zero")
	}
	if engine.NoOfCylinders <=0 {
		return errors.New("noOfCylinders must be greater than zero")
	}
	if engine.CarRange<=0 {
		return errors.New("carRange must be greater than zero")
	}
	return nil
}
