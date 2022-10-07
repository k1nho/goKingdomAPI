package model

import (
	"database/sql"
	"fmt"
)

func FetchProducts() ([]Product, error) {
	var products []Product
	var product Product

	sqlQuery := `SELECT * FROM products;`
	rows, err := DB.Query(sqlQuery)

	if err != nil {
		return products, err
	}

	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&product.ProductID, &product.ProductName, &product.Price)
		if err != nil {
			return products, err
		}
		products = append(products, product)
	}

	return products, nil
}

func FetchProduct(product_id uint64) (Product, error) {
	var product Product

	sqlQuery := `SELECT * FROM products WHERE product_id=$1`

	row := DB.QueryRow(sqlQuery, product_id)

	if err := row.Scan(&product.ProductID, &product.ProductName, &product.Price); err != nil {
		if err == sql.ErrNoRows {
      return product, fmt.Errorf("No product with id %d", product_id)
    }
    return product, fmt.Errorf("Product %d: %v", product_id, err)
	}

  return product, nil
}

func CreateProduct(product Product) error {

	sqlQuery := `INSERT INTO products(product_name, price) VALUES($1, $2);`

	_, err := DB.Exec(sqlQuery, product.ProductName, product.Price)
	if err != nil {
		return err
	}

	return nil
}

func UpdateProduct(product Product) error {
	sqlQuery := `UPDATE products SET product_name=$1 , price=$2 WHERE product_id=$3;`

	_, err := DB.Exec(sqlQuery, product.ProductName, product.Price, product.ProductID)
	if err != nil {
		return err
	}

	return nil
}

func DeleteProduct(id uint64) error {
	sqlQuery := `DELETE FROM products WHERE product_id=$1;`

	_, err := DB.Exec(sqlQuery, id)
	if err != nil {
		return err
	}

	return nil
}
