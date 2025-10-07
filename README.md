# Config Service

Microservice quản lý cấu hình hệ thống cho các ứng dụng IoT trong nông nghiệp, được xây dựng bằng Go và tuân theo nguyên tắc Clean Architecture. Service này cung cấp khả năng quản lý, lưu trữ và truy xuất các cấu hình hệ thống một cách linh hoạt và an toàn, hỗ trợ nhiều loại dữ liệu và validation rules.

## 🏗️ Kiến trúc

Dự án này tuân theo **Clean Architecture** với sự phân tách rõ ràng các mối quan tâm:

```
├── domain/           # Tầng logic nghiệp vụ
│   ├── entity/       # Các thực thể nghiệp vụ cốt lõi
│   │   └── system_configuration.go # Entity cấu hình hệ thống
│   ├── repository/   # Giao diện truy cập dữ liệu
│   │   └── system_configuration_repository.go
│   └── usecase/      # Các trường hợp sử dụng nghiệp vụ
│       └── system_configuration/ # Use cases cấu hình hệ thống
├── infrastructure/   # Các mối quan tâm bên ngoài
│   ├── grpc_service/ # Triển khai API gRPC
│   │   ├── system_configuration/ # gRPC handlers cấu hình hệ thống
│   │   └── server.go            # Thiết lập gRPC server
│   └── repo/         # Triển khai repository cơ sở dữ liệu
├── bootstrap/        # Khởi tạo ứng dụng
└── cmd/             # Điểm vào ứng dụng
```

## 🚀 Tính năng

### Quản lý Cấu hình Hệ thống
- ✅ Tạo, đọc, cập nhật, xóa cấu hình hệ thống
- ✅ Liệt kê cấu hình với bộ lọc (category, data type, editable status)
- ✅ Hỗ trợ nhiều loại dữ liệu: string, number, boolean, JSON, array
- ✅ Phân loại cấu hình theo category: irrigation, fertilization, alerts, sensors, reports
- ✅ Quản lý quyền chỉnh sửa (is_editable) và cấu hình hệ thống (is_system_config)
- ✅ Validation rules linh hoạt cho từng cấu hình
- ✅ Theo dõi người tạo và cập nhật cấu hình
- ✅ Timestamp tự động cho created_at và updated_at
- ✅ Hỗ trợ cấu hình JSON phức tạp với nested objects
- ✅ API gRPC với đầy đủ CRUD operations

## 🛠️ Công nghệ sử dụng

- **Ngôn ngữ**: Go 1.25.0
- **Cơ sở dữ liệu**: PostgreSQL
- **API**: gRPC
- **Kiến trúc**: Clean Architecture
- **Dependencies**:
  - `github.com/go-pg/pg/v10` - PostgreSQL ORM
  - `google.golang.org/grpc` - Framework gRPC
  - `github.com/spf13/viper` - Quản lý cấu hình
  - `go.uber.org/zap` - Logging có cấu trúc
  - `github.com/anhvanhoa/service-core` - Core service utilities
  - `github.com/anhvanhoa/sf-proto` - Protocol buffer definitions

## 📋 Yêu cầu hệ thống

- Go 1.25.0 trở lên
- PostgreSQL 12 trở lên
- [golang-migrate](https://github.com/golang-migrate/migrate) để quản lý migration cơ sở dữ liệu

## 🚀 Hướng dẫn nhanh

### 1. Clone repository
```bash
git clone <repository-url>
cd config-service
```

### 2. Cài đặt dependencies
```bash
go mod download
```

### 3. Thiết lập cơ sở dữ liệu
```bash
# Tạo cơ sở dữ liệu
make create-db

# Chạy migrations
make up
```

### 4. Cấu hình ứng dụng
Sao chép và chỉnh sửa file cấu hình:
```bash
cp dev.config.yml config.yml
```

Cập nhật chuỗi kết nối cơ sở dữ liệu trong `config.yml`:
```yaml
node_env: "development"
url_db: "postgres://pg:123456@localhost:5432/config-service_db?sslmode=disable"
name_service: "MonitoringService"
port_grpc: 50057
host_grpc: "localhost"
interval_check: "20s"
timeout_check: "15s"
```

### 5. Chạy ứng dụng
```bash
# Build và chạy service chính
make run

# Hoặc chạy client để test
make client
```

## 🗄️ Quản lý Cơ sở dữ liệu

Dự án sử dụng `golang-migrate` để quản lý schema cơ sở dữ liệu:

```bash
# Chạy tất cả migrations đang chờ
make up

# Rollback migration cuối cùng
make down

# Reset cơ sở dữ liệu hoàn toàn
make reset

# Tạo migration mới
make create name=migration_name

# Force migration đến phiên bản cụ thể
make force version=1
```

## 🌱 Dữ liệu mẫu

Dự án bao gồm dữ liệu mẫu để phát triển và kiểm thử:

```bash
# Chèn dữ liệu mẫu vào cơ sở dữ liệu
make seed

# Reset cơ sở dữ liệu và chèn dữ liệu mẫu
make seed-reset

# Chèn dữ liệu mẫu vào cơ sở dữ liệu Docker
make docker-seed
```

### Dữ liệu mẫu bao gồm:

**30 cấu hình hệ thống với các category đa dạng:**

**Cấu hình Tưới nước (Irrigation):**
- `irrigation_interval_hours`: Khoảng thời gian tưới (6 giờ)
- `irrigation_duration_minutes`: Thời gian tưới mỗi lần (15 phút)
- `irrigation_auto_mode`: Chế độ tưới tự động (true)
- `irrigation_soil_moisture_threshold`: Ngưỡng độ ẩm đất (30%)

**Cấu hình Bón phân (Fertilization):**
- `fertilization_interval_days`: Khoảng thời gian bón phân (7 ngày)
- `fertilization_amount_ml`: Lượng phân bón (50ml)
- `fertilization_auto_mode`: Chế độ bón phân tự động (false)
- `fertilization_nutrient_ratio`: Tỷ lệ dinh dưỡng NPK

**Cấu hình Cảnh báo (Alerts):**
- `alert_temperature_high/low`: Ngưỡng nhiệt độ (35°C/10°C)
- `alert_humidity_high/low`: Ngưỡng độ ẩm (80%/20%)
- `alert_soil_moisture_low`: Ngưỡng độ ẩm đất (15%)
- `alert_enabled`: Bật/tắt hệ thống cảnh báo
- `alert_notification_methods`: Phương thức thông báo

**Cấu hình Cảm biến (Sensors):**
- `sensor_reading_interval_seconds`: Tần suất đọc dữ liệu (300s)
- `sensor_data_retention_days`: Thời gian lưu trữ (30 ngày)
- `sensor_calibration_enabled`: Hiệu chuẩn cảm biến
- `sensor_temperature/humidity_offset`: Độ lệch hiệu chuẩn

**Cấu hình Báo cáo (Reports):**
- `report_generation_interval`: Tần suất tạo báo cáo (daily)
- `report_include_charts`: Bao gồm biểu đồ (true)
- `report_email_recipients`: Danh sách email nhận báo cáo
- `report_export_formats`: Định dạng xuất báo cáo

## 📁 Cấu trúc Dự án

```
config-service/
├── bootstrap/                 # Khởi tạo ứng dụng
│   ├── app.go               # Khởi tạo app
│   └── env.go               # Cấu hình môi trường
├── cmd/                     # Điểm vào ứng dụng
│   ├── main.go             # Điểm vào service chính
│   └── client/             # gRPC client để test
├── domain/                  # Logic nghiệp vụ (Clean Architecture)
│   ├── entity/             # Các thực thể nghiệp vụ cốt lõi
│   │   ├── system_configuration.go # Entity cấu hình hệ thống
│   │   └── errors.go       # Định nghĩa lỗi
│   ├── repository/         # Giao diện truy cập dữ liệu
│   │   └── system_configuration_repository.go
│   └── usecase/            # Các trường hợp sử dụng nghiệp vụ
│       └── system_configuration/ # Use cases cấu hình hệ thống
│           ├── create_system_configuration_usecase.go
│           ├── get_system_configuration_usecase.go
│           ├── list_system_configuration_usecase.go
│           ├── update_system_configuration_usecase.go
│           ├── delete_system_configuration_usecase.go
│           ├── system_configuration_usecase.go
│           └── errors.go
├── infrastructure/          # Các mối quan tâm bên ngoài
│   ├── grpc_service/       # Triển khai API gRPC
│   │   ├── system_configuration/ # gRPC handlers cấu hình hệ thống
│   │   │   ├── base.go
│   │   │   ├── create.go
│   │   │   ├── get.go
│   │   │   ├── list.go
│   │   │   ├── update.go
│   │   │   └── delete.go
│   │   └── server.go             # Thiết lập gRPC server
│   └── repo/               # Triển khai cơ sở dữ liệu
│       ├── system_configuration_repository.go
│       └── init.go
├── migrations/              # Database migrations
│   ├── 000000_common.up.sql
│   ├── 000002_create_system_configurations.up.sql
│   └── seed/                     # Dữ liệu mẫu
│       └── 000001_seed_system_configurations.up.sql
├── script/seed/             # Script chèn dữ liệu mẫu
├── doc/                     # Tài liệu
└── logs/                    # Log ứng dụng
```

## 🔧 Các lệnh có sẵn

```bash
# Thao tác cơ sở dữ liệu
make up              # Chạy migrations
make down            # Rollback migration
make reset           # Reset cơ sở dữ liệu
make create-db       # Tạo cơ sở dữ liệu
make drop-db         # Xóa cơ sở dữ liệu

# Ứng dụng
make build           # Build ứng dụng
make run             # Chạy service chính
make client          # Chạy client test
make test            # Chạy tests

# Trợ giúp
make help            # Hiển thị tất cả lệnh có sẵn
```

## 📊 Mô hình Dữ liệu

### Cấu hình Hệ thống (System Configuration)
- **ID**: Định danh duy nhất (UUID)
- **ConfigKey**: Khóa cấu hình duy nhất (ví dụ: irrigation_interval_hours)
- **ConfigValue**: Giá trị cấu hình (JSON, có thể là string, number, boolean, object, array)
- **DataType**: Kiểu dữ liệu (string, number, boolean, json, array)
- **Category**: Nhóm cấu hình (irrigation, fertilization, alerts, sensors, reports)
- **Description**: Mô tả chi tiết về cấu hình
- **IsSystemConfig**: Có phải cấu hình hệ thống cốt lõi hay không
- **IsEditable**: Có thể chỉnh sửa hay không
- **ValidationRules**: Quy tắc validate (min, max, regex, enum, etc.)
- **CreatedBy**: Người tạo cấu hình
- **UpdatedBy**: Người cập nhật cấu hình gần nhất
- **CreatedAt**: Thời điểm tạo
- **UpdatedAt**: Thời điểm cập nhật gần nhất

## 🔌 API Endpoints

Service cung cấp các endpoint gRPC:

### System Configuration Service
- `CreateSystemConfiguration` - Tạo cấu hình hệ thống mới
- `GetSystemConfiguration` - Lấy thông tin cấu hình theo ID
- `UpdateSystemConfiguration` - Cập nhật thông tin cấu hình
- `DeleteSystemConfiguration` - Xóa cấu hình hệ thống
- `ListSystemConfigurations` - Liệt kê cấu hình với bộ lọc
- `GetByCategory` - Lấy cấu hình theo category
- `GetByDataType` - Lấy cấu hình theo kiểu dữ liệu
- `GetByKey` - Lấy cấu hình theo config key
- `GetEditableConfigs` - Lấy danh sách cấu hình có thể chỉnh sửa
- `GetSystemConfigs` - Lấy danh sách cấu hình hệ thống
- `ValidateConfiguration` - Validate cấu hình theo rules
- `BulkUpdateConfigurations` - Cập nhật nhiều cấu hình cùng lúc

## 🧪 Testing

Chạy client test để tương tác với service:

```bash
make client
```

Điều này sẽ khởi động một client tương tác nơi bạn có thể test tất cả các endpoint gRPC.

## 📝 Cấu hình

Ứng dụng sử dụng Viper để quản lý cấu hình. Các tùy chọn cấu hình chính:

- `node_env`: Môi trường (development, production)
- `url_db`: Chuỗi kết nối PostgreSQL
- `name_service`: Tên service cho discovery
- `port_grpc`: Cổng gRPC server
- `host_grpc`: Host gRPC server
- `interval_check`: Khoảng thời gian kiểm tra sức khỏe
- `timeout_check`: Timeout kiểm tra sức khỏe

## 🚀 Triển khai

1. **Build ứng dụng**:
   ```bash
   make build
   ```

2. **Thiết lập cơ sở dữ liệu production**:
   ```bash
   make create-db
   make up
   ```

3. **Chạy service**:
   ```bash
   ./bin/app
   ```

## 🤝 Đóng góp

1. Fork repository
2. Tạo feature branch
3. Thực hiện thay đổi
4. Thêm tests nếu cần thiết
5. Submit pull request

## 📄 Giấy phép

Dự án này được cấp phép theo MIT License.

## 🆘 Hỗ trợ

Để được hỗ trợ và đặt câu hỏi, vui lòng tạo issue trong repository.

---

**Lưu ý**: Service này được thiết kế để quản lý cấu hình hệ thống cho các ứng dụng IoT trong nông nghiệp, tuân theo các nguyên tắc kiến trúc microservice để có thể mở rộng và bảo trì dễ dàng. Service cung cấp khả năng quản lý cấu hình linh hoạt, hỗ trợ nhiều loại dữ liệu và validation rules để đảm bảo tính nhất quán và an toàn của hệ thống.
