package mysqldbmodels

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DBClient struct {
	Conn *gorm.DB
}

func InitializeDatabase() (*DBClient, error) {

	// dbUser := os.Getenv("DB_USER")
	// dbPassword := os.Getenv("DB_PASSWORD")
	// dbName := os.Getenv("DB_NAME")
	// dbHost := os.Getenv("DB_HOST")
	// dbPort := os.Getenv("DB_PORT")

	// if dbUser == "" || dbPassword == "" || dbName == "" || dbHost == "" || dbPort == "" {
	// 	return nil, errors.New("One or more required environment variables are not set")
	// }

	//temp

	// MySQL connection parameters
	dbUser := "root"
	dbPassword := "1234"
	dbName := "coursemanagement"
	dbHost := "localhost"
	dbPort := "3306" // Default MySQL port

	// Construct DSN (Data Source Name)
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUser, dbPassword, dbHost, dbPort, dbName)

	// Connect to MySQL database
	DB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	client := &DBClient{Conn: DB}
	return client, nil
}
