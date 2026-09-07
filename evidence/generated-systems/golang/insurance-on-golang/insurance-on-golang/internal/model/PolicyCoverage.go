package model

import (
    "gorm.io/gorm"
)

//==============================================================
// PolicyCoverage Declaration
//==============================================================
type PolicyCoverage struct {
    gorm.Model
     Limit                                                            string
    Deductible                                                            string
    Premium                                                            string
    PolicyId         *uint
    Policy           *Policy `gorm:"foreignKey:PolicyId"`
     InsuredObjects           []InsuredObject `gorm:"foreignKey:InsuredObjectsFromPolicyCoverageId"`
    CoverageType                      CoverageType

// parent associations as their child

}

