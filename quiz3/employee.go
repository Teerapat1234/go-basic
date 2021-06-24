package main

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/shopspring/decimal"
)

type Employee struct {
	FirstName     string `json:"first_name" xml:"first_name" validate:"required"`
	LastName      string `json:"last_name" xml:"last_name" validate:"required"`
	EmployeeId    string `json:"employee_id" xml:"employee_id" validate:"required"`
	EmployeeSubId string `json:"employee_sub_id" xml:"employee_sub_id>sub1"`

	Age    uint8           `json:"age" xml:"age" validate:"required,gte=18,lte=60"`
	Height uint16          `json:"height" xml:"height" validate:"required"`
	Weight uint16          `json:"weight" xml:"weight" validate:"required"`
	Salary decimal.Decimal `json:"salary" xml:"salary"`
}

func (e *Employee) CheckJSONValidate(input []byte, vl *validator.Validate) error {

	if err := json.Unmarshal(input, e); err != nil {
		return err
	}

	if err := vl.Struct(e); err != nil {
		return err
	}
	return nil
}

func (e *Employee) CheckXMLValidate(input []byte, vl *validator.Validate) error {

	if err := xml.Unmarshal(input, e); err != nil {
		return err
	}

	//fmt.Printf("XML DATA : %#v", e)

	if e.FirstName == "" {
		return errors.New("first_name is required")
	}

	if e.LastName == "" {
		return errors.New("last_name is required")
	}

	if e.EmployeeId == "" {
		return errors.New("employee_id is required")
	}

	if err := vl.Var(e.Age, "gte=18,lte=60"); err != nil {
		fmt.Println("age : ", e.Age)
		return err
	}

	if err := vl.Var(e.Height, "gt=0"); err != nil {
		fmt.Println("height : ", e.Height)
		return err
	}

	if err := vl.Var(e.Weight, "gt=0"); err != nil {
		fmt.Println("weight : ", e.Weight)
		return err
	}

	return nil
}
