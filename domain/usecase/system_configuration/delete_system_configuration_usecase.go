package system_configuration

import (
	"config-service/domain/repository"
	"context"
)

type DeleteSystemConfigurationUsecase interface {
	DeleteByConfigKey(ctx context.Context, configKey string) error
}

type DeleteSystemConfigurationUsecaseImpl struct {
	repo repository.SystemConfigurationRepository
}

func NewDeleteSystemConfigurationUsecase(repo repository.SystemConfigurationRepository) DeleteSystemConfigurationUsecase {
	return &DeleteSystemConfigurationUsecaseImpl{
		repo: repo,
	}
}

func (u *DeleteSystemConfigurationUsecaseImpl) DeleteByConfigKey(ctx context.Context, configKey string) error {
	if configKey == "" {
		return ErrInvalidConfigKey
	}

	config, err := u.repo.GetByConfigKey(ctx, configKey)
	if err != nil {
		return ErrConfigNotFound
	}

	if config.IsSystemConfig {
		return ErrCannotDeleteSystemConfig
	}

	return u.repo.Delete(ctx, config.ID)
}
