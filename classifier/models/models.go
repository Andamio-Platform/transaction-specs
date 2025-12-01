package models

// StringArray is a custom type for storing string arrays in PostgreSQL
type StringArray []string

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
	TxHash       string `gorm:"primaryKey" json:"txHash"`
	Alias        string `json:"alias"`
	CourseID     string `json:"courseId"`
	AssignmentID string `json:"assignmentId"`
	Content      string `gorm:"type:text" json:"content"`
}

// StudentCourseCredentialClaim - /student/course/credential/claim
type StudentCourseCredentialClaim struct {
	TxHash      string      `gorm:"primaryKey" json:"txHash"`
	Alias       string      `json:"alias"`
	CourseID    string      `json:"courseId"`
	Credentials StringArray `gorm:"type:jsonb" json:"credentials"`
}

// TeacherCourseModulesManage - /teacher/course/modules/manage
type TeacherCourseModulesManage struct {
	TxHash   string  `gorm:"primaryKey" json:"txHash"`
	Alias    string  `json:"alias"`
	CourseID string  `json:"courseId"`
	Modules  Modules `gorm:"type:jsonb" json:"modules"`
}

type Modules struct {
	Create []ModulesCreated `json:"create"`
	Update []ModulesUpdated `json:"update"`
	Delete []ModulesDeleted `json:"delete"`
}

type ModulesCreated struct {
	AssignmentID string       `json:"assignmentId"`
	Module       ModuleCreate `json:"module"`
}

type ModulesUpdated struct {
	AssignmentID string       `json:"assignmentId"`
	Module       ModuleUpdate `json:"module"`
}

type ModulesDeleted struct {
	AssignmentID string `json:"assignmentId"`
}

type ModuleCreate struct {
	SLTs          StringArray `json:"slts"`
	Prerequisites StringArray `json:"prerequisites"`
}

type ModuleUpdate struct {
	Prerequisites StringArray `json:"prerequisites"`
}

// TeacherCourseAssignmentsAssess - /teacher/course/assignments/assess
type TeacherCourseAssignmentsAssess struct {
	TxHash       string       `gorm:"primaryKey" json:"txHash"`
	Alias        string       `json:"alias"`
	CourseID     string       `json:"courseId"`
	AssignmentID string       `json:"assignmentId"`
	Assessments  []Assessment `gorm:"type:jsonb" json:"assessments"`
}

type Assessment struct {
	StudentAlias string   `json:"studentAlias"`
	Assessment   Decision `json:"assessment"`
}

type Decision string

const (
	Accept Decision = "accept"
	Refuse Decision = "refuse"
)

// UserAccessTokenMint - /user/access-token/mint
type UserAccessTokenMint struct {
	TxHash string `gorm:"primaryKey" json:"txHash"`
	Alias  string `json:"alias"`
}
