package logdb

import (
	"context"
	"fmt"
	"log"
	"test-task/domain"
	"time"

	"github.com/ClickHouse/clickhouse-go/v2"
	driver "github.com/ClickHouse/clickhouse-go/v2/lib/driver"
)

type LogDB struct {
	context context.Context
	conn    driver.Conn
}

func NewLogDB(ctx context.Context) *LogDB {
	conn, err := clickhouse.Open(&clickhouse.Options{
		Addr: []string{"127.0.0.1:9000"},
		Auth: clickhouse.Auth{
			Database: "default",
			Username: "default",
			Password: "default",
		},
		Compression: &clickhouse.Compression{
			Method: clickhouse.CompressionLZ4,
		},
		Settings: clickhouse.Settings{
			"max_execution_time": 60,
		},
		Debug: true,
	})
	if err != nil {
		log.Fatalf("%v", err)
	}
	conn.Ping(ctx)

	err = conn.Exec(ctx,
		`CREATE TABLE IF NOT EXISTS users (
				name String,
				timestamp DateTime
			) ENGINE = Kafka('localhost:9092', 'users', 'group1', 'JSONEachRow');
		`)
	if err != nil {
		log.Fatalf("%v", err)
	}

	return &LogDB{context: ctx, conn: conn}
}

func (logdb *LogDB) LogUser(user domain.User) {
	if err := logdb.conn.Exec(logdb.context, fmt.Sprintf("INSERT INTO users VALUES ('%s', %d);", user.Name, time.Now().Unix())); err != nil {
		log.Fatal(err)
	}
}
