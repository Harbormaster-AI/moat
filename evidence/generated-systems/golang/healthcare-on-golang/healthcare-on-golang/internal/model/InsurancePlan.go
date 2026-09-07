package model

import (
    "gorm.io/gorm"
)

//==============================================================
// InsurancePlan Declaration
//==============================================================
type InsurancePlan struct {
    gorm.Model
     Name                                    string
    PlanCode                                    string
    PayerId         *uint
    Payer           *InsurancePayer `gorm:"foreignKey:PayerId"`
     Coverages           []Coverage `gorm:"foreignKey:CoveragesFromInsurancePlanId"`
    PlanType                      InsurancePlanType

// parent associations as their child

}

