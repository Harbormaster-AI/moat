package model

import (
    "gorm.io/gorm"
)

//==============================================================
// AvionicsSuite Declaration
//==============================================================
type AvionicsSuite struct {
    gorm.Model
     SuiteName                                    string
    SoftwareBaseline                                    string
    SupplierId         *uint
    Supplier           *Supplier `gorm:"foreignKey:SupplierId"`
     Variants           []AircraftVariant `gorm:"foreignKey:VariantsFromAvionicsSuiteId"`
     SoftwareLoads           []SoftwareLoad `gorm:"foreignKey:SoftwareLoadsFromAvionicsSuiteId"`

// parent associations as their child

}

