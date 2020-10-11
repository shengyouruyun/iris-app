package repository

import (
	"fmt"
	"time"

	"github.com/iris-app/model"
)

var productList = []*model.Product{
	&model.Product{

		ID:          1,
		Name:        "Latte",
		Description: "Frothy milky coffe",
		Price:       2.45,
		SKU:         "abc232",
		CreateOn:    time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},

	&model.Product{

		ID:          2,
		Name:        "Espresso",
		Description: "Short and strong coffe without milk",
		Price:       1.48,
		SKU:         "arf232",
		CreateOn:    time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
}

type Products []*model.Product

func GetProducts() Products {

	return productList

}

func AddProduct(p *model.Product) {
	p.ID = getNextID()
	productList = append(productList, p)
}

func UpdateProduct(id int, p *model.Product) error {
	_, pos, err := findProduct(id)
	if err != nil {
		return err
	}

	p.ID = id
	productList[pos] = p

	return nil
}

//ErrProductNotFound defaut type use for erro type
var ErrProductNotFound = fmt.Errorf("Product not found")

func findProduct(id int) (*model.Product, int, error) {
	for i, p := range productList {
		if p.ID == id {
			return p, i, nil
		}
	}

	return nil, -1, ErrProductNotFound
}

func getNextID() int {
	lp := productList[len(productList)-1]
	return lp.ID + 1
}
