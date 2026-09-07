package model

import (
    "time"
    "gorm.io/gorm"
)

//==============================================================
// StockAdjustment Declaration
//==============================================================
type StockAdjustment struct {
    gorm.Model
     AdjustmentNumber                                    string
    Reason                                    string
    AdjustmentDate                                                            time.Time
    WarehouseId         *uint
    Warehouse           *Warehouse `gorm:"foreignKey:WarehouseId"`
     Lines           []StockAdjustmentLine `gorm:"foreignKey:LinesFromStockAdjustmentId"`
     Transactions           []InventoryTransaction `gorm:"foreignKey:TransactionsFromStockAdjustmentId"`
    AdjustmentType                      AdjustmentType
    Status                      AdjustmentStatus

// parent associations as their child

}

