package model

import (
    "time"
    "gorm.io/gorm"
)

//==============================================================
// Quote Declaration
//==============================================================
type Quote struct {
    gorm.Model
     QuoteNumber                                    string
    ValidityStart                                                            time.Time
    ValidityEnd                                                            time.Time
    TotalAmount                                                            string
    DiscountPercent                                                            string
    TaxAmount                                                            string
    ShippingAmount                                                            string
    OrganizationId         *uint
    Organization           *Organization `gorm:"foreignKey:OrganizationId"`
    AccountId         *uint
    Account           *Account `gorm:"foreignKey:AccountId"`
    OpportunityId         *uint
    Opportunity           *Opportunity `gorm:"foreignKey:OpportunityId"`
    OwnerId         *uint
    Owner           *User `gorm:"foreignKey:OwnerId"`
     LineItems           []QuoteLineItem `gorm:"foreignKey:LineItemsFromQuoteId"`
    PriceBookId         *uint
    PriceBook           *PriceBook `gorm:"foreignKey:PriceBookId"`
    OrderId         *uint
    Order           *Order `gorm:"foreignKey:OrderId"`
    Status                      QuoteStatus

// parent associations as their child

}

