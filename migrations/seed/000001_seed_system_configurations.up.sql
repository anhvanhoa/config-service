-- Seed data cho bảng system_configurations
-- Các cấu hình cơ bản cho hệ thống

-- Cấu hình tưới nước
INSERT INTO system_configurations (id, config_key, config_value, data_type, category, description, is_system_config, is_editable, validation_rules, created_by, updated_by) VALUES
('550e8400-e29b-41d4-a716-446655440001', 'irrigation_interval_hours', '{"value": 6, "unit": "hours"}', 'json', 'irrigation', 'Khoảng thời gian giữa các lần tưới nước (giờ)', false, true, '{"min": 1, "max": 24, "type": "number"}', 'system', 'system'),
('550e8400-e29b-41d4-a716-446655440002', 'irrigation_duration_minutes', '{"value": 15, "unit": "minutes"}', 'json', 'irrigation', 'Thời gian tưới nước mỗi lần (phút)', false, true, '{"min": 5, "max": 60, "type": "number"}', 'system', 'system'),
('550e8400-e29b-41d4-a716-446655440003', 'irrigation_auto_mode', 'true', 'boolean', 'irrigation', 'Chế độ tưới tự động', false, true, '{"type": "boolean"}', 'system', 'system'),
('550e8400-e29b-41d4-a716-446655440004', 'irrigation_soil_moisture_threshold', '{"value": 30, "unit": "%"}', 'json', 'irrigation', 'Ngưỡng độ ẩm đất để kích hoạt tưới nước', false, true, '{"min": 0, "max": 100, "type": "number"}', 'system', 'system');

-- Cấu hình bón phân
INSERT INTO system_configurations (id, config_key, config_value, data_type, category, description, is_system_config, is_editable, validation_rules, created_by, updated_by) VALUES
('550e8400-e29b-41d4-a716-446655440005', 'fertilization_interval_days', '{"value": 7, "unit": "days"}', 'json', 'fertilization', 'Khoảng thời gian giữa các lần bón phân (ngày)', false, true, '{"min": 1, "max": 30, "type": "number"}', 'system', 'system'),
('550e8400-e29b-41d4-a716-446655440006', 'fertilization_amount_ml', '{"value": 50, "unit": "ml"}', 'json', 'fertilization', 'Lượng phân bón mỗi lần (ml)', false, true, '{"min": 10, "max": 200, "type": "number"}', 'system', 'system'),
('550e8400-e29b-41d4-a716-446655440007', 'fertilization_auto_mode', 'false', 'boolean', 'fertilization', 'Chế độ bón phân tự động', false, true, '{"type": "boolean"}', 'system', 'system'),
('550e8400-e29b-41d4-a716-446655440008', 'fertilization_nutrient_ratio', '{"N": 20, "P": 10, "K": 15}', 'json', 'fertilization', 'Tỷ lệ dinh dưỡng NPK', false, true, '{"type": "object"}', 'system', 'system');

-- Cấu hình cảnh báo
INSERT INTO system_configurations (id, config_key, config_value, data_type, category, description, is_system_config, is_editable, validation_rules, created_by, updated_by) VALUES
('550e8400-e29b-41d4-a716-446655440009', 'alert_temperature_high', '{"value": 35, "unit": "°C"}', 'json', 'alerts', 'Ngưỡng nhiệt độ cao để cảnh báo', false, true, '{"min": 20, "max": 50, "type": "number"}', 'system', 'system'),
('550e8400-e29b-41d4-a716-446655440010', 'alert_temperature_low', '{"value": 10, "unit": "°C"}', 'json', 'alerts', 'Ngưỡng nhiệt độ thấp để cảnh báo', false, true, '{"min": -10, "max": 30, "type": "number"}', 'system', 'system'),
('550e8400-e29b-41d4-a716-446655440011', 'alert_humidity_high', '{"value": 80, "unit": "%"}', 'json', 'alerts', 'Ngưỡng độ ẩm cao để cảnh báo', false, true, '{"min": 0, "max": 100, "type": "number"}', 'system', 'system'),
('550e8400-e29b-41d4-a716-446655440012', 'alert_humidity_low', '{"value": 20, "unit": "%"}', 'json', 'alerts', 'Ngưỡng độ ẩm thấp để cảnh báo', false, true, '{"min": 0, "max": 100, "type": "number"}', 'system', 'system'),
('550e8400-e29b-41d4-a716-446655440013', 'alert_soil_moisture_low', '{"value": 15, "unit": "%"}', 'json', 'alerts', 'Ngưỡng độ ẩm đất thấp để cảnh báo', false, true, '{"min": 0, "max": 100, "type": "number"}', 'system', 'system'),
('550e8400-e29b-41d4-a716-446655440014', 'alert_enabled', 'true', 'boolean', 'alerts', 'Bật/tắt hệ thống cảnh báo', false, true, '{"type": "boolean"}', 'system', 'system'),
('550e8400-e29b-41d4-a716-446655440015', 'alert_notification_methods', '["email", "sms", "push"]', 'array', 'alerts', 'Phương thức gửi thông báo cảnh báo', false, true, '{"type": "array", "items": {"type": "string", "enum": ["email", "sms", "push", "webhook"]}}', 'system', 'system');

-- Cấu hình cảm biến
INSERT INTO system_configurations (id, config_key, config_value, data_type, category, description, is_system_config, is_editable, validation_rules, created_by, updated_by) VALUES
('550e8400-e29b-41d4-a716-446655440016', 'sensor_reading_interval_seconds', '{"value": 300, "unit": "seconds"}', 'json', 'sensors', 'Khoảng thời gian đọc dữ liệu từ cảm biến (giây)', true, false, '{"min": 60, "max": 3600, "type": "number"}', 'system', 'system'),
('550e8400-e29b-41d4-a716-446655440017', 'sensor_data_retention_days', '{"value": 30, "unit": "days"}', 'json', 'sensors', 'Số ngày lưu trữ dữ liệu cảm biến', true, false, '{"min": 7, "max": 365, "type": "number"}', 'system', 'system'),
('550e8400-e29b-41d4-a716-446655440018', 'sensor_calibration_enabled', 'true', 'boolean', 'sensors', 'Bật/tắt chế độ hiệu chuẩn cảm biến', false, true, '{"type": "boolean"}', 'system', 'system'),
('550e8400-e29b-41d4-a716-446655440019', 'sensor_temperature_offset', '{"value": 0, "unit": "°C"}', 'json', 'sensors', 'Độ lệch hiệu chuẩn cảm biến nhiệt độ', false, true, '{"min": -5, "max": 5, "type": "number"}', 'system', 'system'),
('550e8400-e29b-41d4-a716-446655440020', 'sensor_humidity_offset', '{"value": 0, "unit": "%"}', 'json', 'sensors', 'Độ lệch hiệu chuẩn cảm biến độ ẩm', false, true, '{"min": -10, "max": 10, "type": "number"}', 'system', 'system');

-- Cấu hình báo cáo
INSERT INTO system_configurations (id, config_key, config_value, data_type, category, description, is_system_config, is_editable, validation_rules, created_by, updated_by) VALUES
('550e8400-e29b-41d4-a716-446655440021', 'report_generation_interval', '{"value": "daily", "unit": "frequency"}', 'json', 'reports', 'Tần suất tạo báo cáo tự động', false, true, '{"type": "string", "enum": ["hourly", "daily", "weekly", "monthly"]}', 'system', 'system'),
('550e8400-e29b-41d4-a716-446655440022', 'report_include_charts', 'true', 'boolean', 'reports', 'Bao gồm biểu đồ trong báo cáo', false, true, '{"type": "boolean"}', 'system', 'system'),
('550e8400-e29b-41d4-a716-446655440023', 'report_email_recipients', '["admin@farm.com", "manager@farm.com"]', 'array', 'reports', 'Danh sách email nhận báo cáo', false, true, '{"type": "array", "items": {"type": "string", "format": "email"}}', 'system', 'system'),
('550e8400-e29b-41d4-a716-446655440024', 'report_data_points_limit', '{"value": 1000, "unit": "points"}', 'json', 'reports', 'Giới hạn số điểm dữ liệu trong báo cáo', true, false, '{"min": 100, "max": 10000, "type": "number"}', 'system', 'system'),
('550e8400-e29b-41d4-a716-446655440025', 'report_export_formats', '["pdf", "excel", "csv"]', 'array', 'reports', 'Các định dạng xuất báo cáo', false, true, '{"type": "array", "items": {"type": "string", "enum": ["pdf", "excel", "csv", "json"]}}', 'system', 'system');

-- Cấu hình hệ thống cốt lõi (sử dụng category 'sensors' cho các cấu hình hệ thống)
INSERT INTO system_configurations (id, config_key, config_value, data_type, category, description, is_system_config, is_editable, validation_rules, created_by, updated_by) VALUES
('550e8400-e29b-41d4-a716-446655440026', 'system_timezone', '{"value": "Asia/Ho_Chi_Minh", "unit": "timezone"}', 'json', 'sensors', 'Múi giờ hệ thống', true, false, '{"type": "string"}', 'system', 'system'),
('550e8400-e29b-41d4-a716-446655440027', 'system_language', '{"value": "vi", "unit": "locale"}', 'json', 'sensors', 'Ngôn ngữ hệ thống', true, false, '{"type": "string", "enum": ["vi", "en"]}', 'system', 'system'),
('550e8400-e29b-41d4-a716-446655440028', 'system_maintenance_mode', 'false', 'boolean', 'sensors', 'Chế độ bảo trì hệ thống', true, false, '{"type": "boolean"}', 'system', 'system'),
('550e8400-e29b-41d4-a716-446655440029', 'system_log_level', '{"value": "info", "unit": "level"}', 'json', 'sensors', 'Mức độ ghi log hệ thống', true, false, '{"type": "string", "enum": ["debug", "info", "warn", "error"]}', 'system', 'system'),
('550e8400-e29b-41d4-a716-446655440030', 'system_api_rate_limit', '{"value": 1000, "unit": "requests/hour"}', 'json', 'sensors', 'Giới hạn tốc độ API', true, false, '{"min": 100, "max": 10000, "type": "number"}', 'system', 'system');
