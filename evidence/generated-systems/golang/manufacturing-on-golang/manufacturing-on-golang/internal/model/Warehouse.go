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
    WarehouseCode                                    string
    Address                                                            string
    PlantId         *uint
    Plant           *Plant `gorm:"foreignKey:PlantId"`
     Locations           []Location `gorm:"foreignKey:LocationsFromWarehouseId"`
     InventoryItems           []InventoryItem `gorm:"foreignKey:InventoryItemsFromWarehouseId"`
    WarehouseType                      WarehouseType

// parent associations as their child

}

