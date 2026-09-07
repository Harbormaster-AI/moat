package model

import (
    "gorm.io/gorm"
)

//==============================================================
// OrderItem Declaration
//==============================================================
type OrderItem struct {
    gorm.Model
     Quantity                                                            string
    UnitPrice                                                            string
    DiscountAmount                                                            string
    TaxAmount                                                            string
    TotalAmount                                                            string
    OrderId         *uint
    Order           *Order `gorm:"foreignKey:OrderId"`
    ProductId         *uint
    Product           *Product `gorm:"foreignKey:ProductId"`
    PriceBookEntryId         *uint
    PriceBookEntry           *PriceBookEntry `gorm:"foreignKey:PriceBookEntryId"`

// parent associations as their child

}

