package products

import (
	"github.com/k1nho/goKingdomAPI/model"
)

func FetchProducts() ([]model.Product, error) {
	var products []model.Product
	var product model.Product

	sqlQuery := `SELECT * FROM products;`
	rows, err := model.DB.Query(sqlQuery)

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

func FetchProduct(product_id int) (model.Product, error) {
	var product model.Product

	sqlQuery := `SELECT * FROM products WHERE id=$1`
	row := model.DB.QueryRow(sqlQuery, product_id)
	if row == nil {
		return product, row.Err()
	}

	row.Scan(&product.ProductID, &product.ProductName, &product.Price)

	return product, nil
}
