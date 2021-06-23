package main

import (
	"fmt"
	"github.com/shopspring/decimal"
)

type Employee struct {
	FirstName          string
	LastName           string
	Weight             decimal.Decimal
	HeightInMeter      decimal.Decimal
	EmployeeId         string
	WeightFloat        float64
	HeightInMeterFloat float64
	BMI                decimal.Decimal
	BMIRankingStatus   string
}

func (e *Employee) PrintAll() {
	fmt.Printf("Employee =  %#v", e)
}

func (e *Employee) BMICalculator() {

	him2 := e.HeightInMeter.Mul(e.HeightInMeter)
	bmi := e.Weight.Div(him2)
	e.BMI = bmi
}

func (e *Employee) BMIRanking() {
	bmi := e.BMI

	if bmi.GreaterThanOrEqual(decimal.NewFromFloat(30.0)) {
		fmt.Println("Over-Fat = ", bmi)
		e.BMIRankingStatus = "Over-Fat"
	} else if bmi.GreaterThanOrEqual(decimal.NewFromFloat(29)) {
		fmt.Println("Fat = ", bmi)
		e.BMIRankingStatus = "Fat"
	} else if bmi.GreaterThanOrEqual(decimal.NewFromFloat(24)) {
		fmt.Println("Normal = ", bmi)
		e.BMIRankingStatus = "Normal"
	} else {
		fmt.Println("Thin = ", bmi)
		e.BMIRankingStatus = "Thin"
	}
}

func (e *Employee) PrintHealthStatus() {
	fmt.Println(e.BMIRankingStatus)
}
