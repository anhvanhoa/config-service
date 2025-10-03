CREATE TABLE system_configurations (
    id VARCHAR(36) PRIMARY KEY DEFAULT (UUID()), -- ID duy nhất của cấu hình
    config_key VARCHAR(100) UNIQUE NOT NULL,     -- Khóa cấu hình (tên duy nhất, ví dụ: irrigation_interval, alert_threshold_temp)
    config_value JSON,                           -- Giá trị cấu hình (có thể là số, chuỗi, boolean, object, array…)
    data_type VARCHAR(50) COMMENT 'string, number, boolean, json, array', 
        -- Kiểu dữ liệu của config_value để dễ validate
    category VARCHAR(100) COMMENT 'irrigation, fertilization, alerts, sensors, reports', 
        -- Nhóm cấu hình: tưới, bón phân, cảnh báo, cảm biến, báo cáo…
    description TEXT,                            -- Mô tả chi tiết về cấu hình
    is_system_config BOOLEAN DEFAULT FALSE,      -- Cho biết đây có phải là cấu hình hệ thống (cốt lõi, không thay đổi thường xuyên)
    is_editable BOOLEAN DEFAULT TRUE,            -- Cho phép người dùng có thể chỉnh sửa hay không
    validation_rules JSON,                       -- Quy tắc validate (ví dụ: min, max, regex…)
    created_by VARCHAR(36),                      -- Người tạo cấu hình
    updated_by VARCHAR(36),                      -- Người chỉnh sửa cấu hình gần nhất
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, -- Thời điểm tạo
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP, -- Thời điểm cập nhật gần nhất
    
    INDEX idx_system_configurations_category (category), -- Index để truy vấn theo nhóm cấu hình
    UNIQUE KEY uk_config_key (config_key)         -- Đảm bảo mỗi config_key là duy nhất
);
