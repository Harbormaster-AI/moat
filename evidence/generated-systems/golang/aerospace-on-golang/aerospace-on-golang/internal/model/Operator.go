package model

import (
    "gorm.io/gorm"
)

//==============================================================
// Operator Declaration
//==============================================================
type Operator struct {
    gorm.Model
     Name                                    string
    IcaoDesignator                                    string
     AircraftOrders           []AircraftOrder `gorm:"foreignKey:AircraftOrdersFromOperatorId"`
     OperatedAircraft           []Aircraft `gorm:"foreignKey:OperatedAircraftFromOperatorId"`
    SalesRegionId         *uint
    SalesRegion           *SalesRegion `gorm:"foreignKey:SalesRegionId"`
    OperatorType                      OperatorType

// parent associations as their child

}

