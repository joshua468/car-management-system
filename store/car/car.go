package car

import (
	"context"
	"database/sql"
	"github.com/joshua468/car-management-system/models"
)
type Store struct {
	db *sql.DB

}

func new(db *sql.DB)Store {
return Store{db:db}
}

func(s Store) GetCarById(ctx context.Context,id string) (models.Car,error) {
var car models.Car

query := `SELECT c.id,c.name,c.year,c.brand,c.fuel_type,c.engine_id,c.price,c.created_at,c.updated_at,e.id,e.displacement,e.no_of_cylinders,e.car_range FROM car c LEFT JOIN engine e ON c.engine_id = e.id WHERE c.id=$1 `

}

func(s Store) GetCarByBrand(ctx context.Context,brand string,isEngine bool) {

}

func(s Store) CreateCar(ctx context.Context,CarRequest *models.CarRequest) (models.Car,error) {

}

func(s Store) UpdateCar(ctx context.Context,id string,CarRequest *models.CarRequest) (models.Car,error) {

}

func (s Store) DeleteCar(ctx context.Context,id string)(models.Car,error) {

}

