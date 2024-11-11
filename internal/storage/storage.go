package storage

type GroupCreator interface {
	CreateGroup()
}

type GroupDeleter interface {
	DeleteGroup()
}

type GroupsProvider interface {
	ListGroups()
}

type GroupProvider interface {
	Group()
}
