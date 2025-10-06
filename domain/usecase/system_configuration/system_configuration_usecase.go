package system_configuration

import (
	"config-service/domain/entity"
	"config-service/domain/repository"
	"context"

	"github.com/anhvanhoa/service-core/common"
	"github.com/anhvanhoa/service-core/utils"
)

type SystemConfigurationUsecase interface {
	Create(ctx context.Context, req CreateSystemConfigurationRequest) (*entity.SystemConfiguration, error)
	GetByConfigKey(ctx context.Context, configKey string) (*entity.SystemConfiguration, error)
	Update(ctx context.Context, req UpdateSystemConfigurationRequest) (*entity.SystemConfiguration, error)
	DeleteByConfigKey(ctx context.Context, configKey string) error
	List(ctx context.Context, pagination common.Pagination, filter repository.SystemConfigurationFilter) (common.PaginationResult[*entity.SystemConfiguration], int64, error)
	GetConfigValue(ctx context.Context, configKey string) (interface{}, error)
}

type SystemConfigurationUsecaseImpl struct {
	createUsecase CreateSystemConfigurationUsecase
	getUsecase    GetSystemConfigurationUsecase
	updateUsecase UpdateSystemConfigurationUsecase
	deleteUsecase DeleteSystemConfigurationUsecase
	listUsecase   ListSystemConfigurationUsecase
}

func NewSystemConfigurationUsecase(repo repository.SystemConfigurationRepository, helper utils.Helper) SystemConfigurationUsecase {
	return &SystemConfigurationUsecaseImpl{
		createUsecase: NewCreateSystemConfigurationUsecase(repo),
		getUsecase:    NewGetSystemConfigurationUsecase(repo),
		updateUsecase: NewUpdateSystemConfigurationUsecase(repo),
		deleteUsecase: NewDeleteSystemConfigurationUsecase(repo),
		listUsecase:   NewListSystemConfigurationUsecase(repo, helper),
	}
}

func (u *SystemConfigurationUsecaseImpl) Create(ctx context.Context, req CreateSystemConfigurationRequest) (*entity.SystemConfiguration, error) {
	return u.createUsecase.Execute(ctx, req)
}

func (u *SystemConfigurationUsecaseImpl) GetByConfigKey(ctx context.Context, configKey string) (*entity.SystemConfiguration, error) {
	return u.getUsecase.GetByConfigKey(ctx, configKey)
}

func (u *SystemConfigurationUsecaseImpl) Update(ctx context.Context, req UpdateSystemConfigurationRequest) (*entity.SystemConfiguration, error) {
	return u.updateUsecase.Execute(ctx, req)
}
func (u *SystemConfigurationUsecaseImpl) DeleteByConfigKey(ctx context.Context, configKey string) error {
	return u.deleteUsecase.DeleteByConfigKey(ctx, configKey)
}

func (u *SystemConfigurationUsecaseImpl) List(ctx context.Context, pagination common.Pagination, filter repository.SystemConfigurationFilter) (common.PaginationResult[*entity.SystemConfiguration], int64, error) {
	return u.listUsecase.List(ctx, pagination, filter)
}

func (u *SystemConfigurationUsecaseImpl) GetConfigValue(ctx context.Context, configKey string) (interface{}, error) {
	config, err := u.GetByConfigKey(ctx, configKey)
	if err != nil {
		return nil, err
	}

	return config.ConfigValue.Data, nil
}
