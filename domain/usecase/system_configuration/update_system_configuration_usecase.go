package system_configuration

import (
	"config-service/domain/entity"
	"config-service/domain/repository"
	"context"
)

type UpdateSystemConfigurationUsecase interface {
	Execute(ctx context.Context, req UpdateSystemConfigurationRequest) (*entity.SystemConfiguration, error)
}

type UpdateSystemConfigurationRequest struct {
	ID              string
	ConfigValue     any
	Description     string
	IsEditable      *bool
	ValidationRules any
	UpdatedBy       string
}

type UpdateSystemConfigurationUsecaseImpl struct {
	repo repository.SystemConfigurationRepository
}

func NewUpdateSystemConfigurationUsecase(repo repository.SystemConfigurationRepository) UpdateSystemConfigurationUsecase {
	return &UpdateSystemConfigurationUsecaseImpl{
		repo: repo,
	}
}

func (u *UpdateSystemConfigurationUsecaseImpl) Execute(ctx context.Context, req UpdateSystemConfigurationRequest) (*entity.SystemConfiguration, error) {
	if err := u.validateRequest(req); err != nil {
		return nil, err
	}

	config, err := u.repo.GetByID(ctx, req.ID)
	if err != nil {
		return nil, ErrConfigNotFound
	}

	if config.IsEditable != nil && !*config.IsEditable {
		return nil, ErrConfigNotEditable
	}

	cf := &entity.SystemConfiguration{
		ID:              config.ID,
		ConfigKey:       config.ConfigKey,
		ConfigValue:     entity.JSONValue{Data: req.ConfigValue},
		Description:     req.Description,
		IsEditable:      req.IsEditable,
		DataType:        config.DataType,
		Category:        config.Category,
		IsSystemConfig:  config.IsSystemConfig,
		ValidationRules: entity.JSONValue{Data: req.ValidationRules},
		UpdatedBy:       req.UpdatedBy,
	}

	if req.IsEditable != nil {
		cf.IsEditable = req.IsEditable
	} else {
		cf.IsEditable = config.IsEditable
	}

	if err := cf.Validate(); err != nil {
		return nil, err
	}

	if err := u.repo.Update(ctx, cf); err != nil {
		return nil, err
	}

	return config, nil
}

func (u *UpdateSystemConfigurationUsecaseImpl) validateRequest(req UpdateSystemConfigurationRequest) error {
	if req.ID == "" {
		return ErrInvalidID
	}
	return nil
}
