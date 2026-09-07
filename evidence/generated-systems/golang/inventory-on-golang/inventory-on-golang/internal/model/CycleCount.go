package model

import (
    "time"
    "gorm.io/gorm"
)

//==============================================================
// CycleCount Declaration
//==============================================================
type CycleCount struct {
    gorm.Model
     CountNumber                                    string
    ScheduledDate                                                            time.Time
    PerformedDate                                                            time.Time
    ApprovedBy                                    string
    WarehouseId         *uint
    Warehouse           *Warehouse `gorm:"foreignKey:WarehouseId"`
     Locations           []StorageLocation `gorm:"foreignKey:LocationsFromCycleCountId"`
     Entries           []CycleCountEntry `gorm:"foreignKey:EntriesFromCycleCountId"`
     Transactions           []InventoryTransaction `gorm:"foreignKey:TransactionsFromCycleCountId"`
    Status                      CountStatus

// parent associations as their child

}

