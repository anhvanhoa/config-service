package system_configuration_service

import (
	"context"

	proto_system_configuration "github.com/anhvanhoa/sf-proto/gen/system_configuration/v1"
)

func (s *SystemConfigurationService) GetSystemConfiguration(ctx context.Context, req *proto_system_configuration.GetSystemConfigurationRequest) (*proto_system_configuration.GetSystemConfigurationResponse, error) {
	config, err := s.systemConfigurationUsecase.GetByConfigKey(ctx, req.ConfigKey)
	if err != nil {
		return nil, err
	}
	return &proto_system_configuration.GetSystemConfigurationResponse{
		Configuration: s.createEntitySystemConfigurationToProto(config),
	}, nil
}
