package main

import (
	"fmt"
	"github.com/Rosya-edwica/clickhouse_experience/internal/dbs/clickhouse"
	"github.com/Rosya-edwica/clickhouse_experience/internal/dbs/mysql"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

func main() {
	clickhouseCfg := initClickHouseConfig()
	mysqlCfg := initMySQLConfig()

	mysqlConn, err := mysql.New(mysqlCfg)
	if err != nil {
		log.Fatal(err)
	}
	defer mysqlConn.Close()
	clickhouseConn, err := clickhouse.New(clickhouseCfg)
	if err != nil {
		log.Fatal(err)
	}
	defer clickhouseConn.Close()

	var lastId int
	for {
		vacancies, err := mysql.GetVacancies(mysqlConn, lastId)
		if err != nil {
			log.Fatal(err)
		}
		if len(vacancies) == 0 {
			break
		}
		err = clickhouse.SaveVacancies(clickhouseConn, vacancies)
		if err != nil {
			log.Fatal(err)
		}
		lastId = vacancies[len(vacancies)-1].Id
		fmt.Printf("Вставили:%d вакансий (Последний id: %d)\n", len(vacancies), lastId)
	}
}

func initClickHouseConfig() clickhouse.Config {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}
	portInt, err := strconv.Atoi(os.Getenv("PORT_CLICKHOUSE"))
	if err != nil {
		log.Fatal(err)
	}
	cfg := clickhouse.Config{
		Addr:     os.Getenv("HOST_CLICKHOUSE"),
		Port:     uint16(portInt),
		User:     os.Getenv("USER_CLICKHOUSE"),
		Password: os.Getenv("PASSWORD_CLICKHOUSE"),
		DB:       os.Getenv("DB_CLICKHOUSE"),
	}
	return cfg
}

func initMySQLConfig() mysql.Config {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}
	portInt, err := strconv.Atoi(os.Getenv("PORT_MYSQL"))
	if err != nil {
		log.Fatal(err)
	}
	cfg := mysql.Config{
		Addr:     os.Getenv("HOST_MYSQL"),
		Port:     uint16(portInt),
		User:     os.Getenv("USER_MYSQL"),
		Password: os.Getenv("PASSWORD_MYSQL"),
		DB:       os.Getenv("DB_MYSQL"),
	}
	return cfg
}
