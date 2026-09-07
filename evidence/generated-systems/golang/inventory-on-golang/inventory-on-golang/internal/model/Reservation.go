package model

import (
    "time"
    "gorm.io/gorm"
)

//==============================================================
// Reservation Declaration
//==============================================================
type Reservation struct {
    gorm.Model
     ReferenceNumber                                    string
    ReservedQuantity                                                            string
    PromisedDate                                                            time.Time
    SkuId         *uint
    Sku           *StockKeepingUnit `gorm:"foreignKey:SkuId"`
    WarehouseId         *uint
    Warehouse           *Warehouse `gorm:"foreignKey:WarehouseId"`
    LocationId         *uint
    Location           *StorageLocation `gorm:"foreignKey:LocationId"`
    InventoryItemId         *uint
    InventoryItem           *InventoryItem `gorm:"foreignKey:InventoryItemId"`
    LotId         *uint
    Lot           *Lot `gorm:"foreignKey:LotId"`
     SerialNumbers           []SerialNumber `gorm:"foreignKey:SerialNumbersFromReservationId"`
    DemandSignalId         *uint
    DemandSignal           *DemandSignal `gorm:"foreignKey:DemandSignalId"`
    ReservationStatus                      ReservationStatus
    ReservationType                      ReservationType

// parent associations as their child

}

