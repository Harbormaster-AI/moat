package model

import (
    "gorm.io/gorm"
)

//==============================================================
// OpportunityLineItem Declaration
//==============================================================
type OpportunityLineItem struct {
    gorm.Model
     Quantity                                                            string
    UnitPrice                                                            string
    DiscountPercent                                                            string
    TotalPrice                                                            string
    OpportunityId         *uint
    Opportunity           *Opportunity `gorm:"foreignKey:OpportunityId"`
    ProductId         *uint
    Product           *Product `gorm:"foreignKey:ProductId"`
    PriceBookEntryId         *uint
    PriceBookEntry           *PriceBookEntry `gorm:"foreignKey:PriceBookEntryId"`

// parent associations as their child

}

