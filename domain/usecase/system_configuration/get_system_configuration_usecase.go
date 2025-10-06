package system_configuration

import (
	"config-service/domain/entity"
	"config-service/domain/repository"
	"context"
)

type GetSystemConfigurationUsecase interface {
	GetByConfigKey(ctx context.Context, configKey string) (*entity.SystemConfiguration, error)
}

type GetSystemConfigurationUsecaseImpl struct {
	repo repository.SystemConfigurationRepository
}

func NewGetSystemConfigurationUsecase(repo repository.SystemConfigurationRepository) GetSystemConfigurationUsecase {
	return &GetSystemConfigurationUsecaseImpl{
		repo: repo,
	}
}

func (u *GetSystemConfigurationUsecaseImpl) GetByConfigKey(ctx context.Context, configKey string) (*entity.SystemConfiguration, error) {
	if configKey == "" {
		return nil, ErrInvalidConfigKey
	}

	config, err := u.repo.GetByConfigKey(ctx, configKey)
	if err != nil {
		return nil, ErrConfigNotFound
	}

	return config, nil
}
