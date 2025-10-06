package system_configuration_service

import (
	"config-service/domain/repository"
	"context"

	"github.com/anhvanhoa/service-core/common"
	proto_common "github.com/anhvanhoa/sf-proto/gen/common/v1"
	proto_system_configuration "github.com/anhvanhoa/sf-proto/gen/system_configuration/v1"
)

func (s *SystemConfigurationService) ListSystemConfiguration(ctx context.Context, req *proto_system_configuration.ListSystemConfigurationRequest) (*proto_system_configuration.ListSystemConfigurationResponse, error) {
	// Convert pagination
	pagination := common.Pagination{
		Page:     int(req.Pagination.Page),
		PageSize: int(req.Pagination.PageSize),
	}

	// Convert filter
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

	// Convert entities to proto
	protoConfigs := make([]*proto_system_configuration.SystemConfiguration, len(configs))
	for i, config := range configs {
		protoConfigs[i] = s.createEntitySystemConfigurationToProto(config)
	}

	// Create pagination response
	paginationResponse := &proto_common.PaginationResponse{
		Page:       req.Pagination.Page,
		PageSize:   req.Pagination.PageSize,
		Total:      int32(total),
		TotalPages: int32((total + int64(req.Pagination.PageSize) - 1) / int64(req.Pagination.PageSize)),
	}

	return &proto_system_configuration.ListSystemConfigurationResponse{
		Configurations: protoConfigs,
		Pagination:     paginationResponse,
	}, nil
}
