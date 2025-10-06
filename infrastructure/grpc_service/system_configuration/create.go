package system_configuration_service

import (
	"config-service/domain/entity"
	"config-service/domain/usecase/system_configuration"
	"context"

	proto_system_configuration "github.com/anhvanhoa/sf-proto/gen/system_configuration/v1"
	anypb "google.golang.org/protobuf/types/known/anypb"
	structpb "google.golang.org/protobuf/types/known/structpb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (s *SystemConfigurationService) CreateSystemConfiguration(ctx context.Context, req *proto_system_configuration.CreateSystemConfigurationRequest) (*proto_system_configuration.CreateSystemConfigurationResponse, error) {
	domainReq := s.createSystemConfigurationRequestToDomain(req)
	config, err := s.systemConfigurationUsecase.Create(ctx, domainReq)
	if err != nil {
		return nil, err
	}
	return &proto_system_configuration.CreateSystemConfigurationResponse{
		Configuration: s.createEntitySystemConfigurationToProto(config),
	}, nil
}

func (s *SystemConfigurationService) createSystemConfigurationRequestToDomain(req *proto_system_configuration.CreateSystemConfigurationRequest) system_configuration.CreateSystemConfigurationRequest {
	return system_configuration.CreateSystemConfigurationRequest{
		ConfigKey:       req.ConfigKey,
		ConfigValue:     req.ConfigValue,
		DataType:        req.DataType,
		Category:        req.Category,
		Description:     req.Description,
		IsSystemConfig:  req.IsSystemConfig,
		IsEditable:      req.IsEditable,
		ValidationRules: req.ValidationRules,
		CreatedBy:       req.CreatedBy,
	}
}

func (s *SystemConfigurationService) createEntitySystemConfigurationToProto(config *entity.SystemConfiguration) *proto_system_configuration.SystemConfiguration {
	configValueAny := toAny(config.ConfigValue.Data)
	validationRulesAny := toAny(config.ValidationRules.Data)
	protoConfig := &proto_system_configuration.SystemConfiguration{
		Id:              config.ID,
		ConfigKey:       config.ConfigKey,
		DataType:        config.DataType,
		Category:        config.Category,
		Description:     config.Description,
		IsSystemConfig:  config.IsSystemConfig,
		IsEditable:      config.IsEditable,
		CreatedBy:       config.CreatedBy,
		UpdatedBy:       config.UpdatedBy,
		CreatedAt:       timestamppb.New(config.CreatedAt),
		UpdatedAt:       timestamppb.New(config.UpdatedAt),
		ConfigValue:     configValueAny,
		ValidationRules: validationRulesAny,
	}

	return protoConfig
}

func toAny(v any) *anypb.Any {
	if v == nil {
		return nil
	}
	val, err := structpb.NewValue(v)
	if err != nil {
		return nil
	}
	anyMsg, err := anypb.New(val)
	if err != nil {
		return nil
	}
	return anyMsg
}
