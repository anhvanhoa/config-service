package system_configuration

import (
	"config-service/domain/entity"
	"config-service/domain/repository"
	"context"

	"github.com/anhvanhoa/service-core/common"
)

// SystemConfigurationUsecase defines the main interface for system configuration business logic
type SystemConfigurationUsecase interface {
	// Individual usecase interfaces
	CreateSystemConfigurationUsecase
	GetSystemConfigurationUsecase
	UpdateSystemConfigurationUsecase
	DeleteSystemConfigurationUsecase
	ListSystemConfigurationUsecase

	// Additional business logic methods
	GetConfigValue(ctx context.Context, configKey string) (interface{}, error)
	SetConfigValue(ctx context.Context, configKey string, value interface{}) error
}

// SystemConfigurationUsecaseImpl implements SystemConfigurationUsecase
type SystemConfigurationUsecaseImpl struct {
	createUsecase CreateSystemConfigurationUsecase
	getUsecase    GetSystemConfigurationUsecase
	updateUsecase UpdateSystemConfigurationUsecase
	deleteUsecase DeleteSystemConfigurationUsecase
	listUsecase   ListSystemConfigurationUsecase
}

// NewSystemConfigurationUsecase creates a new system configuration usecase
func NewSystemConfigurationUsecase(repo repository.SystemConfigurationRepository) SystemConfigurationUsecase {
	return &SystemConfigurationUsecaseImpl{
		createUsecase: NewCreateSystemConfigurationUsecase(repo),
		getUsecase:    NewGetSystemConfigurationUsecase(repo),
		updateUsecase: NewUpdateSystemConfigurationUsecase(repo),
		deleteUsecase: NewDeleteSystemConfigurationUsecase(repo),
		listUsecase:   NewListSystemConfigurationUsecase(repo),
	}
}

// Create operations
func (u *SystemConfigurationUsecaseImpl) Execute(ctx context.Context, req CreateSystemConfigurationRequest) (*entity.SystemConfiguration, error) {
	return u.createUsecase.Execute(ctx, req)
}

// Get operations
func (u *SystemConfigurationUsecaseImpl) GetByID(ctx context.Context, id string) (*entity.SystemConfiguration, error) {
	return u.getUsecase.GetByID(ctx, id)
}

func (u *SystemConfigurationUsecaseImpl) GetByConfigKey(ctx context.Context, configKey string) (*entity.SystemConfiguration, error) {
	return u.getUsecase.GetByConfigKey(ctx, configKey)
}

// Update operations
func (u *SystemConfigurationUsecaseImpl) Update(ctx context.Context, req UpdateSystemConfigurationRequest) (*entity.SystemConfiguration, error) {
	return u.updateUsecase.Update(ctx, req)
}

func (u *SystemConfigurationUsecaseImpl) UpdateConfigValue(ctx context.Context, configKey string, value interface{}) (*entity.SystemConfiguration, error) {
	return u.updateUsecase.UpdateConfigValue(ctx, configKey, value)
}

// Delete operations
func (u *SystemConfigurationUsecaseImpl) Delete(ctx context.Context, id string) error {
	return u.deleteUsecase.Delete(ctx, id)
}

func (u *SystemConfigurationUsecaseImpl) DeleteByConfigKey(ctx context.Context, configKey string) error {
	return u.deleteUsecase.DeleteByConfigKey(ctx, configKey)
}

// List operations
func (u *SystemConfigurationUsecaseImpl) List(ctx context.Context, pagination common.Pagination, filter repository.SystemConfigurationFilter) ([]*entity.SystemConfiguration, int64, error) {
	return u.listUsecase.List(ctx, pagination, filter)
}

func (u *SystemConfigurationUsecaseImpl) GetConfigsByCategory(ctx context.Context, category string) ([]*entity.SystemConfiguration, error) {
	return u.listUsecase.GetConfigsByCategory(ctx, category)
}

func (u *SystemConfigurationUsecaseImpl) GetSystemConfigs(ctx context.Context) ([]*entity.SystemConfiguration, error) {
	return u.listUsecase.GetSystemConfigs(ctx)
}

func (u *SystemConfigurationUsecaseImpl) GetUserConfigs(ctx context.Context) ([]*entity.SystemConfiguration, error) {
	return u.listUsecase.GetUserConfigs(ctx)
}

// Additional business logic methods
func (u *SystemConfigurationUsecaseImpl) GetConfigValue(ctx context.Context, configKey string) (interface{}, error) {
	config, err := u.GetByConfigKey(ctx, configKey)
	if err != nil {
		return nil, err
	}

	return config.ConfigValue.Data, nil
}

func (u *SystemConfigurationUsecaseImpl) SetConfigValue(ctx context.Context, configKey string, value interface{}) error {
	_, err := u.UpdateConfigValue(ctx, configKey, value)
	return err
}
