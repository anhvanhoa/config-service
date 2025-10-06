package system_configuration

import (
	"config-service/domain/entity"
	"config-service/domain/repository"
	"context"

	"github.com/anhvanhoa/service-core/common"
)

type ListSystemConfigurationUsecase interface {
	List(ctx context.Context, pagination common.Pagination, filter repository.SystemConfigurationFilter) ([]*entity.SystemConfiguration, int64, error)
}

type ListSystemConfigurationUsecaseImpl struct {
	repo repository.SystemConfigurationRepository
}

func NewListSystemConfigurationUsecase(repo repository.SystemConfigurationRepository) ListSystemConfigurationUsecase {
	return &ListSystemConfigurationUsecaseImpl{
		repo: repo,
	}
}

func (u *ListSystemConfigurationUsecaseImpl) List(ctx context.Context, pagination common.Pagination, filter repository.SystemConfigurationFilter) ([]*entity.SystemConfiguration, int64, error) {
	configs, total, err := u.repo.List(ctx, pagination, filter)
	if err != nil {
		return nil, 0, err
	}

	return configs, total, nil
}
