package model

import (
    "gorm.io/gorm"
)

//==============================================================
// AircraftPackage Declaration
//==============================================================
type AircraftPackage struct {
    gorm.Model
     Name                                    string
     Options           []AircraftOption `gorm:"foreignKey:OptionsFromAircraftPackageId"`
     Variants           []AircraftVariant `gorm:"foreignKey:VariantsFromAircraftPackageId"`
    PackageType                      PackageType

// parent associations as their child

}

