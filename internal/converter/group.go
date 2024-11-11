package converter

import (
	"ens-processing-requests/internal/model"

	desc "github.com/Avalance-rl/processing/pkg/processing_v1"
)

func ToGroupFromService(group *model.Group) *desc.Group {
	return &desc.Group{
		Id:        group.ID,
		Name:      group.Name,
		UserData:  group.UserData,
		CreatedAt: group.CreatedAt.String(),
		UserEmail: group.UserEmail,
	}
}
