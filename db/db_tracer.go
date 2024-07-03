package db

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type pgTracer struct {
	log *zerolog.Logger
}

// zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
// log := zerolog.New(os.Stderr).With().Timestamp().Logger()

func (tracer *pgTracer) TraceQueryStart(ctx context.Context, _ *pgx.Conn, data pgx.TraceQueryStartData) context.Context {

	start := time.Now()

	defer func() {
		log.Info().
			Str("sql statement", data.SQL).
			Str("sql args", fmt.Sprintf("%v", data.Args)).
			Dur("latency", time.Since(start)).
			Msg("executing command")
	}()

	return ctx
}

func (tracer *pgTracer) TraceQueryEnd(ctx context.Context, conn *pgx.Conn, data pgx.TraceQueryEndData) {
}
