package model

import (
    "gorm.io/gorm"
)

//==============================================================
// Registration Declaration
//==============================================================
type Registration struct {
    gorm.Model
     TailNumber                                                            string
    RegistryCountry                                    string
    AircraftId         *uint
    Aircraft           *Aircraft `gorm:"foreignKey:AircraftId"`

// parent associations as their child

}

