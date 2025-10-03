package repository

import (
	"config-service/domain/entity"
	"context"

	"github.com/anhvanhoa/service-core/common"
)

type SystemConfigurationRepository interface {
	Create(ctx context.Context, config *entity.SystemConfiguration) error
	GetByID(ctx context.Context, id string) (*entity.SystemConfiguration, error)
	GetByConfigKey(ctx context.Context, configKey string) (*entity.SystemConfiguration, error)
	List(ctx context.Context, pagination common.Pagination, filter SystemConfigurationFilter) ([]*entity.SystemConfiguration, int64, error)
	Update(ctx context.Context, config *entity.SystemConfiguration) error
	Delete(ctx context.Context, id string) error
}

type SystemConfigurationFilter struct {
	Category       string
	DataType       string
	IsSystemConfig bool
	IsEditable     bool
}
