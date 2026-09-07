package model

import (
    "time"
    "gorm.io/gorm"
)

//==============================================================
// SerialNumber Declaration
//==============================================================
type SerialNumber struct {
    gorm.Model
     Serial                                                            string
    ActivationDate                                                            time.Time
    SkuId         *uint
    Sku           *StockKeepingUnit `gorm:"foreignKey:SkuId"`
    CurrentInventoryItemId         *uint
    CurrentInventoryItem           *InventoryItem `gorm:"foreignKey:CurrentInventoryItemId"`
    LotId         *uint
    Lot           *Lot `gorm:"foreignKey:LotId"`
    Status                      SerialStatus

// parent associations as their child

}

