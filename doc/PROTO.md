syntax = "proto3";

package system_configuration.v1;

option go_package = "github.com/anhvanhoa/sf-proto/gen/system_configuration/v1;proto_system_configuration";

import "google/protobuf/timestamp.proto";
import "google/protobuf/any.proto";
import "buf/validate/validate.proto";
import "common/v1/common.proto";

message SystemConfiguration {
    string id = 1;
    string config_key = 2;
    google.protobuf.Any config_value = 3;
    string data_type = 4;
    string category = 5;
    string description = 6;
    bool is_system_config = 7;
    bool is_editable = 8;
    google.protobuf.Any validation_rules = 9;
    string created_by = 10;
    string updated_by = 11;
    google.protobuf.Timestamp created_at = 12;
    google.protobuf.Timestamp updated_at = 13;
}

message CreateSystemConfigurationRequest {
    string config_key = 1 [
        (buf.validate.field).string.min_len = 1,
        (buf.validate.field).string.max_len = 100
    ];
    google.protobuf.Any config_value = 2;
    string data_type = 3 [
        (buf.validate.field).string.min_len = 1,
        (buf.validate.field).string.max_len = 50
    ];
    string category = 4 [
        (buf.validate.field).string.min_len = 1,
        (buf.validate.field).string.max_len = 100
    ];
    string description = 5 [(buf.validate.field).string.max_len = 1000];
    bool is_system_config = 6;
    bool is_editable = 7;
    google.protobuf.Any validation_rules = 8;
    string created_by = 9 [
        (buf.validate.field).string.uuid = true
    ];
}

message CreateSystemConfigurationResponse {
    SystemConfiguration configuration = 1;
}

message GetSystemConfigurationRequest {
    string config_key = 1 [(buf.validate.field).string.min_len = 1];
}

message GetSystemConfigurationResponse {
    SystemConfiguration configuration = 1;
}

message UpdateSystemConfigurationRequest {
    string id = 1 [
        (buf.validate.field).string.min_len = 1,
        (buf.validate.field).ignore = IGNORE_ALWAYS
    ];
    google.protobuf.Any config_value = 2 [
        (buf.validate.field).ignore = IGNORE_ALWAYS
    ];
    string description = 3 [
        (buf.validate.field).string.max_len = 1000,
        (buf.validate.field).ignore = IGNORE_ALWAYS
    ];
    optional bool is_editable = 4;
    google.protobuf.Any validation_rules = 5 [
        (buf.validate.field).ignore = IGNORE_ALWAYS
    ];
    string updated_by = 6 [
        (buf.validate.field).string.min_len = 1,
        (buf.validate.field).string.max_len = 100,
        (buf.validate.field).ignore = IGNORE_ALWAYS
    ];
}

message UpdateSystemConfigurationResponse {
    SystemConfiguration configuration = 1;
}

message DeleteSystemConfigurationRequest {
    string config_key = 1 [(buf.validate.field).string.min_len = 1];
}

message DeleteSystemConfigurationResponse {
    string message = 1;
}

message ListSystemConfigurationRequest {
    common.PaginationRequest pagination = 1;
    string category = 2 [
        (buf.validate.field).string.max_len = 100,
        (buf.validate.field).ignore = IGNORE_ALWAYS
    ];
    string data_type = 3 [
        (buf.validate.field).string.max_len = 50,
        (buf.validate.field).ignore = IGNORE_ALWAYS
    ];
    optional bool is_system_config = 4;
    optional bool is_editable = 5;
}

message ListSystemConfigurationResponse {
    repeated SystemConfiguration configurations = 1;
    common.PaginationResponse pagination = 2;
}


service SystemConfigurationService {
    rpc CreateSystemConfiguration(CreateSystemConfigurationRequest) returns (CreateSystemConfigurationResponse);
    rpc GetSystemConfiguration(GetSystemConfigurationRequest) returns (GetSystemConfigurationResponse);
    rpc UpdateSystemConfiguration(UpdateSystemConfigurationRequest) returns (UpdateSystemConfigurationResponse);
    rpc DeleteSystemConfiguration(DeleteSystemConfigurationRequest) returns (DeleteSystemConfigurationResponse);
    rpc ListSystemConfiguration(ListSystemConfigurationRequest) returns (ListSystemConfigurationResponse);
}
