-- Rollback seed data cho bảng system_configurations
-- Xóa tất cả dữ liệu seed đã thêm

DELETE FROM system_configurations WHERE created_by = 'system';
