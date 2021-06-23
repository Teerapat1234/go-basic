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
}

func (e Employee) PrintAll() {
	fmt.Printf("Employee =  %#v", e)
}

func (e Employee) BMICalculator() decimal.Decimal {

	him2f := e.HeightInMeterFloat * e.HeightInMeterFloat

	bmif := e.WeightFloat / him2f
	fmt.Println()
	fmt.Println("BMI Float= ", bmif)

	him2 := e.HeightInMeter.Mul(e.HeightInMeter)
	bmi := e.Weight.Div(him2)
	fmt.Println()
	fmt.Println("BMI = ", bmi)
	return bmi
}

func (e Employee) BMIRanking() string {
	bmi := e.BMICalculator()

	if bmi.GreaterThanOrEqual(decimal.NewFromFloat(30.0)) {
		fmt.Println("Over-Fat = ", bmi)
		return "Over-Fat"
	} else if bmi.GreaterThanOrEqual(decimal.NewFromFloat(29)) {
		fmt.Println("Fat = ", bmi)
		return "Fat"
	} else if bmi.GreaterThanOrEqual(decimal.NewFromFloat(24)) {
		fmt.Println("Normal = ", bmi)
		return "Normal"
	} else {
		fmt.Println("Thin = ", bmi)
		return "Thin"
	}
}

func (e Employee) PrintHealthStatus() {
	fmt.Println(e.BMIRanking())
}
