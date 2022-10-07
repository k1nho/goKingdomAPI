package model

import (
	"database/sql"
	"fmt"
)

func GetExports() ([]Exports, error) {
	sqlQuery := `SELECT * FROM exports;`

	var exports []Exports
	var export Exports

	rows, err := DB.Query(sqlQuery)
	if err != nil {
		return exports, err
	}

	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&export.ExportID, &export.ExportName)
		if err != nil {
			return exports, err
		}
		exports = append(exports, export)
	}

	return exports, nil
}

func GetExport(id uint64) (Exports, error) {
	sqlQuery := `SELECT * FROM exports WHERE export_id=$1;`

	var export Exports

	row := DB.QueryRow(sqlQuery, id)

	if err := row.Scan(&export.ExportID, &export.ExportName); err != nil {
		if err == sql.ErrNoRows {
      return export, fmt.Errorf("No export with id: %d", id)
		}
    return export, fmt.Errorf("Export %d: %v", id, err)
	}

	return export, nil

}

func CreateExport(export Exports) error {
	sqlQuery := `INSERT INTO exports(export_name) VALUES($1);`

	_, err := DB.Exec(sqlQuery, export.ExportName)
	if err != nil {
		return err
	}

	return nil
}

func UpdateExport(export Exports) error {
	sqlQuery := `UPDATE exports SET export_name=$1 WHERE export_id=$2;`

	_, err := DB.Exec(sqlQuery, export.ExportName, export.ExportID)
	if err != nil {
		return err
	}

	return nil

}

func DeleteExport(id uint64) error {
	sqlQuery := `DELETE FROM exports WHERE export_id=$1;`

	_, err := DB.Exec(sqlQuery, id)
	if err != nil {
		return err
	}

	return nil
}
