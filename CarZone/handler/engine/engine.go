package models

import (
	"errors"

	"github.com/google/uuid"
)

type Engine struct {
	EngineID      uuid.UUID `json:"engine_id"`
	Displacement  int64     `json:"displacement"`
	NoOfCylinders int64     `json:"noOfCyclinders"`
	CarRange      int64     `json:"car_range"`
}

type EngineRequest struct {
	Displacement  int64 `json:"displacement"`
	NoOfCylinders int64 `json:"noOfCyclinders"`
	CarRange      int64 `json:"car_range"`
}

func ValidateEngineRequest(EngineReq Engine) error {
	if err := validateDisplacement(EngineReq.Displacement); err != nil {
		return err
	}
	if err := validateNoOfCylinders(EngineReq.NoOfCylinders); err != nil {
		return err
	}
	if err := validateCarRange(EngineReq.CarRange); err != nil {
		return err
	}
	return nil
}

func validateDisplacement(displacement int64) error {
	if displacement <= 0 {
		return errors.New("displacement  must be greater than zero")
	}
	return nil
}

func validateNoOfCylinders(noOfCylinders int64) error {
	if noOfCylinders <= 0 {
		return errors.New("noOfCylinders must be greater than zero")
	}
	return nil
}

func validateCarRange(carRange int64) error {
	if carRange <= 0 {
		return errors.New("carRange Must be Greater than Zero")
	}
	return nil
}
