package app

import (
	procReqApi "ens-processing-requests/internal/api/processing-requests"
	"ens-processing-requests/internal/config"
	"ens-processing-requests/internal/service"
	procReqService "ens-processing-requests/internal/service/processing-requests"
	"ens-processing-requests/internal/storage"
	procReqStorage "ens-processing-requests/internal/storage/processing-requests"
	"log"
)

type serviceProvider struct {
	grpcConfig config.GRPCConfig

	groupCreator   storage.GroupCreator
	groupDeleter   storage.GroupDeleter
	groupProvider  storage.GroupProvider
	groupsProvider storage.GroupsProvider

	processReqService service.ProcessReqService

	serverAPI *procReqApi.ServerAPI
}

func newServiceProvider() *serviceProvider {
	return &serviceProvider{}
}

func (s *serviceProvider) GRPCConfig() config.GRPCConfig {
	if s.grpcConfig == nil {
		cfg, err := config.NewGRPCConfig()
		if err != nil {
			log.Fatalf("failed to get grpc config: %s", err)
		}
		s.grpcConfig = cfg
	}

	return s.grpcConfig
}

func (s *serviceProvider) GroupCreator() storage.GroupCreator {
	if s.groupCreator == nil {
		s.groupCreator = procReqStorage.NewPostgreStorage()
	}
	return s.groupCreator
}

func (s *serviceProvider) GroupDeleter() storage.GroupDeleter {
	if s.groupDeleter == nil {
		s.groupDeleter = procReqStorage.NewPostgreStorage()
	}
	return s.groupDeleter
}

func (s *serviceProvider) GroupsProvider() storage.GroupsProvider {
	if s.groupsProvider == nil {
		s.groupsProvider = procReqStorage.NewPostgreStorage()
	}
	return s.groupsProvider
}

func (s *serviceProvider) GroupProvider() storage.GroupProvider {
	if s.groupProvider == nil {
		s.groupProvider = procReqStorage.NewPostgreStorage()
	}
	return s.groupProvider
}

func (s *serviceProvider) ProcessReqService() service.ProcessReqService {
	if s.processReqService == nil {
		s.processReqService = procReqService.NewService(
			s.GroupCreator(),
			s.GroupDeleter(),
			s.GroupsProvider(),
			s.GroupProvider(),
		)
	}
	return s.processReqService
}

func (s *serviceProvider) ServerAPI() *procReqApi.ServerAPI {
	if s.serverAPI == nil {
		s.serverAPI = procReqApi.NewServerAPI(s.ProcessReqService())
	}
	return s.serverAPI
}
