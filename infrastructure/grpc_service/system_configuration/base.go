package system_configuration_service

import (
	"config-service/domain/repository"
	"config-service/domain/usecase/system_configuration"

	"github.com/anhvanhoa/service-core/utils"
	proto_system_configuration "github.com/anhvanhoa/sf-proto/gen/system_configuration/v1"
)

type SystemConfigurationService struct {
	systemConfigurationUsecase system_configuration.SystemConfigurationUsecase
	proto_system_configuration.UnsafeSystemConfigurationServiceServer
}

func NewSystemConfigurationService(repos repository.SystemConfigurationRepository, helper utils.Helper) proto_system_configuration.SystemConfigurationServiceServer {
	systemConfigurationUsecase := system_configuration.NewSystemConfigurationUsecase(repos, helper)
	return &SystemConfigurationService{
		systemConfigurationUsecase: systemConfigurationUsecase,
	}
}
