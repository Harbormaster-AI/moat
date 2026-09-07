package model

import (
    "gorm.io/gorm"
)

//==============================================================
// InsurancePayer Declaration
//==============================================================
type InsurancePayer struct {
    gorm.Model
     Name                                    string
    Website                                    string
     Plans           []InsurancePlan `gorm:"foreignKey:PlansFromInsurancePayerId"`
     Claims           []Claim `gorm:"foreignKey:ClaimsFromInsurancePayerId"`
    PayerType                      PayerType

// parent associations as their child

}

