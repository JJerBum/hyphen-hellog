package initializer

import (
	"context"
	"hyphen-hellog/cerrors/exception"
	"hyphen-hellog/ent"
	"log"
	"math/rand"
	"strconv"
	"time"

	"entgo.io/ent/dialect/sql"

	_ "github.com/go-sql-driver/mysql"
)

func NewDatabase(c config) *ent.Client {
	username := c.Get("DATASOURCE_USERNAME")
	password := c.Get("DATASOURCE_PASSWORD")
	host := c.Get("DATASOURCE_HOST")
	port := c.Get("DATASOURCE_PORT")
	dbName := c.Get("DATASOURCE_DB_NAME")
	maxPoolIdle, err := strconv.Atoi(c.Get("DATASOURCE_POOL_IDLE_CONN"))
	maxPoolOpen, err := strconv.Atoi(c.Get("DATASOURCE_POOL_MAX_CONN"))
	maxPollLifeTime, err := strconv.Atoi(c.Get("DATASOURCE_POOL_LIFE_TIME"))
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

	return client
}
