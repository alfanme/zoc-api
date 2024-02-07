package config

import (
	"fmt"
	"os"
)

func GetDSN() string {
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	pwd := os.Getenv("DB_PWD")
	dbname := os.Getenv("DB_NAME")
	dbport := os.Getenv("DB_PORT")
	timezone := os.Getenv("DB_TIMEZONE")

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=prefer TimeZone=%s",
		host,
		user,
		pwd,
		dbname,
		dbport,
		timezone,
	)

	return dsn
}