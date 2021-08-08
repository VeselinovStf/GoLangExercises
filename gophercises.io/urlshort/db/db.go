// Package db implements database communication functionalities
package db

import (
	"database/sql"

	_ "github.com/denisenkom/go-mssqldb"
)


const (
	selectAllPathsQuery = "Select Path,URL FROM Path"
	driverName = "mssql"
)


// GetURLs return db saved url-pares
func GetURLs(connection SQLConnection) (map[string]string, error) {
	con, err := sql.Open(driverName, connection.ConnectionString())
	if err != nil {
		return nil, err
	}
	defer con.Close()

	rows, err := con.Query(selectAllPathsQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	result := make(map[string]string)
	for rows.Next() {
		var (
			path string
			url  string
		)
		if err := rows.Scan(&path, &url); err != nil {
			return nil, err
		}
		result[path] = url
	}

	return result, nil
}
