package model

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

	sqlQuery := `SELECT * FROM products WHERE id=$1`
	row := DB.QueryRow(sqlQuery, product_id)
	if row == nil {
		return product, row.Err()
	}

	row.Scan(&product.ProductID, &product.ProductName, &product.Price)

	return product, nil
}
