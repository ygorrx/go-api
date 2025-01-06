package repository

import (
	"database/sql"
	"fmt"
	"go-api/model"
)

type ProductRepository struct {
	connection *sql.DB
}

func NewProductRepository(db *sql.DB) ProductRepository {
	return ProductRepository{
		connection: db,
	}
}

func (p *ProductRepository) GetProducts() ([]model.Product, error) {
	query := "SELECT id, product_name, price FROM product"
	rows, err := p.connection.Query(query)
	if err != nil {
		fmt.Println("Query execution error:", err)
		return nil, err
	}
	defer rows.Close()

	var productsList []model.Product

	for rows.Next() {
		var productObject model.Product
		err := rows.Scan(&productObject.ID, &productObject.Name, &productObject.Price)
		if err != nil {
			fmt.Println("Row scan error:", err)
			return nil, err
		}
		productsList = append(productsList, productObject)
	}

	if err = rows.Err(); err != nil {
		fmt.Println("Rows iteration error:", err)
		return nil, err
	}

	return productsList, nil
}

func (p *ProductRepository) CreateProduct(product model.Product) (int, error) {
	var id int
	query, err := p.connection.Prepare("INSERT INTO product" +
		"(product_name, price)" +
		" VALUES ($1, $2) RETURNING id")
	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	err = query.QueryRow(product.Name, product.Price).Scan(&id)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	query.Close()
	return id, nil
}

func (p *ProductRepository) GetProductById(id int) (*model.Product, error) {
	var product model.Product
	query := "SELECT * FROM product WHERE id = $1"
	err := p.connection.QueryRow(query, id).Scan(&product.ID, &product.Name, &product.Price)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return &product, nil
}