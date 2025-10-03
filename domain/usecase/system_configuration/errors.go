package system_configuration

import "errors"

// Usecase errors
var (
	// Validation errors
	ErrInvalidID        = errors.New("invalid ID")
	ErrInvalidConfigKey = errors.New("invalid config key")
	ErrInvalidDataType  = errors.New("invalid data type")
	ErrInvalidCategory  = errors.New("invalid category")
	ErrInvalidValue     = errors.New("invalid config value")

	// Business logic errors
	ErrConfigNotFound           = errors.New("configuration not found")
	ErrConfigKeyAlreadyExists   = errors.New("config key already exists")
	ErrConfigNotEditable        = errors.New("configuration is not editable")
	ErrCannotDeleteSystemConfig = errors.New("cannot delete system configuration")

	// Permission errors
	ErrUnauthorized = errors.New("unauthorized access")
	ErrForbidden    = errors.New("forbidden operation")

	// System errors
	ErrInternalError = errors.New("internal server error")
	ErrTimeout       = errors.New("operation timeout")
)

// Error codes
const (
	ErrCodeInvalidID          = "INVALID_ID"
	ErrCodeInvalidConfigKey   = "INVALID_CONFIG_KEY"
	ErrCodeInvalidDataType    = "INVALID_DATA_TYPE"
	ErrCodeInvalidCategory    = "INVALID_CATEGORY"
	ErrCodeInvalidValue       = "INVALID_VALUE"
	ErrCodeConfigNotFound     = "CONFIG_NOT_FOUND"
	ErrCodeConfigKeyExists    = "CONFIG_KEY_EXISTS"
	ErrCodeConfigNotEditable  = "CONFIG_NOT_EDITABLE"
	ErrCodeCannotDeleteSystem = "CANNOT_DELETE_SYSTEM_CONFIG"
	ErrCodeUnauthorized       = "UNAUTHORIZED"
	ErrCodeForbidden          = "FORBIDDEN"
	ErrCodeInternalError      = "INTERNAL_ERROR"
	ErrCodeTimeout            = "TIMEOUT"
)
