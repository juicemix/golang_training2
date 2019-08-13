package main

import (
	"golang_training2/entity"
)

func main() {
	entity.ViewCustomers()
	entity.AddCustomer(1, "rendy", "muhardianto", "jakarta", "malang")
	entity.ViewCustomers()
}
