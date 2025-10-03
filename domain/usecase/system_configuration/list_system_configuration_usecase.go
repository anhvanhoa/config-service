package system_configuration

import (
	"config-service/domain/entity"
	"config-service/domain/repository"
	"context"

	"github.com/anhvanhoa/service-core/common"
)

// ListSystemConfigurationUsecase defines the interface for listing system configurations
type ListSystemConfigurationUsecase interface {
	List(ctx context.Context, pagination common.Pagination, filter repository.SystemConfigurationFilter) ([]*entity.SystemConfiguration, int64, error)
	GetConfigsByCategory(ctx context.Context, category string) ([]*entity.SystemConfiguration, error)
	GetSystemConfigs(ctx context.Context) ([]*entity.SystemConfiguration, error)
	GetUserConfigs(ctx context.Context) ([]*entity.SystemConfiguration, error)
}

// ListSystemConfigurationUsecaseImpl implements ListSystemConfigurationUsecase
type ListSystemConfigurationUsecaseImpl struct {
	repo repository.SystemConfigurationRepository
}

// NewListSystemConfigurationUsecase creates a new list system configuration usecase
func NewListSystemConfigurationUsecase(repo repository.SystemConfigurationRepository) ListSystemConfigurationUsecase {
	return &ListSystemConfigurationUsecaseImpl{
		repo: repo,
	}
}

// List gets a list of system configurations with filtering and pagination
func (u *ListSystemConfigurationUsecaseImpl) List(ctx context.Context, pagination common.Pagination, filter repository.SystemConfigurationFilter) ([]*entity.SystemConfiguration, int64, error) {
	configs, total, err := u.repo.List(ctx, pagination, filter)
	if err != nil {
		return nil, 0, err
	}

	return configs, total, nil
}

// GetConfigsByCategory gets configurations by category
func (u *ListSystemConfigurationUsecaseImpl) GetConfigsByCategory(ctx context.Context, category string) ([]*entity.SystemConfiguration, error) {
	if category == "" {
		return nil, ErrInvalidCategory
	}

	filter := repository.SystemConfigurationFilter{
		Category: category,
	}

	configs, _, err := u.repo.List(ctx, common.Pagination{}, filter)
	return configs, err
}

// GetSystemConfigs gets all system configurations
func (u *ListSystemConfigurationUsecaseImpl) GetSystemConfigs(ctx context.Context) ([]*entity.SystemConfiguration, error) {
	filter := repository.SystemConfigurationFilter{
		IsSystemConfig: true,
	}

	configs, _, err := u.repo.List(ctx, common.Pagination{}, filter)
	return configs, err
}

// GetUserConfigs gets all user configurations
func (u *ListSystemConfigurationUsecaseImpl) GetUserConfigs(ctx context.Context) ([]*entity.SystemConfiguration, error) {
	filter := repository.SystemConfigurationFilter{
		IsSystemConfig: false,
	}

	configs, _, err := u.repo.List(ctx, common.Pagination{}, filter)
	return configs, err
}
