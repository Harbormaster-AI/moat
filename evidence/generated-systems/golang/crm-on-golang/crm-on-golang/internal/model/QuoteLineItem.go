package model

import (
    "gorm.io/gorm"
)

//==============================================================
// QuoteLineItem Declaration
//==============================================================
type QuoteLineItem struct {
    gorm.Model
     Quantity                                                            string
    UnitPrice                                                            string
    DiscountAmount                                                            string
    TaxAmount                                                            string
    TotalAmount                                                            string
    QuoteId         *uint
    Quote           *Quote `gorm:"foreignKey:QuoteId"`
    ProductId         *uint
    Product           *Product `gorm:"foreignKey:ProductId"`
    PriceBookEntryId         *uint
    PriceBookEntry           *PriceBookEntry `gorm:"foreignKey:PriceBookEntryId"`
    OpportunityLineItemId         *uint
    OpportunityLineItem           *OpportunityLineItem `gorm:"foreignKey:OpportunityLineItemId"`

// parent associations as their child

}

