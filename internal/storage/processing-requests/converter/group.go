package converter

import (
	"ens-processing-requests/internal/model"
	storageModel "ens-processing-requests/internal/storage/processing-requests/model"
)

func groupFromStorage(group *storageModel.Group) *model.Group {
	return &model.Group{
		ID:        group.ID,
		Name:      group.Name,
		UserData:  group.UserData,
		CreatedAt: group.CreatedAt.Time,
		UserEmail: group.UserEmail,
	}
}
