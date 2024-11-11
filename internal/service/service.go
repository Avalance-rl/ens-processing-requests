package service

import "context"

// TODO: переделать

type ProcessReqService interface {
	CreateGroup(ctx context.Context) (string, error)
	DeleteGroup(ctx context.Context, id string) error
	ListGroups(ctx context.Context, id string) ([]string, error)
	ExecuteDispatch(ctx context.Context, id string) error
}
