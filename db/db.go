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
	pool *pgxpool.Pool
}

var (
	pg   *Postgres
	once sync.Once

	dbLogger   *slog.Logger
	dbTimeZone string
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

	once.Do(func() {
		dbLogger = logger
		dbTimeZone = conf.TimeZone

		dbPool, dbErr := pgxpool.NewWithConfig(ctx, getParsedConfig(conf))
		if dbErr != nil {
			err = dbErr
		}

		pg = &Postgres{
			pool: dbPool,
		}
	})

	return pg, err
}

func (pg *Postgres) Conn() *pgxpool.Pool {
	return pg.pool
}

func (pg *Postgres) Ping(ctx context.Context) error {
	return pg.pool.Ping(ctx)
}

func (pg *Postgres) Close() {
	pg.pool.Close()
}

func setDbTimeZone(ctx context.Context, conn *pgx.Conn) error {
	// SET does not support parameterized queries
	query := fmt.Sprintf("SET TIME ZONE '%s'", dbTimeZone)

	if _, err := conn.Exec(ctx, query); err != nil {
		return fmt.Errorf("unable to set timezone: %w", err)
	}
	log.Printf("Timezone set to %s", dbTimeZone)
	return nil
}
