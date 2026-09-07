package model

import (
    "time"
    "gorm.io/gorm"
)

//==============================================================
// SalesOrderLine Declaration
//==============================================================
type SalesOrderLine struct {
    gorm.Model
     LineNumber                                                            string
    Quantity                                                            string
    UnitPrice                                                            string
    DueDate                                                            time.Time
    SalesOrderId         *uint
    SalesOrder           *SalesOrder `gorm:"foreignKey:SalesOrderId"`
    ItemId         *uint
    Item           *Item `gorm:"foreignKey:ItemId"`

// parent associations as their child

}

