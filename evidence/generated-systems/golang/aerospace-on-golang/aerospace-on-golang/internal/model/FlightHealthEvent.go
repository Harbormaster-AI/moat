package model

import (
    "gorm.io/gorm"
)

//==============================================================
// FlightHealthEvent Declaration
//==============================================================
type FlightHealthEvent struct {
    gorm.Model
     EventCode                                    string
    ConnectedAircraftId         *uint
    ConnectedAircraft           *ConnectedAircraft `gorm:"foreignKey:ConnectedAircraftId"`
    Severity                      EventSeverity

// parent associations as their child

}

