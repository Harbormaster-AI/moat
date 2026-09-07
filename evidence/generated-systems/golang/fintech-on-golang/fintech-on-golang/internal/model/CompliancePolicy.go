package model

import (
    "gorm.io/gorm"
)

//==============================================================
// CompliancePolicy Declaration
//==============================================================
type CompliancePolicy struct {
    gorm.Model
     Name                                    string
    PolicyCode                                    string
    Description                                    string
    InstitutionId         *uint
    Institution           *FinancialInstitution `gorm:"foreignKey:InstitutionId"`
    Status                      PolicyStatus

// parent associations as their child

}

