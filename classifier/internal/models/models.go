package models

import (
	"database/sql/driver"
	"encoding/json"
	"time"
)

// StringArray is a custom type for storing string arrays in PostgreSQL
type StringArray []string

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

// AdminCourseCreate - /admin/course/create
type AdminCourseCreate struct {
	TxHash              string      `gorm:"primaryKey" json:"txHash"`
	CourseID            string      `json:"courseId"`
	Admin               string      `json:"admin"`
	Teachers            StringArray `gorm:"type:jsonb" json:"teachers"`
	CourseAddress       string      `json:"courseAddress"`
	CourseStatePolicyId string      `json:"courseStatePolicyId"`
}

// AdminCourseTeachersUpdate - /admin/course/teachers/update
type AdminCourseTeachersUpdate struct {
	TxHash   string `gorm:"primaryKey" json:"txHash"`
	CourseID string `json:"courseId"`
	Teachers string `gorm:"type:jsonb" json:"teachers"` // Stores add/remove operations as JSON
}

// StudentCourseEnroll - /student/course/enroll
type StudentCourseEnroll struct {
	TxHash   string `gorm:"primaryKey" json:"txHash"`
	Alias    string `json:"alias"`
	CourseID string `json:"courseId"`
}

// StudentCourseAssignmentSubmit - /student/course/assignment/submit
type StudentCourseAssignmentSubmit struct {
	TxHash       string `gorm:"primaryKey" json:"txHash"`
	Alias        string `json:"alias"`
	CourseID     string `json:"courseId"`
	AssignmentID string `json:"assignmentId"`
	Content      string `gorm:"type:text" json:"content"`
}

// StudentCourseAssignmentUpdate - /student/course/assignment/update
type StudentCourseAssignmentUpdate struct {
	TxHash       string    `gorm:"primaryKey" json:"txHash"`
	Alias        string    `json:"alias"`
	CourseID     string    `json:"courseId"`
	AssignmentID string    `json:"assignmentId"`
	Content      string    `gorm:"type:text" json:"content"`
	UpdatedAt    time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

// StudentCourseCredentialClaim - /student/course/credential/claim
type StudentCourseCredentialClaim struct {
	TxHash   string `gorm:"primaryKey" json:"txHash"`
	Alias    string `json:"alias"`
	CourseID string `json:"courseId"`
}

// TeacherCourseModulesManage - /teacher/course/modules/manage
type TeacherCourseModulesManage struct {
	TxHash   string `gorm:"primaryKey" json:"txHash"`
	Alias    string `json:"alias"`
	CourseID string `json:"courseId"`
	Modules  string `gorm:"type:jsonb" json:"modules"` // Stores create/update/delete operations as JSON
}

// TeacherCourseAssignmentsAssess - /teacher/course/assignments/assess
type TeacherCourseAssignmentsAssess struct {
	TxHash       string `gorm:"primaryKey" json:"txHash"`
	Alias        string `json:"alias"`
	CourseID     string `json:"courseId"`
	AssignmentID string `json:"assignmentId"`
	Assessments  string `gorm:"type:jsonb" json:"assessments"` // Stores assessment array as JSON
}

// UserAccessTokenMint - /user/access-token/mint
type UserAccessTokenMint struct {
	TxHash string `gorm:"primaryKey" json:"txHash"`
	Alias  string `json:"alias"`
}
