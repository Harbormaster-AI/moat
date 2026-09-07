package model

import (
    "time"
    "gorm.io/gorm"
)

//==============================================================
// InboundShipment Declaration
//==============================================================
type InboundShipment struct {
    gorm.Model
     ShipmentNumber                                    string
    ExpectedArrivalDate                                                            time.Time
    ArrivalDate                                                            time.Time
    CarrierName                                    string
    WarehouseId         *uint
    Warehouse           *Warehouse `gorm:"foreignKey:WarehouseId"`
     Lines           []InboundShipmentLine `gorm:"foreignKey:LinesFromInboundShipmentId"`
     Transactions           []InventoryTransaction `gorm:"foreignKey:TransactionsFromInboundShipmentId"`
    Status                      InboundShipmentStatus

// parent associations as their child

}

