package entity

import (
	"fmt"
	"golang_training2/flp"
)

type Address struct {
	City string `json:"city"`
}

type Customer struct {
	Id        int     `json:"id"`
	FirstName string  `json:"fisrtname"`
	LastName  string  `json:"lastname"`
	Shipping  Address `json:"shipping"`
	Default   Address `json:"default"`
}

type Customers struct {
	Customers []Customer
}

func ViewCustomers() {
	var c2 Customers

	err := flp.ReadToStructTest("D:\\FILES\\CUSTOMERS.txt", &c2)
	//c, err := flp.ReadToStruct("D:\\FILES\\CUSTOMERS.txt", new(Customers))
	if err != "success" {
		println(err)
	} else {
		//c2 := c.(Customers)
		fmt.Println("|ID\t|Firstname\t|Lastname\t|Shipping\t|Default\t|")
		for _, v := range c2.Customers {
			fmt.Printf("|%d\t|%s\t|%s\t|%s\t|%s\t|\n", v.Id, v.FirstName, v.LastName, v.Shipping.City, v.Default.City)
		}
	}
}

func AddCustomer(id int, firstname string, lastname string, shipping string, def string) {
	var c2 Customers

	flp.ReadToStructTest("D:\\FILES\\CUSTOMERS.txt", &c2)

	c2.Customers = append(c2.Customers, Customer{id, firstname, lastname, Address{shipping}, Address{def}})
	flp.WriteData(c2, "D:\\FILES\\CUSTOMERS.txt")
}
