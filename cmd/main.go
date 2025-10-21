package main

import (
	"config-service/bootstrap"
	"config-service/infrastructure/grpc_client"
	"config-service/infrastructure/grpc_service"
	system_configuration_service "config-service/infrastructure/grpc_service/system_configuration"
	"context"

	gc "github.com/anhvanhoa/service-core/domain/grpc_client"

	"github.com/anhvanhoa/service-core/domain/discovery"
)

func main() {
	StartGRPCServer()
}

func StartGRPCServer() {
	app := bootstrap.App()
	env := app.Env
	log := app.Log

	discoveryConfig := &discovery.DiscoveryConfig{
		ServiceName:   env.NameService,
		ServicePort:   env.PortGrpc,
		ServiceHost:   env.HostGprc,
		IntervalCheck: env.IntervalCheck,
		TimeoutCheck:  env.TimeoutCheck,
	}

	discovery, err := discovery.NewDiscovery(discoveryConfig)
	if err != nil {
		log.Fatal("Failed to create discovery: " + err.Error())
	}
	discovery.Register()

	clientFactory := gc.NewClientFactory(env.GrpcClients...)
	client := clientFactory.GetClient(env.PermissionServiceAddr)
	permissionClient := grpc_client.NewPermissionClient(client)

	systemConfigurationService := system_configuration_service.NewSystemConfigurationService(app.Repo.SystemConfiguration(), app.Helper)

	grpcSrv := grpc_service.NewGRPCServer(
		env, log,
		systemConfigurationService,
	)

	ctx, cancel := context.WithCancel(context.Background())
	permissions := app.Helper.ConvertResourcesToPermissions(grpcSrv.GetResources())
	if _, err := permissionClient.PermissionServiceClient.RegisterPermission(ctx, permissions); err != nil {
		log.Fatal("Failed to register permission: " + err.Error())
	}

	defer cancel()
	if err := grpcSrv.Start(ctx); err != nil {
		log.Fatal("gRPC server error: " + err.Error())
	}
}
