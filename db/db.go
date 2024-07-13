package db

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"sync"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/pgx/v5/tracelog"

	"github.com/AxelTahmid/golang-starter/config"
)

type Postgres struct {
	DB *pgxpool.Pool
}

var (
	pool     *Postgres
	dbLogger *slog.Logger
	pgOnce   sync.Once
)

func getParsedConfig(conf config.Database) *pgxpool.Config {

	dbConfig, err := pgxpool.ParseConfig(conf.Url)

	if err != nil {
		log.Fatalf("Unable to parse database URL: %v", err)
	}

	dbConfig.MaxConns = conf.PoolMax
	dbConfig.MinConns = conf.PoolMin
	dbConfig.MaxConnLifetime = conf.MaxConnLifetime
	dbConfig.MaxConnIdleTime = conf.MaxConnIdleTime
	dbConfig.HealthCheckPeriod = conf.HealthCheckPeriod
	dbConfig.ConnConfig.ConnectTimeout = conf.ConnectTimeout

	dbConfig.AfterConnect = setDbTimeZone

	dbConfig.ConnConfig.Tracer = &tracelog.TraceLog{
		Logger:   InitLogger(dbLogger),
		LogLevel: tracelog.LogLevelTrace,
	}

	return dbConfig
}

func CreatePool(ctx context.Context, conf config.Database, logger *slog.Logger) (*Postgres, error) {
	var err error

	pgOnce.Do(func() {
		dbLogger = logger

		dbPool, dbErr := pgxpool.NewWithConfig(ctx, getParsedConfig(conf))
		if dbErr != nil {
			err = dbErr
		}

		pool = &Postgres{dbPool}
	})

	return pool, err
}

func (pg *Postgres) Ping(ctx context.Context) error {
	return pg.DB.Ping(ctx)
}

func (pg *Postgres) Close() {
	pg.DB.Close()
}

func setDbTimeZone(ctx context.Context, conn *pgx.Conn) error {
	_, err := conn.Exec(ctx, "SET TIME ZONE 'UTC';")

	if err != nil {
		return fmt.Errorf("unable to set timezone: %w", err)
	}
	log.Printf("Timezone set to UTC\n")
	return nil
}
