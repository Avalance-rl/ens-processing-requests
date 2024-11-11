package config

import (
	"errors"
	"fmt"
	"os"
)

const (
	storageUser     = "STORAGE_USER"
	storagePassword = "STORAGE_PASSWORD"
	storageHost     = "STORAGE_HOST"
	storagePort     = "STORAGE_PORT"
	storageName     = "STORAGE_NAME"
	storagePoolSize = "STORAGE_POLL_SIZE"
)

type StorageConfig interface {
	DNS() string
}

type storageConfig struct {
	user     string
	password string
	host     string
	port     string
	name     string
	poolSize string
}

func NewStorageConfig() (StorageConfig, error) {
	user := os.Getenv(storageUser)
	if user == "" {
		return nil, errors.New("storage user not found")
	}

	password := os.Getenv(storagePassword)
	if password == "" {
		return nil, errors.New("storage password not found")
	}

	host := os.Getenv(storageHost)
	if host == "" {
		return nil, errors.New("storage host not found")
	}

	port := os.Getenv(storagePort)
	if port == "" {

		return nil, errors.New("storage port not found")
	}

	name := os.Getenv(storageName)
	if name == "" {
		return nil, errors.New("storage name not found")
	}

	pollSize := os.Getenv(storagePoolSize)
	if pollSize == "" {
		return nil, errors.New("storage pool size not found")
	}

	return &storageConfig{
		user:     user,
		password: password,
		host:     host,
		port:     port,
		name:     name,
		poolSize: pollSize,
	}, nil
}

func (cfg *storageConfig) DNS() string {
	return fmt.Sprintf(
		"postgresql://%s:%s@%s:%d/%s?pool_max_conns=%d",
		cfg.user, cfg.password, cfg.host, cfg.port, cfg.name, cfg.poolSize,
	)
}
