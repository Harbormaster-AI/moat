package model

import (
    "gorm.io/gorm"
)

//==============================================================
// InsuranceProduct Declaration
//==============================================================
type InsuranceProduct struct {
    gorm.Model
     Name                                    string
    ProductCode                                    string
    InsurerId         *uint
    Insurer           *Insurer `gorm:"foreignKey:InsurerId"`
     CoverageDefinitions           []CoverageDefinition `gorm:"foreignKey:CoverageDefinitionsFromInsuranceProductId"`
    LineOfBusiness                      LineOfBusiness

// parent associations as their child

}

