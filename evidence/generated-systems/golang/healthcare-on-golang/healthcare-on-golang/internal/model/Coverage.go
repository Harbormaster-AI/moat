package model

import (
    "time"
    "gorm.io/gorm"
)

//==============================================================
// Coverage Declaration
//==============================================================
type Coverage struct {
    gorm.Model
     MemberId                                    string
    GroupNumber                                    string
    EffectiveDate                                                            time.Time
    EndDate                                                            time.Time
    PatientId         *uint
    Patient           *Patient `gorm:"foreignKey:PatientId"`
    PlanId         *uint
    Plan           *InsurancePlan `gorm:"foreignKey:PlanId"`
     Claims           []Claim `gorm:"foreignKey:ClaimsFromCoverageId"`
     Authorizations           []Authorization `gorm:"foreignKey:AuthorizationsFromCoverageId"`
    CoverageType                      CoverageType

// parent associations as their child

}

