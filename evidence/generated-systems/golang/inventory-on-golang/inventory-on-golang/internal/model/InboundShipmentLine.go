package model

import (
    "gorm.io/gorm"
)

//==============================================================
// InboundShipmentLine Declaration
//==============================================================
type InboundShipmentLine struct {
    gorm.Model
     LineNumber                                                            string
    Quantity                                                            string
    InboundShipmentId         *uint
    InboundShipment           *InboundShipment `gorm:"foreignKey:InboundShipmentId"`
    SkuId         *uint
    Sku           *StockKeepingUnit `gorm:"foreignKey:SkuId"`
    LotId         *uint
    Lot           *Lot `gorm:"foreignKey:LotId"`
     SerialNumbers           []SerialNumber `gorm:"foreignKey:SerialNumbersFromInboundShipmentLineId"`
    DestinationLocationId         *uint
    DestinationLocation           *StorageLocation `gorm:"foreignKey:DestinationLocationId"`
    UnitOfMeasure                      UnitOfMeasure
    StockStatus                      StockStatus

// parent associations as their child

}

