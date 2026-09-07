package model

import (
    "gorm.io/gorm"
)

//==============================================================
// Product Declaration
//==============================================================
type Product struct {
    gorm.Model
     Sku                                    string
    Name                                    string
    AsActive                                    bool
    StandardPrice                                                            string
    Description                                    string
    OrganizationId         *uint
    Organization           *Organization `gorm:"foreignKey:OrganizationId"`
     PriceBookEntries           []PriceBookEntry `gorm:"foreignKey:PriceBookEntriesFromProductId"`
     OpportunityLineItems           []OpportunityLineItem `gorm:"foreignKey:OpportunityLineItemsFromProductId"`
     QuoteLineItems           []QuoteLineItem `gorm:"foreignKey:QuoteLineItemsFromProductId"`
     OrderItems           []OrderItem `gorm:"foreignKey:OrderItemsFromProductId"`
    ProductType                      ProductType
    Uom                      UnitOfMeasure

// parent associations as their child

}

