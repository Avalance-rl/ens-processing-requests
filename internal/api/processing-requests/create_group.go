package processing_requests

import (
	"context"

	desc "github.com/Avalance-rl/processing/pkg/processing_v1"
)

func (s *ServerAPI) CreateGroup(
	ctx context.Context,
	req *desc.CreateGroupRequest,
) (*desc.CreateGroupResponse, error) {

	return &desc.CreateGroupResponse{
		GroupID:     req.GroupName,
		InfoMessage: req.UserEmail,
	}, nil
}
