package system_configuration

import (
	"config-service/domain/entity"
	"config-service/domain/repository"
	"context"

	"github.com/anhvanhoa/service-core/common"
	"github.com/anhvanhoa/service-core/utils"
)

type ListSystemConfigurationUsecase interface {
	List(ctx context.Context, pagination common.Pagination, filter repository.SystemConfigurationFilter) (common.PaginationResult[*entity.SystemConfiguration], int64, error)
}

type ListSystemConfigurationUsecaseImpl struct {
	repo   repository.SystemConfigurationRepository
	helper utils.Helper
}

func NewListSystemConfigurationUsecase(repo repository.SystemConfigurationRepository, helper utils.Helper) ListSystemConfigurationUsecase {
	return &ListSystemConfigurationUsecaseImpl{
		repo:   repo,
		helper: helper,
	}
}

func (u *ListSystemConfigurationUsecaseImpl) List(ctx context.Context, pagination common.Pagination, filter repository.SystemConfigurationFilter) (common.PaginationResult[*entity.SystemConfiguration], int64, error) {
	configs, total, err := u.repo.List(ctx, pagination, filter)
	if err != nil {
		return common.PaginationResult[*entity.SystemConfiguration]{}, 0, err
	}

	return common.PaginationResult[*entity.SystemConfiguration]{
		Data:       configs,
		Total:      total,
		TotalPages: u.helper.CalculateTotalPages(total, int64(pagination.PageSize)),
		PageSize:   pagination.PageSize,
		Page:       pagination.Page,
	}, total, nil
}
