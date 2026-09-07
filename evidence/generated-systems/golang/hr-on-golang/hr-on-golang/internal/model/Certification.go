package model

import (
    "time"
    "gorm.io/gorm"
)

//==============================================================
// Certification Declaration
//==============================================================
type Certification struct {
    gorm.Model
     Name                                    string
    Issuer                                    string
    ValidFrom                                                            time.Time
    ValidTo                                                            time.Time
    CredentialId                                    string
    EmployeeId         *uint
    Employee           *Employee `gorm:"foreignKey:EmployeeId"`
    CourseId         *uint
    Course           *TrainingCourse `gorm:"foreignKey:CourseId"`

// parent associations as their child

}

