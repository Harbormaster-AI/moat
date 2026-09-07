package model

import (
    "gorm.io/gorm"
)

//==============================================================
// SoftwareLoad Declaration
//==============================================================
type SoftwareLoad struct {
    gorm.Model
     Version                                    string
    ConnectedAircraftId         *uint
    ConnectedAircraft           *ConnectedAircraft `gorm:"foreignKey:ConnectedAircraftId"`
    AvionicsSuiteId         *uint
    AvionicsSuite           *AvionicsSuite `gorm:"foreignKey:AvionicsSuiteId"`
    LoadType                      SoftwareLoadType

// parent associations as their child

}

