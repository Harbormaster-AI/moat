package model

import (
    "time"
    "gorm.io/gorm"
)

//==============================================================
// InventoryTransaction Declaration
//==============================================================
type InventoryTransaction struct {
    gorm.Model
     TransactionNumber                                    string
    Quantity                                                            string
    TransactionDateTime                                                            time.Time
    ReferenceDocument                                    string
    ItemId         *uint
    Item           *Item `gorm:"foreignKey:ItemId"`
    LocationId         *uint
    Location           *Location `gorm:"foreignKey:LocationId"`
    WorkOrderId         *uint
    WorkOrder           *WorkOrder `gorm:"foreignKey:WorkOrderId"`
    PurchaseOrderId         *uint
    PurchaseOrder           *PurchaseOrder `gorm:"foreignKey:PurchaseOrderId"`
    SalesOrderId         *uint
    SalesOrder           *SalesOrder `gorm:"foreignKey:SalesOrderId"`
    TransactionType                      InventoryTransactionType

// parent associations as their child

}

