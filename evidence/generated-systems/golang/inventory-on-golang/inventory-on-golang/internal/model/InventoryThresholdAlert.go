package model

import (
    "time"
    "gorm.io/gorm"
)

//==============================================================
// InventoryThresholdAlert Declaration
//==============================================================
type InventoryThresholdAlert struct {
    gorm.Model
     AlertNumber                                    string
    DetectedAt                                                            time.Time
    Message                                    string
    SkuId         *uint
    Sku           *StockKeepingUnit `gorm:"foreignKey:SkuId"`
    WarehouseId         *uint
    Warehouse           *Warehouse `gorm:"foreignKey:WarehouseId"`
    LocationId         *uint
    Location           *StorageLocation `gorm:"foreignKey:LocationId"`
    RelatedPolicyId         *uint
    RelatedPolicy           *ReplenishmentPolicy `gorm:"foreignKey:RelatedPolicyId"`
    AlertType                      InventoryAlertType
    Status                      AlertStatus

// parent associations as their child

}

