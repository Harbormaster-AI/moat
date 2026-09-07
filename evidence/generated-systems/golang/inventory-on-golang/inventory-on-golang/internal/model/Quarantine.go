package model

import (
    "time"
    "gorm.io/gorm"
)

//==============================================================
// Quarantine Declaration
//==============================================================
type Quarantine struct {
    gorm.Model
     Reason                                    string
    StartedAt                                                            time.Time
    ReleasedAt                                                            time.Time
    WarehouseId         *uint
    Warehouse           *Warehouse `gorm:"foreignKey:WarehouseId"`
     Items           []InventoryItem `gorm:"foreignKey:ItemsFromQuarantineId"`
    LotId         *uint
    Lot           *Lot `gorm:"foreignKey:LotId"`
     SerialNumbers           []SerialNumber `gorm:"foreignKey:SerialNumbersFromQuarantineId"`
    Disposition                      Disposition

// parent associations as their child

}

