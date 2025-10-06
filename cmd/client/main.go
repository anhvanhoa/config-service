package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"encoding/json"

	proto_common "github.com/anhvanhoa/sf-proto/gen/common/v1"
	proto_system_configuration "github.com/anhvanhoa/sf-proto/gen/system_configuration/v1"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/anypb"
	structpb "google.golang.org/protobuf/types/known/structpb"
)

var serverAddress string

func init() {
	viper.SetConfigFile("dev.config.yml")
	viper.ReadInConfig()
	serverAddress = fmt.Sprintf("%s:%s", viper.GetString("host_grpc"), viper.GetString("port_grpc"))
}

func inputPaging(reader *bufio.Reader) (int32, int32) {
	fmt.Print("Nhập trang (mặc định 1): ")
	offsetStr, _ := reader.ReadString('\n')
	offsetStr = cleanInput(offsetStr)
	offset := int32(1)
	if offsetStr != "" {
		if o, err := strconv.Atoi(offsetStr); err == nil {
			offset = int32(o)
		}
	}

	fmt.Print("Nhập số bản ghi mỗi trang (mặc định 10): ")
	limitStr, _ := reader.ReadString('\n')
	limitStr = cleanInput(limitStr)
	limit := int32(10)
	if limitStr != "" {
		if l, err := strconv.Atoi(limitStr); err == nil {
			limit = int32(l)
		}
	}

	return offset, limit
}

type SystemConfigurationClient struct {
	systemConfigurationClient proto_system_configuration.SystemConfigurationServiceClient
	conn                      *grpc.ClientConn
}

func NewSystemConfigurationClient(address string) (*SystemConfigurationClient, error) {
	conn, err := grpc.NewClient(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, fmt.Errorf("failed to connect to gRPC server: %v", err)
	}

	return &SystemConfigurationClient{
		systemConfigurationClient: proto_system_configuration.NewSystemConfigurationServiceClient(conn),
		conn:                      conn,
	}, nil
}

func (c *SystemConfigurationClient) Close() {
	if c.conn != nil {
		c.conn.Close()
	}
}

// --- Helper để làm sạch input ---
func cleanInput(s string) string {
	return strings.ToValidUTF8(strings.TrimSpace(s), "")
}

func anyToString(a *anypb.Any) string {
	if a == nil {
		return ""
	}
	var v structpb.Value
	if err := a.UnmarshalTo(&v); err == nil {
		if b, err := json.Marshal(v.AsInterface()); err == nil {
			return string(b)
		}
		return fmt.Sprintf("%v", v.AsInterface())
	}
	return fmt.Sprintf("[Any type=%s bytes=%d]", a.TypeUrl, len(a.Value))
}

func parseInputToAny(s string) *anypb.Any {
	if s == "" {
		return nil
	}
	var v interface{}
	if err := json.Unmarshal([]byte(s), &v); err == nil {
		if val, err := structpb.NewValue(v); err == nil {
			if anyMsg, err := anypb.New(val); err == nil {
				return anyMsg
			}
		}
	}
	// fallback to plain string
	anyMsg, _ := anypb.New(structpb.NewStringValue(s))
	return anyMsg
}

// ================== System Configuration Service Tests ==================

func (c *SystemConfigurationClient) TestCreateSystemConfiguration() {
	fmt.Println("\n=== Tạo System Configuration ===")

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Nhập config_key: ")
	configKey, _ := reader.ReadString('\n')
	configKey = cleanInput(configKey)

	fmt.Print("Nhập data_type (string/number/boolean/json/array): ")
	dataType, _ := reader.ReadString('\n')
	dataType = cleanInput(dataType)

	fmt.Print("Nhập category (irrigation/fertilization/alerts/sensors/reports): ")
	category, _ := reader.ReadString('\n')
	category = cleanInput(category)

	fmt.Print("Nhập description: ")
	description, _ := reader.ReadString('\n')
	description = cleanInput(description)

	fmt.Print("Is system config? (true/false): ")
	isSysStr, _ := reader.ReadString('\n')
	isSysStr = cleanInput(isSysStr)
	isSystemConfig := isSysStr == "true"

	fmt.Print("Is editable? (true/false): ")
	isEditableStr, _ := reader.ReadString('\n')
	isEditableStr = cleanInput(isEditableStr)
	isEditable := isEditableStr != "false"

	fmt.Print("Nhập giá trị config_value (hỗ trợ JSON, ví dụ: 123, true, \"abc\", {\"a\":1}): ")
	configValueStr, _ := reader.ReadString('\n')
	configValueStr = cleanInput(configValueStr)
	configValueAny := parseInputToAny(configValueStr)

	fmt.Print("Nhập created_by (uuid): ")
	createdBy, _ := reader.ReadString('\n')
	createdBy = cleanInput(createdBy)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	req := &proto_system_configuration.CreateSystemConfigurationRequest{
		ConfigKey:       configKey,
		ConfigValue:     configValueAny,
		DataType:        dataType,
		Category:        category,
		Description:     description,
		IsSystemConfig:  isSystemConfig,
		IsEditable:      isEditable,
		ValidationRules: nil,
		CreatedBy:       createdBy,
	}

	resp, err := c.systemConfigurationClient.CreateSystemConfiguration(ctx, req)
	if err != nil {
		fmt.Printf("Error CreateSystemConfiguration: %v\n", err)
		return
	}

	fmt.Printf("Kết quả:\n")
	if resp.Configuration != nil {
		cfg := resp.Configuration
		fmt.Printf("ID: %s\n", cfg.Id)
		fmt.Printf("Key: %s\n", cfg.ConfigKey)
		fmt.Printf("DataType: %s\n", cfg.DataType)
		fmt.Printf("Category: %s\n", cfg.Category)
		fmt.Printf("Description: %s\n", cfg.Description)
		fmt.Printf("IsSystem: %t, IsEditable: %t\n", cfg.IsSystemConfig, cfg.IsEditable)
		fmt.Printf("Value: %s\n", anyToString(cfg.ConfigValue))
	}
}

func (c *SystemConfigurationClient) TestGetSystemConfiguration() {
	fmt.Println("\n=== Lấy System Configuration theo config_key ===")

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Nhập config_key: ")
	configKey, _ := reader.ReadString('\n')
	configKey = cleanInput(configKey)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	resp, err := c.systemConfigurationClient.GetSystemConfiguration(ctx, &proto_system_configuration.GetSystemConfigurationRequest{
		ConfigKey: configKey,
	})
	if err != nil {
		fmt.Printf("Error GetSystemConfiguration: %v\n", err)
		return
	}

	if resp.Configuration != nil {
		cfg := resp.Configuration
		fmt.Printf("ID: %s\n", cfg.Id)
		fmt.Printf("Key: %s\n", cfg.ConfigKey)
		fmt.Printf("Value: %s\n", anyToString(cfg.ConfigValue))
		fmt.Printf("DataType: %s\n", cfg.DataType)
		fmt.Printf("Category: %s\n", cfg.Category)
		fmt.Printf("Description: %s\n", cfg.Description)
		fmt.Printf("IsSystem: %t, IsEditable: %t\n", cfg.IsSystemConfig, cfg.IsEditable)
		fmt.Printf("ValidationRules: %s\n", anyToString(cfg.ValidationRules))
		fmt.Printf("CreatedBy: %s\n", cfg.CreatedBy)
		fmt.Printf("UpdatedBy: %s\n", cfg.UpdatedBy)
		fmt.Printf("CreatedAt: %s\n", cfg.CreatedAt.AsTime())
		fmt.Printf("UpdatedAt: %s\n", cfg.UpdatedAt.AsTime())
	}
}

func (c *SystemConfigurationClient) TestListSystemConfiguration() {
	fmt.Println("\n=== Liệt kê System Configuration ===")

	reader := bufio.NewReader(os.Stdin)
	offset, limit := inputPaging(reader)

	fmt.Println("\n--- Tùy chọn bộ lọc (để trống để bỏ qua) ---")

	fmt.Print("Nhập category: ")
	category, _ := reader.ReadString('\n')
	category = cleanInput(category)

	fmt.Print("Nhập data_type: ")
	dataType, _ := reader.ReadString('\n')
	dataType = cleanInput(dataType)

	fmt.Print("is_system_config? (true/false, trống = bỏ qua): ")
	isSysStr, _ := reader.ReadString('\n')
	isSysStr = cleanInput(isSysStr)
	var isSystemConfig *bool
	if isSysStr != "" {
		b := isSysStr == "true"
		isSystemConfig = &b
	}

	fmt.Print("is_editable? (true/false, trống = bỏ qua): ")
	isEditableStr, _ := reader.ReadString('\n')
	isEditableStr = cleanInput(isEditableStr)
	var isEditable *bool
	if isEditableStr != "" {
		b := isEditableStr == "true"
		isEditable = &b
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	req := &proto_system_configuration.ListSystemConfigurationRequest{
		Pagination: &proto_common.PaginationRequest{Page: offset, PageSize: limit},
		Category:   category,
		DataType:   dataType,
	}
	if isSystemConfig != nil {
		req.IsSystemConfig = isSystemConfig
	}
	if isEditable != nil {
		req.IsEditable = isEditable
	}

	resp, err := c.systemConfigurationClient.ListSystemConfiguration(ctx, req)
	if err != nil {
		fmt.Printf("Error ListSystemConfiguration: %v\n", err)
		return
	}

	fmt.Printf("Tổng số: %d\n", resp.Pagination.Total)
	for i, cfg := range resp.Configurations {
		fmt.Printf("  [%d] ID: %s, Key: %s, Type: %s, Category: %s, Editable: %t\n", i+1, cfg.Id, cfg.ConfigKey, cfg.DataType, cfg.Category, cfg.IsEditable)
		fmt.Printf("      Value: %s\n", anyToString(cfg.ConfigValue))
		fmt.Printf("      Description: %s\n", cfg.Description)
		fmt.Printf("      IsSystem: %t, IsEditable: %t\n", cfg.IsSystemConfig, cfg.IsEditable)
		fmt.Printf("      ValidationRules: %s\n", anyToString(cfg.ValidationRules))
		fmt.Printf("      CreatedBy: %s\n", cfg.CreatedBy)
		fmt.Printf("      UpdatedBy: %s\n", cfg.UpdatedBy)
		fmt.Printf("      CreatedAt: %s\n", cfg.CreatedAt.AsTime())
		fmt.Printf("      UpdatedAt: %s\n", cfg.UpdatedAt.AsTime())
		fmt.Printf("      Value: %s\n", anyToString(cfg.ConfigValue))
	}
}

func (c *SystemConfigurationClient) TestUpdateSystemConfiguration() {
	fmt.Println("\n=== Cập nhật System Configuration ===")

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Nhập ID cấu hình: ")
	id, _ := reader.ReadString('\n')
	id = cleanInput(id)

	fmt.Print("Nhập description (trống = bỏ qua): ")
	description, _ := reader.ReadString('\n')
	description = cleanInput(description)

	fmt.Print("is_editable? (true/false, trống = bỏ qua): ")
	isEditableStr, _ := reader.ReadString('\n')
	isEditableStr = cleanInput(isEditableStr)
	var isEditable *bool
	if isEditableStr != "" {
		b := isEditableStr == "true"
		isEditable = &b
	}

	fmt.Print("Giá trị mới cho config_value (hỗ trợ JSON, trống = bỏ qua): ")
	valueStr, _ := reader.ReadString('\n')
	valueStr = cleanInput(valueStr)
	var valueAny *anypb.Any
	if valueStr != "" {
		valueAny = parseInputToAny(valueStr)
	}

	fmt.Print("Nhập updated_by: ")
	updatedBy, _ := reader.ReadString('\n')
	updatedBy = cleanInput(updatedBy)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	req := &proto_system_configuration.UpdateSystemConfigurationRequest{
		Id:          id,
		Description: description,
		UpdatedBy:   updatedBy,
	}
	if isEditable != nil {
		req.IsEditable = isEditable
	}
	if valueAny != nil {
		req.ConfigValue = valueAny
	}

	resp, err := c.systemConfigurationClient.UpdateSystemConfiguration(ctx, req)
	if err != nil {
		fmt.Printf("Error UpdateSystemConfiguration: %v\n", err)
		return
	}

	if resp.Configuration != nil {
		cfg := resp.Configuration
		fmt.Printf("Đã cập nhật: ID: %s, Key: %s, Editable: %t\n", cfg.Id, cfg.ConfigKey, cfg.IsEditable)
		fmt.Printf("Giá trị: %s\n", anyToString(cfg.ConfigValue))
	}
}

func (c *SystemConfigurationClient) TestDeleteSystemConfiguration() {
	fmt.Println("\n=== Xóa System Configuration theo config_key ===")
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Nhập config_key: ")
	configKey, _ := reader.ReadString('\n')
	configKey = cleanInput(configKey)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	resp, err := c.systemConfigurationClient.DeleteSystemConfiguration(ctx, &proto_system_configuration.DeleteSystemConfigurationRequest{ConfigKey: configKey})
	if err != nil {
		fmt.Printf("Error DeleteSystemConfiguration: %v\n", err)
		return
	}
	fmt.Printf("Kết quả: %s\n", resp.Message)
}

// ================== Menu Functions ==================

func printMainMenu() {
	fmt.Println("\n=== Ứng dụng kiểm thử gRPC System Configuration Service ===")
	fmt.Println("1. Dịch vụ System Configuration")
	fmt.Println("0. Thoát")
	fmt.Print("Nhập lựa chọn của bạn: ")
}

func printSystemConfigurationMenu() {
	fmt.Println("\n=== Dịch vụ System Configuration ===")
	fmt.Println("1. Tạo cấu hình")
	fmt.Println("2. Lấy cấu hình theo config_key")
	fmt.Println("3. Liệt kê cấu hình")
	fmt.Println("4. Cập nhật cấu hình")
	fmt.Println("5. Xóa cấu hình theo config_key")
	fmt.Println("0. Quay lại menu chính")
	fmt.Print("Nhập lựa chọn của bạn: ")
}

func main() {
	address := serverAddress
	if len(os.Args) > 1 {
		address = os.Args[1]
	}

	fmt.Printf("Đang kết nối tới máy chủ gRPC tại %s...\n", address)
	client, err := NewSystemConfigurationClient(address)
	if err != nil {
		log.Fatalf("Failed to create gRPC client: %v", err)
	}
	defer client.Close()

	fmt.Println("Kết nối thành công!")

	reader := bufio.NewReader(os.Stdin)

	for {
		printMainMenu()
		choice, _ := reader.ReadString('\n')
		choice = cleanInput(choice)

		switch choice {
		case "1":
			// Dịch vụ System Configuration
			for {
				printSystemConfigurationMenu()
				subChoice, _ := reader.ReadString('\n')
				subChoice = cleanInput(subChoice)

				switch subChoice {
				case "1":
					client.TestCreateSystemConfiguration()
				case "2":
					client.TestGetSystemConfiguration()
				case "3":
					client.TestListSystemConfiguration()
				case "4":
					client.TestUpdateSystemConfiguration()
				case "5":
					client.TestDeleteSystemConfiguration()
				case "0":
				default:
					fmt.Println("Invalid choice. Please try again.")
					continue
				}
				if subChoice == "0" {
					break
				}
			}
		case "0":
			fmt.Println("Tạm biệt!")
			return
		default:
			fmt.Println("Lựa chọn không hợp lệ. Vui lòng thử lại.")
		}
	}
}
