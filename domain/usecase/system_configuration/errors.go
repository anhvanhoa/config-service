package system_configuration

import (
	"github.com/anhvanhoa/service-core/domain/oops"
)

var (
	ErrInvalidID                = oops.New("Id không hợp lệ")
	ErrInvalidConfigKey         = oops.New("Config key không hợp lệ")
	ErrInvalidDataType          = oops.New("Data type không hợp lệ")
	ErrInvalidCategory          = oops.New("Category không hợp lệ")
	ErrInvalidValue             = oops.New("Giá trị không hợp lệ")
	ErrConfigNotFound           = oops.New("Configuration không tồn tại")
	ErrConfigKeyAlreadyExists   = oops.New("Config key đã tồn tại")
	ErrConfigNotEditable        = oops.New("Configuration không thể chỉnh sửa")
	ErrCannotDeleteSystemConfig = oops.New("Không thể xóa Configuration hệ thống")
	ErrUnauthorized             = oops.New("Truy cập không được phép")
	ErrForbidden                = oops.New("Thao tác không được phép")
	ErrInternalError            = oops.New("Lỗi server")
	ErrTimeout                  = oops.New("Thời gian chờ")
)
