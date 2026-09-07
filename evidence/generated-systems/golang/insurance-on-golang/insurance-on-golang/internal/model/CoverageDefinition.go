package model

import (
    "gorm.io/gorm"
)

//==============================================================
// CoverageDefinition Declaration
//==============================================================
type CoverageDefinition struct {
    gorm.Model
     Name                                    string
    DefaultLimit                                                            string
    DefaultDeductible                                                            string
    AsMandatory                                    bool
    ProductId         *uint
    Product           *InsuranceProduct `gorm:"foreignKey:ProductId"`
    CoverageType                      CoverageType

// parent associations as their child

}

