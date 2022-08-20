package model

import "time"

// Defining schemata structs to use when querying

// The first table from the kingdom of Kod
// report_code | date | export_dest_code | export_dest_name | product_code | product_name | unit_price | quantity
// Needs to be normalized since one export can have multiple products

// 1st normal form (no items appear more than once per record)
// (1) (report_code | date | export_dest_code | export_dest_name) (2) (report_code | product_code | product_name | unit_price | quantity)
// Although this is better, the problem is that we cannot add a product if it has not be sold yet, so we need to normalize further (2)

// 2nd normal form (columns are determined when the PK is determined)
// (product_code | product_name |unit_price) (report_code | product_code | quantity)
// (2) was normalized, however, (1) eventhough we can determine the rest of the fields from pk report_code this is not beneficial since
// mixin of of kingdoms that do not have exports will not be able to be managed without having exports

// 3rd normal form (Table does not allow any non-PK to determine values from other columns)
// (report_code | date | export_dest_code) (export_dest_code | export_dest_name)

// Final result
// Sales: (report_code | date | export_dest_code)
// SalesStatement: (report_code | product_code | quantity)
// Product: (product_code | product_name |unit_price)
// Exports: (export_dest_code | export_dest_name)

type Sales struct {
	ReportID uint64    `json:"report_id"`
	Date     time.Time `json:"date"`
	ExportID uint64    `json:"export_id"`
}

type SalesStatement struct {
	ReportID  uint64 `json:"report_id"`
	ProductID uint64 `json:"product_id"`
	Quantity  uint64 `json:"quantity"`
}

type Product struct {
	ProductID   uint64  `json:"product_id"`
	ProductName string  `json:"product_name"`
	Price       float64 `json:"price"`
}

type Exports struct {
	ExportID   uint64 `json:"export_id"`
	ExportName string `json:"export_name"`
}
