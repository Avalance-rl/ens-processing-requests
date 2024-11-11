package processing_requests

// descriptor alias
import (
	"ens-processing-requests/internal/service"

	desc "github.com/Avalance-rl/processing/pkg/processing_v1"
)

type ServerAPI struct {
	desc.UnimplementedProcessingServiceServer
	processReqService service.ProcessReqService
}

func NewServerAPI(processReqService service.ProcessReqService) *ServerAPI {
	return &ServerAPI{
		processReqService: processReqService,
	}
}
