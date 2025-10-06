package entity

import (
	"github.com/anhvanhoa/service-core/domain/oops"
)

var (
	ErrInvalidConfigKey = oops.New("Khóa cấu hình không hợp lệ")
	ErrInvalidDataType  = oops.New("Loại dữ liệu không hợp lệ")
	ErrInvalidCategory  = oops.New("Loại cấu hình không hợp lệ")
)

const (
	DataTypeString  = "string"
	DataTypeNumber  = "number"
	DataTypeBoolean = "boolean"
	DataTypeJSON    = "json"
	DataTypeArray   = "array"
)

const (
	CategoryIrrigation    = "irrigation"
	CategoryFertilization = "fertilization"
	CategoryAlerts        = "alerts"
	CategorySensors       = "sensors"
	CategoryReports       = "reports"
)

func ValidDataTypes() []string {
	return []string{
		DataTypeString,
		DataTypeNumber,
		DataTypeBoolean,
		DataTypeJSON,
		DataTypeArray,
	}
}

func ValidCategories() []string {
	return []string{
		CategoryIrrigation,
		CategoryFertilization,
		CategoryAlerts,
		CategorySensors,
		CategoryReports,
	}
}
