package model

import (
    "time"
    "gorm.io/gorm"
)

//==============================================================
// Customer Declaration
//==============================================================
type Customer struct {
    gorm.Model
     FirstName                                    string
    LastName                                    string
    OrganizationName                                    string
    TaxId                                    string
    DateOfBirth                                                            time.Time
    PrimaryAddress                                                            string
     Applications           []Application `gorm:"foreignKey:ApplicationsFromCustomerId"`
     Policies           []Policy `gorm:"foreignKey:PoliciesFromCustomerId"`
     Claims           []Claim `gorm:"foreignKey:ClaimsFromCustomerId"`
     Agents           []Agent `gorm:"foreignKey:AgentsFromCustomerId"`
     Beneficiaries           []Beneficiary `gorm:"foreignKey:BeneficiariesFromCustomerId"`
    CustomerType                      CustomerType

// parent associations as their child

}

