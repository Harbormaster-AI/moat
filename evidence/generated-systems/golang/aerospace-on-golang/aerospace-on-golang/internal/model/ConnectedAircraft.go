package model

import (
    "gorm.io/gorm"
)

//==============================================================
// ConnectedAircraft Declaration
//==============================================================
type ConnectedAircraft struct {
    gorm.Model
     CommunicationsProvider                                    string
    AircraftId         *uint
    Aircraft           *Aircraft `gorm:"foreignKey:AircraftId"`
     FlightHealthEvents           []FlightHealthEvent `gorm:"foreignKey:FlightHealthEventsFromConnectedAircraftId"`
     SoftwareLoads           []SoftwareLoad `gorm:"foreignKey:SoftwareLoadsFromConnectedAircraftId"`
    ConnectivityStatus                      ConnectivityStatus

// parent associations as their child

}

