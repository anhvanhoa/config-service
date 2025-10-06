CREATE TABLE system_configurations (
    id VARCHAR(36) PRIMARY KEY DEFAULT (gen_random_uuid()), -- ID duy nhất của cấu hình
    config_key VARCHAR(100) UNIQUE NOT NULL,     -- Khóa cấu hình (tên duy nhất, ví dụ: irrigation_interval, alert_threshold_temp)
    config_value JSON,                           -- Giá trị cấu hình (có thể là số, chuỗi, boolean, object, array…)
    data_type VARCHAR(50),                       -- Kiểu dữ liệu của config_value để dễ validate
    category VARCHAR(100),                       -- Nhóm cấu hình: tưới, bón phân, cảnh báo, cảm biến, báo cáo…
    description TEXT,                            -- Mô tả chi tiết về cấu hình
    is_system_config BOOLEAN DEFAULT FALSE,      -- Cho biết đây có phải là cấu hình hệ thống (cốt lõi, không thay đổi thường xuyên)
    is_editable BOOLEAN NOT NULL DEFAULT TRUE,   -- Cho phép người dùng có thể chỉnh sửa hay không
    validation_rules JSON,                       -- Quy tắc validate (ví dụ: min, max, regex…)
    created_by VARCHAR(36),                      -- Người tạo cấu hình
    updated_by VARCHAR(36),                      -- Người chỉnh sửa cấu hình gần nhất
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, -- Thời điểm tạo
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, -- Thời điểm cập nhật gần nhất
    
    CONSTRAINT chk_data_type CHECK (data_type IN ('string', 'number', 'boolean', 'json', 'array')),
    CONSTRAINT chk_category CHECK (category IN ('irrigation', 'fertilization', 'alerts', 'sensors', 'reports'))
);

-- Tạo index để truy vấn theo nhóm cấu hình
CREATE INDEX idx_system_configurations_category ON system_configurations (category);

-- Tạo trigger để tự động cập nhật updated_at
CREATE TRIGGER trigger_system_configurations_updated_at
    BEFORE UPDATE ON system_configurations
    FOR EACH ROW
    EXECUTE FUNCTION update_updated_at_column();

-- Thêm comment cho bảng và các cột
COMMENT ON TABLE system_configurations IS 'Bảng lưu trữ các cấu hình hệ thống';
COMMENT ON COLUMN system_configurations.id IS 'ID duy nhất của cấu hình';
COMMENT ON COLUMN system_configurations.config_key IS 'Khóa cấu hình (tên duy nhất)';
COMMENT ON COLUMN system_configurations.config_value IS 'Giá trị cấu hình (có thể là số, chuỗi, boolean, object, array)';
COMMENT ON COLUMN system_configurations.data_type IS 'Kiểu dữ liệu của config_value để dễ validate';
COMMENT ON COLUMN system_configurations.category IS 'Nhóm cấu hình: tưới, bón phân, cảnh báo, cảm biến, báo cáo';
COMMENT ON COLUMN system_configurations.description IS 'Mô tả chi tiết về cấu hình';
COMMENT ON COLUMN system_configurations.is_system_config IS 'Cho biết đây có phải là cấu hình hệ thống (cốt lõi, không thay đổi thường xuyên)';
COMMENT ON COLUMN system_configurations.is_editable IS 'Cho phép người dùng có thể chỉnh sửa hay không';
COMMENT ON COLUMN system_configurations.validation_rules IS 'Quy tắc validate (ví dụ: min, max, regex)';
COMMENT ON COLUMN system_configurations.created_by IS 'Người tạo cấu hình';
COMMENT ON COLUMN system_configurations.updated_by IS 'Người chỉnh sửa cấu hình gần nhất';
COMMENT ON COLUMN system_configurations.created_at IS 'Thời điểm tạo';
COMMENT ON COLUMN system_configurations.updated_at IS 'Thời điểm cập nhật gần nhất';
