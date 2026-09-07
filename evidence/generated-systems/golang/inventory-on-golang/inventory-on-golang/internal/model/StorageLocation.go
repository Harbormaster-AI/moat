package model

import (
    "gorm.io/gorm"
)

//==============================================================
// StorageLocation Declaration
//==============================================================
type StorageLocation struct {
    gorm.Model
     Code                                    string
    TemperatureControlled                                    bool
    Capacity                                                            string
    CapacityUnit                                    string
    WarehouseId         *uint
    Warehouse           *Warehouse `gorm:"foreignKey:WarehouseId"`
    ParentLocationId         *uint
    ParentLocation           *StorageLocation `gorm:"foreignKey:ParentLocationId"`
     ChildLocations           []StorageLocation `gorm:"foreignKey:ChildLocationsFromStorageLocationId"`
     InventoryItems           []InventoryItem `gorm:"foreignKey:InventoryItemsFromStorageLocationId"`
    LocationType                      LocationType

// parent associations as their child

}

