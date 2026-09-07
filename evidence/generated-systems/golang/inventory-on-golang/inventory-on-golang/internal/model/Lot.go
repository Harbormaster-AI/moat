package model

import (
    "time"
    "gorm.io/gorm"
)

//==============================================================
// Lot Declaration
//==============================================================
type Lot struct {
    gorm.Model
     BatchNumber                                                            string
    ManufactureDate                                                            time.Time
    ExpirationDate                                                            time.Time
    SkuId         *uint
    Sku           *StockKeepingUnit `gorm:"foreignKey:SkuId"`
     InventoryItems           []InventoryItem `gorm:"foreignKey:InventoryItemsFromLotId"`
    LotStatus                      LotStatus

// parent associations as their child

}

