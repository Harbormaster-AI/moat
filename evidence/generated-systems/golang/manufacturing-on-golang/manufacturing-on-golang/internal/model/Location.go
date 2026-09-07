package model

import (
    "gorm.io/gorm"
)

//==============================================================
// Location Declaration
//==============================================================
type Location struct {
    gorm.Model
     LocationCode                                    string
    Description                                    string
    WarehouseId         *uint
    Warehouse           *Warehouse `gorm:"foreignKey:WarehouseId"`
     InventoryItems           []InventoryItem `gorm:"foreignKey:InventoryItemsFromLocationId"`
    LocationType                      LocationType

// parent associations as their child

}

