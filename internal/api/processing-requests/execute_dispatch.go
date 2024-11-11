package processing_requests

import (
	"context"

	desc "github.com/Avalance-rl/processing/pkg/processing_v1"
)

func (s *ServerAPI) ExecuteDispatch(
	ctx context.Context,
	req *desc.ExecuteDispatchRequest,
) (*desc.ExecuteDispatchResponse, error) {
	panic("implement me")
}
