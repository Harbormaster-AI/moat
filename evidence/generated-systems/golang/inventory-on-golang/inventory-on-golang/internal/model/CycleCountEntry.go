package model

import (
    "gorm.io/gorm"
)

//==============================================================
// CycleCountEntry Declaration
//==============================================================
type CycleCountEntry struct {
    gorm.Model
     LineNumber                                                            string
    SystemQuantity                                                            string
    CountedQuantity                                                            string
    VarianceQuantity                                                            string
    RecountRequired                                    bool
    CycleCountId         *uint
    CycleCount           *CycleCount `gorm:"foreignKey:CycleCountId"`
    SkuId         *uint
    Sku           *StockKeepingUnit `gorm:"foreignKey:SkuId"`
    LotId         *uint
    Lot           *Lot `gorm:"foreignKey:LotId"`
    LocationId         *uint
    Location           *StorageLocation `gorm:"foreignKey:LocationId"`
     SerialNumbers           []SerialNumber `gorm:"foreignKey:SerialNumbersFromCycleCountEntryId"`
    StockStatus                      StockStatus

// parent associations as their child

}

