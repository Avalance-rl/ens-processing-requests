package processing_requests

import (
	"context"

	desc "github.com/Avalance-rl/processing/pkg/processing_v1"
)

func (s *ServerAPI) DeleteGroup(
	ctx context.Context,
	req *desc.DeleteGroupRequest,
) (*desc.DeleteGroupResponse, error) {
	panic("implement me")
}
