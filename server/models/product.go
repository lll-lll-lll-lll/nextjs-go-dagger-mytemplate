package models

import (
	"database/sql"
	"errors"
)

type Product struct {
	ID int `json:"id"`
	Name string `json:"name"`
	Price int `json:"price"`
}


func (p *Product) getProduct(db *sql.DB) error {
	return errors.New("Not Implemented")
}

func (p *Product) updateProduct(db *sql.DB) error {
	return errors.New("Not Implemented")
}

func (p *Product) deleteProduct(db *sql.DB) error {
	return errors.New("Not implemented")
  }
  
func (p *Product) createProduct(db *sql.DB) error {
return errors.New("Not implemented")
}


func (p *Product) getProducts(db *sql.DB) ([]*Product, error) {
	return nil, errors.New("Not implemented")
}