package model

import (
    "time"
    "gorm.io/gorm"
)

//==============================================================
// InventoryItem Declaration
//==============================================================
type InventoryItem struct {
    gorm.Model
     QuantityOnHand                                                            string
    QuantityAvailable                                                            string
    QuantityReserved                                                            string
    UnitCost                                                            string
    LastUpdated                                                            time.Time
    SkuId         *uint
    Sku           *StockKeepingUnit `gorm:"foreignKey:SkuId"`
    WarehouseId         *uint
    Warehouse           *Warehouse `gorm:"foreignKey:WarehouseId"`
    LocationId         *uint
    Location           *StorageLocation `gorm:"foreignKey:LocationId"`
    LotId         *uint
    Lot           *Lot `gorm:"foreignKey:LotId"`
     SerialNumbers           []SerialNumber `gorm:"foreignKey:SerialNumbersFromInventoryItemId"`
     Transactions           []InventoryTransaction `gorm:"foreignKey:TransactionsFromInventoryItemId"`
     Reservations           []Reservation `gorm:"foreignKey:ReservationsFromInventoryItemId"`
    StockStatus                      StockStatus

// parent associations as their child

}

