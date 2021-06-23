package main

import "github.com/shopspring/decimal"

func main() {
	emp := Employee{
		FirstName:          "Gopher",
		LastName:           "Cons2020",
		Weight:             decimal.NewFromFloat(83),
		HeightInMeter:      decimal.NewFromFloat(1.67),
		EmployeeId:         "999999",
		WeightFloat:        75.0,
		HeightInMeterFloat: 1.67,
	}

	emp.PrintAll()
	emp.BMICalculator()
	emp.BMIRanking()
	emp.PrintHealthStatus()

}
