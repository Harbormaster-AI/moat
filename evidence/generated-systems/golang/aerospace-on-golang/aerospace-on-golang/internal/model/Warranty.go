package model

import (
    "gorm.io/gorm"
)

//==============================================================
// Warranty Declaration
//==============================================================
type Warranty struct {
    gorm.Model
     CoverageMonths                                                            string
    AircraftId         *uint
    Aircraft           *Aircraft `gorm:"foreignKey:AircraftId"`
    WarrantyType                      WarrantyType

// parent associations as their child

}

