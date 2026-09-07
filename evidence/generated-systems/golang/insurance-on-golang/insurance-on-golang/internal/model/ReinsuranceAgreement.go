package model

import (
    "gorm.io/gorm"
)

//==============================================================
// ReinsuranceAgreement Declaration
//==============================================================
type ReinsuranceAgreement struct {
    gorm.Model
     AgreementNumber                                    string
    EffectivePeriod                                                            string
    Retention                                                            string
    Limit                                                            string
    CessionPercentage                                                            string
    InsurerId         *uint
    Insurer           *Insurer `gorm:"foreignKey:InsurerId"`
     Policies           []Policy `gorm:"foreignKey:PoliciesFromReinsuranceAgreementId"`
    ReinsuranceType                      ReinsuranceType
    TreatyType                      TreatyType

// parent associations as their child

}

