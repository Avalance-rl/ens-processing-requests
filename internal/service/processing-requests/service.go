package processing_requests

import (
	def "ens-processing-requests/internal/service"
	"ens-processing-requests/internal/storage"
)

var _ def.ProcessReqService = (*service)(nil)

type service struct {
	groupCreator   storage.GroupCreator
	groupDeleter   storage.GroupDeleter
	groupsProvider storage.GroupsProvider
	groupProvider  storage.GroupProvider
}

func NewService(
	groupCreator storage.GroupCreator,
	groupDeleter storage.GroupDeleter,
	groupsProvider storage.GroupsProvider,
	groupProvider storage.GroupProvider,
) *service {
	return &service{
		groupCreator:   groupCreator,
		groupDeleter:   groupDeleter,
		groupsProvider: groupsProvider,
		groupProvider:  groupProvider,
	}
}
