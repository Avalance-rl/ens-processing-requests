package processing_requests

import (
	"context"
	serviceModel "ens-processing-requests/internal/model"

	"github.com/jackc/pgx/v5/pgxpool"
)

func (p *Postgre) CreateGroup(ctx context.Context, group serviceModel.Group) (serviceModel.Group, error) {
	if err := p.Go(func(db *pgxpool.Pool) error {

		return nil
	}); err != nil {
		return serviceModel.Group{}, err
	}

	return group, nil

}

func (p *Postgre) DeleteGroup() {

}

func (p *Postgre) ListGroups() {

}

func (p *Postgre) Group() {

}
