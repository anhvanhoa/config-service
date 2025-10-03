package entity

import (
	"github.com/anhvanhoa/service-core/domain/oops"
)

var (
	ErrInvalidConfigKey        = oops.New("invalid config key")
	ErrInvalidDataType         = oops.New("invalid data type")
	ErrInvalidCategory         = oops.New("invalid category")
	ErrInvalidValue            = oops.New("invalid config value")
	ErrValidationFailed        = oops.New("validation failed")
	ErrSystemConfigNotEditable = oops.New("system config is not editable")
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
