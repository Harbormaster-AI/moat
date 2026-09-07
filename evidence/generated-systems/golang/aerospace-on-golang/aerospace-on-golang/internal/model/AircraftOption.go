package model

import (
    "gorm.io/gorm"
)

//==============================================================
// AircraftOption Declaration
//==============================================================
type AircraftOption struct {
    gorm.Model
     Code                                    string
    Name                                    string
     Variants           []AircraftVariant `gorm:"foreignKey:VariantsFromAircraftOptionId"`
     Packages           []AircraftPackage `gorm:"foreignKey:PackagesFromAircraftOptionId"`
    OptionCategory                      OptionCategory

// parent associations as their child

}

