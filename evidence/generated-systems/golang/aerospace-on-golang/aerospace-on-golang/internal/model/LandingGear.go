package model

import (
    "gorm.io/gorm"
)

//==============================================================
// LandingGear Declaration
//==============================================================
type LandingGear struct {
    gorm.Model
     SupplierPartNumber                                    string
    SupplierId         *uint
    Supplier           *Supplier `gorm:"foreignKey:SupplierId"`
     Variants           []AircraftVariant `gorm:"foreignKey:VariantsFromLandingGearId"`
    GearType                      LandingGearType

// parent associations as their child

}

