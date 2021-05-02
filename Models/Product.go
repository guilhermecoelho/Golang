package model

import uuid "github.com/satori/go.uuid"

type Product struct {
	Id   string
	Name string
}

type Products struct {
	Product []Product
}

func (p *Products) Add(product *Product) {
	p.Product = append(p.Product, *product)
}

func NewProduct() *Product {
	return &Product{
		Id: uuid.NewV4().String(),
	}
}
