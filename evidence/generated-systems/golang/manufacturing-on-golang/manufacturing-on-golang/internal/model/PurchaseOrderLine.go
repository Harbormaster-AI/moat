package model

import (
    "time"
    "gorm.io/gorm"
)

//==============================================================
// PurchaseOrderLine Declaration
//==============================================================
type PurchaseOrderLine struct {
    gorm.Model
     LineNumber                                                            string
    Quantity                                                            string
    UnitPrice                                                            string
    DueDate                                                            time.Time
    PurchaseOrderId         *uint
    PurchaseOrder           *PurchaseOrder `gorm:"foreignKey:PurchaseOrderId"`
    ItemId         *uint
    Item           *Item `gorm:"foreignKey:ItemId"`

// parent associations as their child

}

