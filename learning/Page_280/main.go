package main

import (
	"fmt"
	"learning/magazine"
)

func main() {
	/*var s magazine.Subscriber
	s.Rate = 4.99
	fmt.Println(s.Rate)
	var employee magazine.Employee
	employee.Name = "Joy Carr"
	employee.Salary = 60000
	fmt.Println(employee.Name)
	fmt.Println(employee.Salary)*/
	subscriber := magazine.Subscriber{Name: "Aman Singh"}
	subscriber.HomeAddress.Street = "123 Oak St"
	subscriber.HomeAddress.City = "Omaha"
	subscriber.HomeAddress.State = "Ne"
	subscriber.HomeAddress.PostalCode = "68111"
	fmt.Println("Subscriber name:", subscriber.Name)
	fmt.Println("Street:", subscriber.HomeAddress.Street)
	fmt.Println("City:", subscriber.HomeAddress.City)
	fmt.Println("State:", subscriber.HomeAddress.State)
	fmt.Println("Postal Code:", subscriber.HomeAddress.PostalCode)
	employee := magazine.Employee{Name: "Joy Carr"}
	employee.HomeAddress.Street = "456 Elm St"
	employee.HomeAddress.City = "Portland"
	employee.HomeAddress.State = "Or"
	employee.HomeAddress.PostalCode = "97222"
	fmt.Println("Employee name:", employee.Name)
	fmt.Println("Street:", employee.HomeAddress.Street)
	fmt.Println("City:", employee.HomeAddress.City)
	fmt.Println("State:", employee.HomeAddress.State)
	fmt.Println("Postal Code:", employee.HomeAddress.PostalCode)
}
