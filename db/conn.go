package db

import (
	"database/sql"
	"fmt"

	"go-api/model"

	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

const (
	host     = "go_db"
	port     = 5432
	user     = "postgres"
	password = "1234"
	dbname   = "postgres"
)

func ConnectDB() (*gorm.DB, error) {

	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	fmt.Println("String Conexão: ", dsn)

	err := CreateBD(dsn)
	if err != nil {
		panic(err)
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // Desativa a pluralização
		},
	})

	if err != nil {
		return nil, err
	}

	// Migrar automaticamente a estrutura das tabelas
	err = db.AutoMigrate(&model.Product{})
	if err != nil {
		return nil, err
	}

	fmt.Println("Connected to", dbname)
	return db, nil
}

func CreateBD(strconn string) error {

	// Criar o banco de dados se ele não existir
	db, err := sql.Open("postgres", strconn)
	if err != nil {
		return err
	}
	defer db.Close()

	// Nome do banco de dados
	targetDBName := "go_api"

	exists, err := checkDatabaseExists(db, targetDBName)
	if err != nil {
		panic(err)
	}

	if !exists {
		// Query para criar o banco de dados
		createDBQuery := fmt.Sprintf("CREATE DATABASE %s", targetDBName)

		_, err = db.Exec(createDBQuery)
		if err != nil {
			return err
		}

		db.Close()

		fmt.Println("Database created successfully")
	}

	return nil

}

func checkDatabaseExists(db *sql.DB, dbName string) (bool, error) {
	var exists bool

	query := `SELECT EXISTS(SELECT 1 FROM pg_database WHERE datname = $1)`
	err := db.QueryRow(query, dbName).Scan(&exists)
	if err != nil {
		return false, err
	}

	return exists, nil
}
