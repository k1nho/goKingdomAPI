package model

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

	rows, err := DB.Query(sqlQuery, id)
	if err != nil {
		return export, err
	}

	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&export.ExportID, &export.ExportName)
		if err != nil {
			return export, err
		}
	}

	return export, nil

}

func CreateExport(export Exports) error {
	sqlQuery = `INSERT INTO exports(export_name) VALUES($1);`

}

func UpdateExport() error {

}

func DeleteExport() error {

}
