package system_configuration

import (
	"config-service/domain/entity"
	"config-service/domain/repository"
	"context"
)

// UpdateSystemConfigurationUsecase defines the interface for updating system configurations
type UpdateSystemConfigurationUsecase interface {
	Update(ctx context.Context, req UpdateSystemConfigurationRequest) (*entity.SystemConfiguration, error)
	UpdateConfigValue(ctx context.Context, configKey string, value interface{}) (*entity.SystemConfiguration, error)
}

// UpdateSystemConfigurationRequest represents the request to update a system configuration
type UpdateSystemConfigurationRequest struct {
	ID              string      `json:"id" validate:"required"`
	ConfigValue     interface{} `json:"config_value,omitempty"`
	Description     string      `json:"description,omitempty"`
	IsEditable      *bool       `json:"is_editable,omitempty"`
	ValidationRules interface{} `json:"validation_rules,omitempty"`
	UpdatedBy       string      `json:"updated_by"`
}

// UpdateSystemConfigurationUsecaseImpl implements UpdateSystemConfigurationUsecase
type UpdateSystemConfigurationUsecaseImpl struct {
	repo repository.SystemConfigurationRepository
}

// NewUpdateSystemConfigurationUsecase creates a new update system configuration usecase
func NewUpdateSystemConfigurationUsecase(repo repository.SystemConfigurationRepository) UpdateSystemConfigurationUsecase {
	return &UpdateSystemConfigurationUsecaseImpl{
		repo: repo,
	}
}

// Update updates a system configuration
func (u *UpdateSystemConfigurationUsecaseImpl) Update(ctx context.Context, req UpdateSystemConfigurationRequest) (*entity.SystemConfiguration, error) {
	// Validate request
	if err := u.validateRequest(req); err != nil {
		return nil, err
	}

	// Get existing config
	config, err := u.repo.GetByID(ctx, req.ID)
	if err != nil {
		return nil, err
	}

	if config == nil {
		return nil, ErrConfigNotFound
	}

	// Check if config is editable
	if !config.IsEditable {
		return nil, ErrConfigNotEditable
	}

	// Update fields
	if req.ConfigValue != nil {
		config.ConfigValue = entity.JSONValue{Data: req.ConfigValue}
	}
	if req.Description != "" {
		config.Description = req.Description
	}
	if req.IsEditable != nil {
		config.IsEditable = *req.IsEditable
	}
	if req.ValidationRules != nil {
		config.ValidationRules = entity.JSONValue{Data: req.ValidationRules}
	}
	if req.UpdatedBy != "" {
		config.UpdatedBy = req.UpdatedBy
	}

	// Validate updated entity
	if err := config.Validate(); err != nil {
		return nil, err
	}

	// Update in repository
	if err := u.repo.Update(ctx, config); err != nil {
		return nil, err
	}

	return config, nil
}

// UpdateConfigValue updates only the config value
func (u *UpdateSystemConfigurationUsecaseImpl) UpdateConfigValue(ctx context.Context, configKey string, value interface{}) (*entity.SystemConfiguration, error) {
	if configKey == "" {
		return nil, ErrInvalidConfigKey
	}

	// Get existing config
	config, err := u.repo.GetByConfigKey(ctx, configKey)
	if err != nil {
		return nil, err
	}

	if config == nil {
		return nil, ErrConfigNotFound
	}

	// Check if config is editable
	if !config.IsEditable {
		return nil, ErrConfigNotEditable
	}

	// Update config value
	config.ConfigValue = entity.JSONValue{Data: value}

	// Update in repository
	if err := u.repo.Update(ctx, config); err != nil {
		return nil, err
	}

	return config, nil
}

// validateRequest validates the update request
func (u *UpdateSystemConfigurationUsecaseImpl) validateRequest(req UpdateSystemConfigurationRequest) error {
	if req.ID == "" {
		return ErrInvalidID
	}
	return nil
}
