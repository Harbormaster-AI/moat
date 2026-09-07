package model

import (
    "time"
    "gorm.io/gorm"
)

//==============================================================
// InventoryTransaction Declaration
//==============================================================
type InventoryTransaction struct {
    gorm.Model
     TransactionNumber                                    string
    Quantity                                                            string
    UnitCost                                                            string
    TransactionDate                                                            time.Time
    ReasonCode                                    string
    SkuId         *uint
    Sku           *StockKeepingUnit `gorm:"foreignKey:SkuId"`
    WarehouseId         *uint
    Warehouse           *Warehouse `gorm:"foreignKey:WarehouseId"`
    LocationId         *uint
    Location           *StorageLocation `gorm:"foreignKey:LocationId"`
    LotId         *uint
    Lot           *Lot `gorm:"foreignKey:LotId"`
     SerialNumbers           []SerialNumber `gorm:"foreignKey:SerialNumbersFromInventoryTransactionId"`
    RelatedReservationId         *uint
    RelatedReservation           *Reservation `gorm:"foreignKey:RelatedReservationId"`
    TransferOrderId         *uint
    TransferOrder           *TransferOrder `gorm:"foreignKey:TransferOrderId"`
    AdjustmentId         *uint
    Adjustment           *StockAdjustment `gorm:"foreignKey:AdjustmentId"`
    CycleCountId         *uint
    CycleCount           *CycleCount `gorm:"foreignKey:CycleCountId"`
    TransactionType                      TransactionType
    UnitOfMeasure                      UnitOfMeasure
    Status                      TransactionStatus

// parent associations as their child

}

