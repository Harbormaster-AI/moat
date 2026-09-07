package model

import (
    "time"
    "gorm.io/gorm"
)

//==============================================================
// Dependent Declaration
//==============================================================
type Dependent struct {
    gorm.Model
     FirstName                                    string
    LastName                                    string
    BirthDate                                                            time.Time
    BenefitEnrollmentId         *uint
    BenefitEnrollment           *BenefitEnrollment `gorm:"foreignKey:BenefitEnrollmentId"`
    EmployeeId         *uint
    Employee           *Employee `gorm:"foreignKey:EmployeeId"`
    Relationship                      DependentRelationship

// parent associations as their child

}

