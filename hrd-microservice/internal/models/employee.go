package models

import (
	"encoding/json"
	"time"

	"github.com/IKHINtech/erp_microservice_go_grpc/hrd-microservice/internal/utils"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type (
	EmploymentType   string
	EmploymentStatus string
	Gender           string
	MaritalStatus    string
)

const (
	FullTime EmploymentType = "full_time"
	PartTime EmploymentType = "part_time"
	Contract EmploymentType = "contract"
	Intern   EmploymentType = "intern"

	Active     EmploymentStatus = "active"
	OnLeave    EmploymentStatus = "on_leave"
	Terminated EmploymentStatus = "terminated"

	Male      Gender = "male"
	Female    Gender = "female"
	NonBinary Gender = "non_binary"

	Single   MaritalStatus = "single"
	Married  MaritalStatus = "married"
	Divorced MaritalStatus = "divorced"
	Widowed  MaritalStatus = "widowed"
)

type Employee struct {
	ID                uuid.UUID        `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	EmployeeID        string           `gorm:"type:varchar(20);uniqueIndex;not null"`
	FirstName         string           `gorm:"type:varchar(50);not null"`
	LastName          string           `gorm:"type:varchar(50);not null"`
	Email             string           `gorm:"type:varchar(100);uniqueIndex;not null"`
	Phone             string           `gorm:"type:varchar(20)"`
	DateOfBirth       time.Time        `gorm:"type:date"`
	Gender            Gender           `gorm:"type:varchar(20);check:gender IN ('male','female','non_binary')"`
	Nationality       string           `gorm:"type:varchar(50)"`
	MaritalStatus     MaritalStatus    `gorm:"type:varchar(20);check:marital_status IN ('single','married','divorced','widowed')"`
	EmergencyContact  json.RawMessage  `gorm:"type:jsonb"` // {name, relationship, phone}
	Address           json.RawMessage  `gorm:"type:jsonb"` // {street, city, postal_code, country}
	JoinDate          time.Time        `gorm:"type:date;not null"`
	DepartmentID      uuid.UUID        `gorm:"type:uuid;index"`
	PositionID        uuid.UUID        `gorm:"type:uuid;index"`
	EmploymentType    EmploymentType   `gorm:"type:varchar(20);check:employment_type IN ('full_time','part_time','contract','intern')"`
	EmploymentStatus  EmploymentStatus `gorm:"type:varchar(20);check:employment_status IN ('active','on_leave','terminated')"`
	ManagerID         *uuid.UUID       `gorm:"type:uuid;index;null"`
	SalaryCurrency    string           `gorm:"type:varchar(3)"`
	BaseSalary        float64          `gorm:"type:numeric(12,2)"`
	BankAccount       json.RawMessage  `gorm:"type:jsonb"` // {bank_name, account_number, branch}
	GovernmentID      string           `gorm:"type:varchar(50)"`
	TaxID             string           `gorm:"type:varchar(50)"`
	SocialSecurityNo  string           `gorm:"type:varchar(50)"`
	ProfilePictureURL string           `gorm:"type:varchar(255)"`
	CreatedAt         time.Time        `gorm:"autoCreateTime"`
	UpdatedAt         time.Time        `gorm:"autoUpdateTime"`

	// Relations
	Department Department `gorm:"foreignKey:DepartmentID"`
	Position   Position   `gorm:"foreignKey:PositionID"`
	Manager    *Employee  `gorm:"foreignKey:ManagerID"`
}

// Encrypt sensitive fields before save
func (e *Employee) BeforeSave(tx *gorm.DB) error {
	e.GovernmentID = utils.Encrypt(e.GovernmentID)
	e.TaxID = utils.Encrypt(e.TaxID)
	e.SocialSecurityNo = utils.Encrypt(e.SocialSecurityNo)
	return nil
}

func (e *Employee) BeforeUpdate(tx *gorm.DB) error {
	e.GovernmentID = utils.Encrypt(e.GovernmentID)
	e.TaxID = utils.Encrypt(e.TaxID)
	e.SocialSecurityNo = utils.Encrypt(e.SocialSecurityNo)
	return nil
}

func (e *Employee) AfterFind(tx *gorm.DB) (err error) {
	e.GovernmentID, err = utils.Decrypt(e.GovernmentID)
	if err != nil {
		return err
	}

	e.TaxID, err = utils.Decrypt(e.TaxID)
	if err != nil {
		return err
	}

	e.SocialSecurityNo, err = utils.Decrypt(e.SocialSecurityNo)
	if err != nil {
		return err
	}
	return nil
}
