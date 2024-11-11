package processing_requests

import (
	"context"
	"ens-processing-requests/internal/config"
	"ens-processing-requests/utils"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Postgre struct {
	db  *pgxpool.Pool
	sem *utils.Semaphore
}

func NewPostgreStorage(ctx context.Context, cfg config.StorageConfig, maxPool int) (*Postgre, error) {
	poolConfig, err := pgxpool.ParseConfig(cfg.DNS())
	if err != nil {
		return nil, err
	}
	poolConfig.ConnConfig.ConnectTimeout = 5 * time.Second
	poolConfig.MaxConnLifetime = 1 * time.Hour
	poolConfig.MaxConnIdleTime = 30 * time.Minute
	poolConfig.HealthCheckPeriod = 1 * time.Minute

	pool, err := pgxpool.NewWithConfig(ctx, poolConfig)
	if err != nil {
		return nil, err
	}
	sem := utils.NewSemaphore(maxPool)
	return &Postgre{
		db:  pool,
		sem: sem,
	}, nil
}

func (p *Postgre) Go(f func(db *pgxpool.Pool) error) error {
	resultChan := make(chan error)
	go func() {
		p.sem.Acquire()
		defer p.sem.Release()

		result := f(p.db)
		resultChan <- result
	}()

	return <-resultChan
}

func (p *Postgre) Close(ctx context.Context) error {
	done := make(chan struct{}, 1)
	go func() {
		p.db.Close()
		done <- struct{}{}
	}()

	select {
	case <-ctx.Done():
		return ctx.Err()
	case <-done:
		return nil
	}

}
