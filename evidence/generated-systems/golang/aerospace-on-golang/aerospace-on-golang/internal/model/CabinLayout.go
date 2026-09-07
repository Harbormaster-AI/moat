package model

import (
    "gorm.io/gorm"
)

//==============================================================
// CabinLayout Declaration
//==============================================================
type CabinLayout struct {
    gorm.Model
     LayoutCode                                    string
    TotalSeats                                                            string
    ClassLayout                                    string
    VariantId         *uint
    Variant           *AircraftVariant `gorm:"foreignKey:VariantId"`
     Aircraft           []Aircraft `gorm:"foreignKey:AircraftFromCabinLayoutId"`
     Options           []AircraftOption `gorm:"foreignKey:OptionsFromCabinLayoutId"`

// parent associations as their child

}

