package model

import (
    "gorm.io/gorm"
)

//==============================================================
// Warehouse Declaration
//==============================================================
type Warehouse struct {
    gorm.Model
     Name                                    string
     InventoryItems           []InventoryItem `gorm:"foreignKey:InventoryItemsFromWarehouseId"`

// parent associations as their child

}

