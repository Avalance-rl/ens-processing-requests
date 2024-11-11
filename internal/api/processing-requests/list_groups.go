package processing_requests

import (
	"context"

	desc "github.com/Avalance-rl/processing/pkg/processing_v1"
)

func (s *ServerAPI) ListGroups(
	ctx context.Context,
	req *desc.ListGroupsRequest,
) (*desc.ListGroupsResponse, error) {
	panic("implement me")
}
