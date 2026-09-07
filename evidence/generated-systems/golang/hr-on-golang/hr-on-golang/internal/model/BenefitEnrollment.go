package model

import (
    "time"
    "gorm.io/gorm"
)

//==============================================================
// BenefitEnrollment Declaration
//==============================================================
type BenefitEnrollment struct {
    gorm.Model
     EnrollmentId                                    string
    EffectiveFrom                                                            time.Time
    EffectiveTo                                                            time.Time
    BenefitPlanId         *uint
    BenefitPlan           *BenefitPlan `gorm:"foreignKey:BenefitPlanId"`
    EmployeeId         *uint
    Employee           *Employee `gorm:"foreignKey:EmployeeId"`
     Dependents           []Dependent `gorm:"foreignKey:DependentsFromBenefitEnrollmentId"`
    Status                      BenefitEnrollmentStatus
    CoverageLevel                      CoverageLevel

// parent associations as their child

}

