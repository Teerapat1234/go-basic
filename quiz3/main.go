package main

import (
	"fmt"
	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

func main() {
	validate = validator.New()
	someoneJSON := []byte(`{
					"first_name":"Gopher",
					"last_name":"Conference",
					"age": 60,
					"employee_id":"999999",
					"height":169,
					"weight": 67,
					"salary": 100000.01
				}`)

	var empJson Employee

	if err := empJson.CheckJSONValidate(someoneJSON, validate); err != nil {
		fmt.Println("json validate error", err)
		return
	}

	fmt.Printf("JSON Employee OK: %#v", empJson)
	fmt.Println()

	someoneXML := []byte(`
							<employee>
							<first_name>Gopher2</first_name>
							<last_name>Cons2021</last_name>
							<employee_id>999999</employee_id>
							<employee_sub_id>
								<sub1>9999901</sub1>
							</employee_sub_id>
							<age>60</age>
                            <height>1</height>
							<weight>1</weight>
						  </employee>
						`)

	var empXml Employee

	if err := empXml.CheckXMLValidate(someoneXML, validate); err != nil {
		fmt.Println("xml validate error", err)
		return
	}

	fmt.Printf("XML Employee OK: %#v", empXml)
}
