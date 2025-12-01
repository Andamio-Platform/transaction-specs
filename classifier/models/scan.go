package models

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
)

// Value implements the driver.Valuer interface
func (a StringArray) Value() (driver.Value, error) {
	return json.Marshal(a)
}

// Scan implements the sql.Scanner interface
func (a *StringArray) Scan(value interface{}) error {
	if value == nil {
		*a = []string{}
		return nil
	}
	bytes, ok := value.([]byte)
	if !ok {
		return nil
	}
	return json.Unmarshal(bytes, &a)
}

// For jsonb column storing the Modules struct
func (m Modules) Value() (driver.Value, error) {
	return json.Marshal(m)
}

func (m *Modules) Scan(value interface{}) error {
	if value == nil {
		*m = Modules{}
		return nil
	}
	bytes, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("Modules: expected []byte, got %T", value)
	}
	return json.Unmarshal(bytes, m)
}

// For jsonb column storing AssessmentArray
func (a AssessmentArray) Value() (driver.Value, error) {
	return json.Marshal(a)
}

func (a *AssessmentArray) Scan(value interface{}) error {
	if value == nil {
		*a = []Assessment{}
		return nil
	}
	bytes, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("AssessmentArray: expected []byte, got %T", value)
	}
	return json.Unmarshal(bytes, a)
}
