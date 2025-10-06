package grpc_service

import (
	"config-service/bootstrap"

	grpc_server "github.com/anhvanhoa/service-core/bootstrap/grpc"
	"github.com/anhvanhoa/service-core/domain/log"
	proto_system_configuration "github.com/anhvanhoa/sf-proto/gen/system_configuration/v1"
	"google.golang.org/grpc"
)

func NewGRPCServer(
	env *bootstrap.Env,
	log *log.LogGRPCImpl,
	SystemConfigurationServer proto_system_configuration.SystemConfigurationServiceServer,
) *grpc_server.GRPCServer {
	config := &grpc_server.GRPCServerConfig{
		IsProduction: env.IsProduction(),
		PortGRPC:     env.PortGrpc,
		NameService:  env.NameService,
	}
	return grpc_server.NewGRPCServer(
		config,
		log,
		func(server *grpc.Server) {
			proto_system_configuration.RegisterSystemConfigurationServiceServer(server, SystemConfigurationServer)
		},
	)
}
