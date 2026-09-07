package model

import (
    "gorm.io/gorm"
)

//==============================================================
// FinancialInstitution Declaration
//==============================================================
type FinancialInstitution struct {
    gorm.Model
     Name                                    string
    LegalName                                    string
    CountryOfIncorporation                                    string
    Bic                                                            string
    Website                                    string
     Branches           []Branch `gorm:"foreignKey:BranchesFromFinancialInstitutionId"`
     Customers           []Customer `gorm:"foreignKey:CustomersFromFinancialInstitutionId"`
     ProductOfferings           []ProductOffering `gorm:"foreignKey:ProductOfferingsFromFinancialInstitutionId"`
     PaymentProcessors           []PaymentProcessor `gorm:"foreignKey:PaymentProcessorsFromFinancialInstitutionId"`
     CompliancePolicies           []CompliancePolicy `gorm:"foreignKey:CompliancePoliciesFromFinancialInstitutionId"`

// parent associations as their child

}

