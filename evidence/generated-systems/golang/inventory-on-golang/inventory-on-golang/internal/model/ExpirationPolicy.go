package model

import (
    "gorm.io/gorm"
)

//==============================================================
// ExpirationPolicy Declaration
//==============================================================
type ExpirationPolicy struct {
    gorm.Model
     RejectIfDaysToExpireLessThan                                                            string
    AutoQuarantineDaysToExpire                                                            string
    SkuId         *uint
    Sku           *StockKeepingUnit `gorm:"foreignKey:SkuId"`
    WarehouseId         *uint
    Warehouse           *Warehouse `gorm:"foreignKey:WarehouseId"`
    RotationMethod                      RotationMethod

// parent associations as their child

}

