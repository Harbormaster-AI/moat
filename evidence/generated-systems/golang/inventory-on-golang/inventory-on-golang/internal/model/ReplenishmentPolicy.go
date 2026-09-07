package model

import (
    "gorm.io/gorm"
)

//==============================================================
// ReplenishmentPolicy Declaration
//==============================================================
type ReplenishmentPolicy struct {
    gorm.Model
     MinLevel                                                            string
    MaxLevel                                                            string
    ReorderPoint                                                            string
    ReorderQuantity                                                            string
    LeadTimeDays                                                            string
    ReviewPeriodDays                                                            string
    SkuId         *uint
    Sku           *StockKeepingUnit `gorm:"foreignKey:SkuId"`
    WarehouseId         *uint
    Warehouse           *Warehouse `gorm:"foreignKey:WarehouseId"`
    LocationId         *uint
    Location           *StorageLocation `gorm:"foreignKey:LocationId"`
    PolicyType                      ReplenishmentPolicyType

// parent associations as their child

}

