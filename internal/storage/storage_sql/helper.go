package storage_sql

import "fmt"

func createTableStatement(tableName string) string {
	return fmt.Sprintf(`CREATE TABLE %s (
    id      INTEGER PRIMARY KEY AUTOINCREMENT,
    entity 	TEXT NOT NULL
);`, tableName)
}
