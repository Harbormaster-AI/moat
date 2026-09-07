package model

import (
    "gorm.io/gorm"
)

//==============================================================
// Branch Declaration
//==============================================================
type Branch struct {
    gorm.Model
     Name                                    string
    BranchCode                                    string
    Address                                                            string
    InstitutionId         *uint
    Institution           *FinancialInstitution `gorm:"foreignKey:InstitutionId"`

// parent associations as their child

}

