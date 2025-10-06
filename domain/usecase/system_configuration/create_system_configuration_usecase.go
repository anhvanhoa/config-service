package system_configuration

import (
	"config-service/domain/entity"
	"config-service/domain/repository"
	"context"
)

type CreateSystemConfigurationUsecase interface {
	Execute(ctx context.Context, req CreateSystemConfigurationRequest) (*entity.SystemConfiguration, error)
}

type CreateSystemConfigurationRequest struct {
	ConfigKey       string
	ConfigValue     any
	DataType        string
	Category        string
	Description     string
	IsSystemConfig  *bool
	IsEditable      *bool
	ValidationRules any
	CreatedBy       string
}

type CreateSystemConfigurationUsecaseImpl struct {
	repo repository.SystemConfigurationRepository
}

func NewCreateSystemConfigurationUsecase(repo repository.SystemConfigurationRepository) CreateSystemConfigurationUsecase {
	return &CreateSystemConfigurationUsecaseImpl{
		repo: repo,
	}
}

func (u *CreateSystemConfigurationUsecaseImpl) Execute(ctx context.Context, req CreateSystemConfigurationRequest) (*entity.SystemConfiguration, error) {
	if err := u.validateRequest(req); err != nil {
		return nil, err
	}

	if existing, err := u.repo.CheckConfigKeyExists(ctx, req.ConfigKey); err != nil {
		return nil, err
	} else if existing {
		return nil, ErrConfigKeyAlreadyExists
	}

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

	if err := config.Validate(); err != nil {
		return nil, err
	}

	if err := u.repo.Create(ctx, config); err != nil {
		return nil, err
	}

	return config, nil
}

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
