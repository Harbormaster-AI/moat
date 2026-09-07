package model

import (
    "time"
    "gorm.io/gorm"
)

//==============================================================
// OutboundAllocation Declaration
//==============================================================
type OutboundAllocation struct {
    gorm.Model
     AllocationNumber                                    string
    AllocatedQuantity                                                            string
    AllocationDate                                                            time.Time
    WarehouseId         *uint
    Warehouse           *Warehouse `gorm:"foreignKey:WarehouseId"`
    SkuId         *uint
    Sku           *StockKeepingUnit `gorm:"foreignKey:SkuId"`
    InventoryItemId         *uint
    InventoryItem           *InventoryItem `gorm:"foreignKey:InventoryItemId"`
    ReservationId         *uint
    Reservation           *Reservation `gorm:"foreignKey:ReservationId"`
    LotId         *uint
    Lot           *Lot `gorm:"foreignKey:LotId"`
     SerialNumbers           []SerialNumber `gorm:"foreignKey:SerialNumbersFromOutboundAllocationId"`
    SourceLocationId         *uint
    SourceLocation           *StorageLocation `gorm:"foreignKey:SourceLocationId"`
    Status                      AllocationStatus

// parent associations as their child

}

