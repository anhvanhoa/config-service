package main

import (
	"config-service/bootstrap"
	"config-service/infrastructure/grpc_client"
	"config-service/infrastructure/grpc_service"
	system_configuration_service "config-service/infrastructure/grpc_service/system_configuration"
	"context"

	gc "github.com/anhvanhoa/service-core/domain/grpc_client"
)

func main() {
	StartGRPCServer()
}

func StartGRPCServer() {
	app := bootstrap.App()
	env := app.Env
	log := app.Log

	clientFactory := gc.NewClientFactory(env.GrpcClients...)
	client := clientFactory.GetClient(env.PermissionServiceAddr)
	permissionClient := grpc_client.NewPermissionClient(client)

	systemConfigurationService := system_configuration_service.NewSystemConfigurationService(app.Repo.SystemConfiguration(), app.Helper)

	grpcSrv := grpc_service.NewGRPCServer(
		env, log, app.Cacher,
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
