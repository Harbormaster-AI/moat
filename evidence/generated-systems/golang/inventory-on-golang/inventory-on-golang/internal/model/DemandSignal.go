package model

import (
    "time"
    "gorm.io/gorm"
)

//==============================================================
// DemandSignal Declaration
//==============================================================
type DemandSignal struct {
    gorm.Model
     ExternalReference                                    string
    RequestedDate                                                            time.Time
    Quantity                                                            string
    SkuId         *uint
    Sku           *StockKeepingUnit `gorm:"foreignKey:SkuId"`
     Reservations           []Reservation `gorm:"foreignKey:ReservationsFromDemandSignalId"`
    DemandType                      DemandType

// parent associations as their child

}

