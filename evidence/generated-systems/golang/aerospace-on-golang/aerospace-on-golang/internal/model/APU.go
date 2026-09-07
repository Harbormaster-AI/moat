package model

import (
    "gorm.io/gorm"
)

//==============================================================
// APU Declaration
//==============================================================
type APU struct {
    gorm.Model
     Model_                                    string
    SupplierId         *uint
    Supplier           *Supplier `gorm:"foreignKey:SupplierId"`
     Variants           []AircraftVariant `gorm:"foreignKey:VariantsFromAPUId"`

// parent associations as their child

}

