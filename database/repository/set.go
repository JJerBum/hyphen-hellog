package repository

import (
	"context"
	"hyphen-hellog/cerrors/exception"
	"hyphen-hellog/ent"
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"

	"entgo.io/ent/dialect/sql"

	_ "github.com/go-sql-driver/mysql"
)

type DBType ent.Client

var instance *DBType = nil

func load() *DBType {
	username := os.Getenv("DATASOURCE_USERNAME")
	password := os.Getenv("DATASOURCE_PASSWORD")
	host := os.Getenv("DATASOURCE_HOST")
	port := os.Getenv("DATASOURCE_PORT")
	dbName := os.Getenv("DATASOURCE_DB_NAME")
	maxPoolIdle, err := strconv.Atoi(os.Getenv("DATASOURCE_POOL_IDLE_CONN"))
	maxPoolOpen, err := strconv.Atoi(os.Getenv("DATASOURCE_POOL_MAX_CONN"))
	maxPollLifeTime, err := strconv.Atoi(os.Getenv("DATASOURCE_POOL_LIFE_TIME"))
	exception.Sniff(err)

	dsn := username + ":" + password + "@tcp(" + host + ":" + port + ")/" + dbName + "?parseTime=true"
	drv, err := sql.Open("mysql", dsn)
	exception.Sniff(err)

	// Get the underlying sql.DB object of the driver.
	db := drv.DB()
	db.SetMaxIdleConns(maxPoolIdle)
	db.SetMaxOpenConns(maxPoolOpen)
	db.SetConnMaxLifetime(time.Duration(rand.Int31n(int32(maxPollLifeTime))) * time.Millisecond)
	client := ent.NewClient(ent.Driver(drv))

	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	var temp = DBType(*client)
	return &temp
}

func Get() *DBType {
	if instance == nil {
		instance = load()
	}

	return instance
}
