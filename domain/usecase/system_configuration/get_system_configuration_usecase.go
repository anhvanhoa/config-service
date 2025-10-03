package system_configuration

import (
	"config-service/domain/entity"
	"config-service/domain/repository"
	"context"
)

// GetSystemConfigurationUsecase defines the interface for getting system configurations
type GetSystemConfigurationUsecase interface {
	GetByID(ctx context.Context, id string) (*entity.SystemConfiguration, error)
	GetByConfigKey(ctx context.Context, configKey string) (*entity.SystemConfiguration, error)
}

// GetSystemConfigurationUsecaseImpl implements GetSystemConfigurationUsecase
type GetSystemConfigurationUsecaseImpl struct {
	repo repository.SystemConfigurationRepository
}

// NewGetSystemConfigurationUsecase creates a new get system configuration usecase
func NewGetSystemConfigurationUsecase(repo repository.SystemConfigurationRepository) GetSystemConfigurationUsecase {
	return &GetSystemConfigurationUsecaseImpl{
		repo: repo,
	}
}

// GetByID gets a system configuration by ID
func (u *GetSystemConfigurationUsecaseImpl) GetByID(ctx context.Context, id string) (*entity.SystemConfiguration, error) {
	if id == "" {
		return nil, ErrInvalidID
	}

	config, err := u.repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if config == nil {
		return nil, ErrConfigNotFound
	}

	return config, nil
}

// GetByConfigKey gets a system configuration by config key
func (u *GetSystemConfigurationUsecaseImpl) GetByConfigKey(ctx context.Context, configKey string) (*entity.SystemConfiguration, error) {
	if configKey == "" {
		return nil, ErrInvalidConfigKey
	}

	config, err := u.repo.GetByConfigKey(ctx, configKey)
	if err != nil {
		return nil, err
	}

	if config == nil {
		return nil, ErrConfigNotFound
	}

	return config, nil
}
