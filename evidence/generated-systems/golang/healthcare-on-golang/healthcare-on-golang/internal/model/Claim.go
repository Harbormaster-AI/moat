package model

import (
    "gorm.io/gorm"
)

//==============================================================
// Claim Declaration
//==============================================================
type Claim struct {
    gorm.Model
     ClaimNumber                                    string
    TotalAmount                                                            string
    PatientId         *uint
    Patient           *Patient `gorm:"foreignKey:PatientId"`
    CoverageId         *uint
    Coverage           *Coverage `gorm:"foreignKey:CoverageId"`
    EncounterId         *uint
    Encounter           *Encounter `gorm:"foreignKey:EncounterId"`
     Invoices           []Invoice `gorm:"foreignKey:InvoicesFromClaimId"`
    PayerId         *uint
    Payer           *InsurancePayer `gorm:"foreignKey:PayerId"`
    Status                      ClaimStatus

// parent associations as their child

}

