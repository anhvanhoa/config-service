package system_configuration

import (
	"config-service/domain/repository"
	"context"
)

// DeleteSystemConfigurationUsecase defines the interface for deleting system configurations
type DeleteSystemConfigurationUsecase interface {
	Delete(ctx context.Context, id string) error
	DeleteByConfigKey(ctx context.Context, configKey string) error
}

// DeleteSystemConfigurationUsecaseImpl implements DeleteSystemConfigurationUsecase
type DeleteSystemConfigurationUsecaseImpl struct {
	repo repository.SystemConfigurationRepository
}

// NewDeleteSystemConfigurationUsecase creates a new delete system configuration usecase
func NewDeleteSystemConfigurationUsecase(repo repository.SystemConfigurationRepository) DeleteSystemConfigurationUsecase {
	return &DeleteSystemConfigurationUsecaseImpl{
		repo: repo,
	}
}

// Delete deletes a system configuration by ID
func (u *DeleteSystemConfigurationUsecaseImpl) Delete(ctx context.Context, id string) error {
	if id == "" {
		return ErrInvalidID
	}

	// Check if config exists
	config, err := u.repo.GetByID(ctx, id)
	if err != nil {
		return err
	}

	if config == nil {
		return ErrConfigNotFound
	}

	// Check if it's a system config (cannot delete)
	if config.IsSystemConfig {
		return ErrCannotDeleteSystemConfig
	}

	// Delete from repository
	return u.repo.Delete(ctx, id)
}

// DeleteByConfigKey deletes a system configuration by config key
func (u *DeleteSystemConfigurationUsecaseImpl) DeleteByConfigKey(ctx context.Context, configKey string) error {
	if configKey == "" {
		return ErrInvalidConfigKey
	}

	// Get config by key first
	config, err := u.repo.GetByConfigKey(ctx, configKey)
	if err != nil {
		return err
	}

	if config == nil {
		return ErrConfigNotFound
	}

	// Check if it's a system config (cannot delete)
	if config.IsSystemConfig {
		return ErrCannotDeleteSystemConfig
	}

	// Delete from repository
	return u.repo.Delete(ctx, config.ID)
}
