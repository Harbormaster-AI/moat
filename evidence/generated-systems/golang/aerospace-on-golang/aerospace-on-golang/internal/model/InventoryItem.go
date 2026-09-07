package model

import (
    "gorm.io/gorm"
)

//==============================================================
// InventoryItem Declaration
//==============================================================
type InventoryItem struct {
    gorm.Model
     QuantityOnHand                                                            string
    QuantityReserved                                                            string
    LotNumber                                    string
    ComponentId         *uint
    Component           *Component_ `gorm:"foreignKey:ComponentId"`
    WarehouseId         *uint
    Warehouse           *Warehouse `gorm:"foreignKey:WarehouseId"`

// parent associations as their child

}

