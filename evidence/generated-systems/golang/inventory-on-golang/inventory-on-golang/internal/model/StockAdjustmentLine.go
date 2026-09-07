package model

import (
    "gorm.io/gorm"
)

//==============================================================
// StockAdjustmentLine Declaration
//==============================================================
type StockAdjustmentLine struct {
    gorm.Model
     LineNumber                                                            string
    Quantity                                                            string
    AdjustmentId         *uint
    Adjustment           *StockAdjustment `gorm:"foreignKey:AdjustmentId"`
    SkuId         *uint
    Sku           *StockKeepingUnit `gorm:"foreignKey:SkuId"`
    LotId         *uint
    Lot           *Lot `gorm:"foreignKey:LotId"`
    LocationId         *uint
    Location           *StorageLocation `gorm:"foreignKey:LocationId"`
     SerialNumbers           []SerialNumber `gorm:"foreignKey:SerialNumbersFromStockAdjustmentLineId"`
    UnitOfMeasure                      UnitOfMeasure
    StockStatus                      StockStatus

// parent associations as their child

}

