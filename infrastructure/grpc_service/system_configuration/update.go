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
		ID:          req.Id,
		UpdatedBy:   req.UpdatedBy,
		Description: req.Description,
		IsEditable:  req.IsEditable,
	}

	if req.ConfigValue != nil {
		domainReq.ConfigValue = s.anyToInterface(req.ConfigValue)
	}

	if req.ValidationRules != nil {
		domainReq.ValidationRules = s.anyToInterface(req.ValidationRules)
	}

	return domainReq
}

func (s *SystemConfigurationService) anyToInterface(any *anypb.Any) interface{} {
	if any == nil {
		return nil
	}

	var value structpb.Value
	if err := any.UnmarshalTo(&value); err == nil {
		return value.AsInterface()
	}

	return any.Value
}
