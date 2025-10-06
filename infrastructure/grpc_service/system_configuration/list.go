package system_configuration_service

import (
	"config-service/domain/repository"
	"context"

	"github.com/anhvanhoa/service-core/common"
	proto_common "github.com/anhvanhoa/sf-proto/gen/common/v1"
	proto_system_configuration "github.com/anhvanhoa/sf-proto/gen/system_configuration/v1"
)

func (s *SystemConfigurationService) ListSystemConfiguration(ctx context.Context, req *proto_system_configuration.ListSystemConfigurationRequest) (*proto_system_configuration.ListSystemConfigurationResponse, error) {
	pagination := common.Pagination{
		Page:     int(req.Pagination.Page),
		PageSize: int(req.Pagination.PageSize),
	}

	filter := repository.SystemConfigurationFilter{
		Category:       req.Category,
		DataType:       req.DataType,
		IsSystemConfig: req.IsSystemConfig != nil && *req.IsSystemConfig,
		IsEditable:     req.IsEditable != nil && *req.IsEditable,
	}

	configs, total, err := s.systemConfigurationUsecase.List(ctx, pagination, filter)
	if err != nil {
		return nil, err
	}

	protoConfigs := make([]*proto_system_configuration.SystemConfiguration, len(configs.Data))
	for i, config := range configs.Data {
		protoConfigs[i] = s.createEntitySystemConfigurationToProto(config)
	}

	paginationResponse := &proto_common.PaginationResponse{
		Page:       int32(configs.Page),
		PageSize:   int32(configs.PageSize),
		Total:      int32(total),
		TotalPages: int32(configs.TotalPages),
	}

	return &proto_system_configuration.ListSystemConfigurationResponse{
		Configurations: protoConfigs,
		Pagination:     paginationResponse,
	}, nil
}
