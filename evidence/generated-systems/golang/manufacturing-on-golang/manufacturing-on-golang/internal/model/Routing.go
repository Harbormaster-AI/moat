package model

import (
    "time"
    "gorm.io/gorm"
)

//==============================================================
// Routing Declaration
//==============================================================
type Routing struct {
    gorm.Model
     RoutingNumber                                    string
    Revision                                    string
    EffectivityStart                                                            time.Time
    EffectivityEnd                                                            time.Time
    ItemId         *uint
    Item           *Item `gorm:"foreignKey:ItemId"`
     Operations           []Operation `gorm:"foreignKey:OperationsFromRoutingId"`
    RoutingType                      RoutingType
    Status                      RoutingStatus

// parent associations as their child

}

