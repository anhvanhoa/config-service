package system_configuration_service

import (
	"config-service/domain/usecase/system_configuration"

	proto_system_configuration "github.com/anhvanhoa/sf-proto/gen/system_configuration/v1"
)

type SystemConfigurationService struct {
	systemConfigurationUsecase system_configuration.SystemConfigurationUsecase
	proto_system_configuration.UnsafeSystemConfigurationServiceServer
}

func NewSystemConfigurationService(systemConfigurationUsecase system_configuration.SystemConfigurationUsecase) proto_system_configuration.SystemConfigurationServiceServer {
	return &SystemConfigurationService{
		systemConfigurationUsecase: systemConfigurationUsecase,
	}
}
