package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

var (
	host     = "localhost"
	port     = "5432"
	user     = "postgres"
	password = "postgres"
	dbname   = "h8-product"
	dialect  = "postgres"
)
var (
	db  *sql.DB
	err error
)

func handleDatabaseConnection() {
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname,
	)

	db, err = sql.Open(dialect, psqlInfo)

	if err != nil {
		log.Panic("error occured while trying to validate database arguments:", err)
	}

	err = db.Ping()

	if err != nil {
		log.Panic("error occured while trying to connect to database:", err)
	}
}
func handleCreateRequiredTables() {
	userTable := `
		CREATE TABLE IF NOT EXISTS "users" (
			id SERIAL PRIMARY KEY,
			email VARCHAR(255) UNIQUE NOT NULL,
			password TEXT NOT NULL,
			level varchar(100) NOT NULL DEFAULT "customer",
			createdAt timestamptz DEFAULT now(),
			updatedAt timestamptz DEFAULT now()
		);
	`

	productTable := `
		CREATE TABLE IF NOT EXISTS "products" (
			id SERIAL PRIMARY KEY,
			title VARCHAR(255) NOT NULL,
			description TEXT NOT NULL,
			price int NOT NULL,
			userId int NOT NULL,
			createdAt timestamptz DEFAULT now(),
			updatedAt timestamptz DEFAULT now(),
			CONSTRAINT movies_user_id_fk
				FOREIGN KEY(userId)
					REFERENCES users(id)
						ON DELETE CASCADE
		);
	`

	createTableQueries := fmt.Sprintf("%s %s", userTable, productTable)

	_, err = db.Exec(createTableQueries)

	if err != nil {
		log.Panic("error occured while trying to create required tables:", err)
	}
}
func InitiliazeDatabase() {
	handleDatabaseConnection()
	handleCreateRequiredTables()
}
func GetDatabaseInstance() *sql.DB {
	return db
}
