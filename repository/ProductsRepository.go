package repository

import(
	
	"time"
	"github.com/iris-app/model"
	
	
	) 

var productlist = []*Product{
	&Product{

		ID:          1,
		Name:        "Latte",
		Description: "Frothy milky coffe",
		Price:       2.45,
		SKU:         "abc232",
		CreateOn:    time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},

	&Product{

		ID:          2,
		Name:        "Espresso",
		Description: "Short and strong coffe without milk",
		Price:       1.48,
		SKU:         "arf232",
		CreateOn:    time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
}



type Products []*Product


