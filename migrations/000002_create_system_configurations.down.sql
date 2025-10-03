-- Xóa trigger trước khi xóa bảng
DROP TRIGGER IF EXISTS trigger_system_configurations_updated_at ON system_configurations;

-- Xóa bảng system_configurations
DROP TABLE IF EXISTS system_configurations;
