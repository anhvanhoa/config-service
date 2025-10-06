package system_configuration_service

import (
	"context"

	proto_system_configuration "github.com/anhvanhoa/sf-proto/gen/system_configuration/v1"
)

func (s *SystemConfigurationService) DeleteSystemConfiguration(ctx context.Context, req *proto_system_configuration.DeleteSystemConfigurationRequest) (*proto_system_configuration.DeleteSystemConfigurationResponse, error) {
	err := s.systemConfigurationUsecase.DeleteByConfigKey(ctx, req.ConfigKey)
	if err != nil {
		return nil, err
	}
	return &proto_system_configuration.DeleteSystemConfigurationResponse{
		Message: "System configuration deleted successfully",
	}, nil
}
