package system_configuration

import (
	"config-service/domain/entity"
	"config-service/domain/repository"
	"context"
)

type UpdateSystemConfigurationUsecase interface {
	Update(ctx context.Context, req UpdateSystemConfigurationRequest) (*entity.SystemConfiguration, error)
}

type UpdateSystemConfigurationRequest struct {
	ID              string      `json:"id" validate:"required"`
	ConfigValue     interface{} `json:"config_value,omitempty"`
	Description     string      `json:"description,omitempty"`
	IsEditable      *bool       `json:"is_editable,omitempty"`
	ValidationRules interface{} `json:"validation_rules,omitempty"`
	UpdatedBy       string      `json:"updated_by"`
}

type UpdateSystemConfigurationUsecaseImpl struct {
	repo repository.SystemConfigurationRepository
}

func NewUpdateSystemConfigurationUsecase(repo repository.SystemConfigurationRepository) UpdateSystemConfigurationUsecase {
	return &UpdateSystemConfigurationUsecaseImpl{
		repo: repo,
	}
}

func (u *UpdateSystemConfigurationUsecaseImpl) Update(ctx context.Context, req UpdateSystemConfigurationRequest) (*entity.SystemConfiguration, error) {
	if err := u.validateRequest(req); err != nil {
		return nil, err
	}

	config, err := u.repo.GetByID(ctx, req.ID)
	if err != nil {
		return nil, ErrConfigNotFound
	}

	if !config.IsEditable {
		return nil, ErrConfigNotEditable
	}

	if err := config.Validate(); err != nil {
		return nil, err
	}

	if err := u.repo.Update(ctx, config); err != nil {
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
