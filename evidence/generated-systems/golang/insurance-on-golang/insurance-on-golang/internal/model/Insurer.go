package model

import (
    "gorm.io/gorm"
)

//==============================================================
// Insurer Declaration
//==============================================================
type Insurer struct {
    gorm.Model
     Name                                    string
    LegalName                                    string
    DomicileCountry                                    string
    NaicNumber                                    string
    Website                                    string
     Products           []InsuranceProduct `gorm:"foreignKey:ProductsFromInsurerId"`
     DistributionPartners           []Distributor `gorm:"foreignKey:DistributionPartnersFromInsurerId"`
     Policies           []Policy `gorm:"foreignKey:PoliciesFromInsurerId"`
     Claims           []Claim `gorm:"foreignKey:ClaimsFromInsurerId"`
     ReinsuranceAgreements           []ReinsuranceAgreement `gorm:"foreignKey:ReinsuranceAgreementsFromInsurerId"`

// parent associations as their child

}

