package system_configuration_service

import (
	"config-service/domain/usecase/system_configuration"
	"context"

	proto_system_configuration "github.com/anhvanhoa/sf-proto/gen/system_configuration/v1"
	anypb "google.golang.org/protobuf/types/known/anypb"
	structpb "google.golang.org/protobuf/types/known/structpb"
)

func (s *SystemConfigurationService) UpdateSystemConfiguration(ctx context.Context, req *proto_system_configuration.UpdateSystemConfigurationRequest) (*proto_system_configuration.UpdateSystemConfigurationResponse, error) {
	domainReq := s.updateSystemConfigurationRequestToDomain(req)
	config, err := s.systemConfigurationUsecase.Update(ctx, domainReq)
	if err != nil {
		return nil, err
	}
	return &proto_system_configuration.UpdateSystemConfigurationResponse{
		Configuration: s.createEntitySystemConfigurationToProto(config),
	}, nil
}

func (s *SystemConfigurationService) updateSystemConfigurationRequestToDomain(req *proto_system_configuration.UpdateSystemConfigurationRequest) system_configuration.UpdateSystemConfigurationRequest {
	domainReq := system_configuration.UpdateSystemConfigurationRequest{
		ID:        req.Id,
		UpdatedBy: req.UpdatedBy,
	}

	// Convert config_value from Any to interface{}
	if req.ConfigValue != nil {
		domainReq.ConfigValue = s.anyToInterface(req.ConfigValue)
	}

	// Set description if provided
	if req.Description != "" {
		domainReq.Description = req.Description
	}

	// Set is_editable if provided
	if req.IsEditable != nil {
		domainReq.IsEditable = req.IsEditable
	}

	// Convert validation_rules from Any to interface{}
	if req.ValidationRules != nil {
		domainReq.ValidationRules = s.anyToInterface(req.ValidationRules)
	}

	return domainReq
}

func (s *SystemConfigurationService) anyToInterface(any *anypb.Any) interface{} {
	if any == nil {
		return nil
	}

	// Try to unmarshal to a structpb.Value first
	var value structpb.Value
	if err := any.UnmarshalTo(&value); err == nil {
		// Convert structpb.Value to interface{}
		return value.AsInterface()
	}

	// If unmarshaling fails, return the raw bytes
	return any.Value
}
