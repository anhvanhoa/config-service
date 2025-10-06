package entity

import (
	"database/sql/driver"
	"encoding/json"
	"slices"
	"time"
)

type SystemConfiguration struct {
	tableName       struct{} `pg:"system_configurations"`
	ID              string
	ConfigKey       string
	ConfigValue     JSONValue
	DataType        string
	Category        string
	Description     string
	IsSystemConfig  bool
	IsEditable      bool
	ValidationRules JSONValue
	CreatedBy       string
	UpdatedBy       string
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

type JSONValue struct {
	Data any `json:"-"`
}

func (j *JSONValue) Scan(value any) error {
	if value == nil {
		j.Data = nil
		return nil
	}

	bytes, ok := value.([]byte)
	if !ok {
		return nil
	}

	return json.Unmarshal(bytes, &j.Data)
}

func (j JSONValue) Value() (driver.Value, error) {
	if j.Data == nil {
		return nil, nil
	}
	return json.Marshal(j.Data)
}

func (j *JSONValue) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &j.Data)
}

func (j JSONValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(j.Data)
}

func (sc *SystemConfiguration) GetStringValue() string {
	if str, ok := sc.ConfigValue.Data.(string); ok {
		return str
	}
	return ""
}

func (sc *SystemConfiguration) GetNumberValue() float64 {
	switch v := sc.ConfigValue.Data.(type) {
	case float64:
		return v
	case int:
		return float64(v)
	case int64:
		return float64(v)
	case float32:
		return float64(v)
	}
	return 0
}

func (sc *SystemConfiguration) GetBoolValue() bool {
	if b, ok := sc.ConfigValue.Data.(bool); ok {
		return b
	}
	return false
}

func (sc *SystemConfiguration) GetArrayValue() []any {
	if arr, ok := sc.ConfigValue.Data.([]any); ok {
		return arr
	}
	return nil
}

func (sc *SystemConfiguration) GetObjectValue() map[string]any {
	if obj, ok := sc.ConfigValue.Data.(map[string]any); ok {
		return obj
	}
	return nil
}

func (sc *SystemConfiguration) IsValidDataType() bool {
	validTypes := []string{"string", "number", "boolean", "json", "array"}
	return slices.Contains(validTypes, sc.DataType)
}

func (sc *SystemConfiguration) IsValidCategory() bool {
	validCategories := []string{"irrigation", "fertilization", "alerts", "sensors", "reports"}
	return slices.Contains(validCategories, sc.Category)
}

func (sc *SystemConfiguration) Validate() error {
	if sc.ConfigKey == "" {
		return ErrInvalidConfigKey
	}

	if !sc.IsValidDataType() {
		return ErrInvalidDataType
	}

	if !sc.IsValidCategory() {
		return ErrInvalidCategory
	}

	return nil
}

func (sc *SystemConfiguration) SetConfigValue(value any) {
	sc.ConfigValue.Data = value
}

func (sc *SystemConfiguration) TableName() struct{} {
	return sc.tableName
}
