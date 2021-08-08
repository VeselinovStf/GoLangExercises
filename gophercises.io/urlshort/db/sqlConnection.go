package db

import "fmt"

// SQLConnection struct implements connection string parrameters
// for database
type SQLConnection struct {
	dataSource     string
	initialCatalog string
	user           string
	password       string
	port           int
}

// NewConnection creates sqlConnection
func NewConnection(dataSource, initialCatalog, user, password string) SQLConnection {
	return SQLConnection{
		dataSource:     dataSource,
		initialCatalog: initialCatalog,
		user:           user,
		password:       password,
	}
}

// ConnectionString returns full connection string
func (c *SQLConnection) ConnectionString() string {
	return fmt.Sprintf("Server=%s;Database=%s;User=%s;Password=%s;Port=1433", c.dataSource, c.initialCatalog, c.user, c.password)
}
