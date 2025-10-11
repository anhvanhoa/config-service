package repo

import (
	"config-service/domain/entity"
	"config-service/domain/repository"
	"context"
	"fmt"

	"github.com/anhvanhoa/service-core/common"
	"github.com/anhvanhoa/service-core/utils"
	"github.com/go-pg/pg/v10"
)

type systemConfigurationRepository struct {
	db     *pg.DB
	helper utils.Helper
}

func NewSystemConfigurationRepository(
	db *pg.DB,
	helper utils.Helper,
) repository.SystemConfigurationRepository {
	return &systemConfigurationRepository{
		db:     db,
		helper: helper,
	}
}

func (r *systemConfigurationRepository) Create(ctx context.Context, config *entity.SystemConfiguration) error {
	_, err := r.db.Model(config).Context(ctx).Insert()
	return err
}

func (r *systemConfigurationRepository) GetByID(ctx context.Context, id string) (*entity.SystemConfiguration, error) {
	config := &entity.SystemConfiguration{}
	err := r.db.Model(config).Context(ctx).Where("id = ?", id).Select()
	if err != nil {
		return nil, err
	}
	return config, nil
}

func (r *systemConfigurationRepository) GetByConfigKey(ctx context.Context, configKey string) (*entity.SystemConfiguration, error) {
	config := &entity.SystemConfiguration{}
	err := r.db.Model(config).Context(ctx).Where("config_key = ?", configKey).Select()
	if err != nil {
		return nil, err
	}
	return config, nil
}

func (r *systemConfigurationRepository) List(ctx context.Context, pagination common.Pagination, filter repository.SystemConfigurationFilter) ([]*entity.SystemConfiguration, int64, error) {
	var configs []*entity.SystemConfiguration
	query := r.db.Model(&configs).Context(ctx)

	if filter.Category != "" {
		query = query.Where("category = ?", filter.Category)
	}
	if filter.DataType != "" {
		query = query.Where("data_type = ?", filter.DataType)
	}
	if filter.IsSystemConfig {
		query = query.Where("is_system_config = ?", filter.IsSystemConfig)
	}
	if filter.IsEditable {
		query = query.Where("is_editable = ?", filter.IsEditable)
	}

	totalCount, err := query.Count()
	if err != nil {
		return nil, 0, err
	}

	if pagination.Page > 0 && pagination.PageSize > 0 {
		offset := r.helper.CalculateOffset(pagination.Page, pagination.PageSize)
		query = query.Limit(pagination.PageSize).Offset(offset)
	}

	if pagination.SortBy != "" {
		sortOrder := "ASC"
		if pagination.SortOrder == "desc" || pagination.SortOrder == "DESC" {
			sortOrder = "DESC"
		}
		query = query.Order(fmt.Sprintf("%s %s", pagination.SortBy, sortOrder))
	} else {
		query = query.Order("created_at DESC")
	}

	err = query.Select()
	if err != nil {
		return nil, 0, err
	}

	return configs, int64(totalCount), nil
}

func (r *systemConfigurationRepository) Update(ctx context.Context, config *entity.SystemConfiguration) error {
	_, err := r.db.Model(config).Context(ctx).WherePK().Update()
	return err
}

func (r *systemConfigurationRepository) Delete(ctx context.Context, id string) error {
	config := &entity.SystemConfiguration{ID: id}
	_, err := r.db.Model(config).Context(ctx).WherePK().Delete()
	return err
}

func (r *systemConfigurationRepository) CheckConfigKeyExists(ctx context.Context, configKey string) (bool, error) {
	count, err := r.db.Model(&entity.SystemConfiguration{}).Context(ctx).Where("config_key = ?", configKey).Count()
	if err != nil {
		return false, err
	}
	return count > 0, nil
}
