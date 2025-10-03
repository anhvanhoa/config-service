package system_configuration

import (
	"config-service/domain/entity"
	"config-service/domain/repository"
	"context"
)

// CreateSystemConfigurationUsecase defines the interface for creating system configurations
type CreateSystemConfigurationUsecase interface {
	Execute(ctx context.Context, req CreateSystemConfigurationRequest) (*entity.SystemConfiguration, error)
}

// CreateSystemConfigurationRequest represents the request to create a system configuration
type CreateSystemConfigurationRequest struct {
	ConfigKey       string      `json:"config_key" validate:"required"`
	ConfigValue     interface{} `json:"config_value" validate:"required"`
	DataType        string      `json:"data_type" validate:"required,oneof=string number boolean json array"`
	Category        string      `json:"category" validate:"required,oneof=irrigation fertilization alerts sensors reports"`
	Description     string      `json:"description"`
	IsSystemConfig  bool        `json:"is_system_config"`
	IsEditable      bool        `json:"is_editable"`
	ValidationRules interface{} `json:"validation_rules,omitempty"`
	CreatedBy       string      `json:"created_by"`
}

// CreateSystemConfigurationUsecaseImpl implements CreateSystemConfigurationUsecase
type CreateSystemConfigurationUsecaseImpl struct {
	repo repository.SystemConfigurationRepository
}

// NewCreateSystemConfigurationUsecase creates a new create system configuration usecase
func NewCreateSystemConfigurationUsecase(repo repository.SystemConfigurationRepository) CreateSystemConfigurationUsecase {
	return &CreateSystemConfigurationUsecaseImpl{
		repo: repo,
	}
}

// Execute creates a new system configuration
func (u *CreateSystemConfigurationUsecaseImpl) Execute(ctx context.Context, req CreateSystemConfigurationRequest) (*entity.SystemConfiguration, error) {
	// Validate request
	if err := u.validateRequest(req); err != nil {
		return nil, err
	}

	// Check if config key already exists
	existing, err := u.repo.GetByConfigKey(ctx, req.ConfigKey)
	if err == nil && existing != nil {
		return nil, ErrConfigKeyAlreadyExists
	}

	// Create entity
	config := &entity.SystemConfiguration{
		ConfigKey:       req.ConfigKey,
		ConfigValue:     entity.JSONValue{Data: req.ConfigValue},
		DataType:        req.DataType,
		Category:        req.Category,
		Description:     req.Description,
		IsSystemConfig:  req.IsSystemConfig,
		IsEditable:      req.IsEditable,
		ValidationRules: entity.JSONValue{Data: req.ValidationRules},
		CreatedBy:       req.CreatedBy,
		UpdatedBy:       req.CreatedBy,
	}

	// Validate entity
	if err := config.Validate(); err != nil {
		return nil, err
	}

	// Create in repository
	if err := u.repo.Create(ctx, config); err != nil {
		return nil, err
	}

	return config, nil
}

// validateRequest validates the create request
func (u *CreateSystemConfigurationUsecaseImpl) validateRequest(req CreateSystemConfigurationRequest) error {
	if req.ConfigKey == "" {
		return ErrInvalidConfigKey
	}
	if req.DataType == "" {
		return ErrInvalidDataType
	}
	if req.Category == "" {
		return ErrInvalidCategory
	}
	return nil
}
