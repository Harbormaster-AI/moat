package model

import (
    "time"
    "gorm.io/gorm"
)

//==============================================================
// Order Declaration
//==============================================================
type Order struct {
    gorm.Model
     OrderNumber                                    string
    OrderDate                                                            time.Time
    TotalAmount                                                            string
    TaxAmount                                                            string
    ShippingAmount                                                            string
    OrganizationId         *uint
    Organization           *Organization `gorm:"foreignKey:OrganizationId"`
    AccountId         *uint
    Account           *Account `gorm:"foreignKey:AccountId"`
    OpportunityId         *uint
    Opportunity           *Opportunity `gorm:"foreignKey:OpportunityId"`
    QuoteId         *uint
    Quote           *Quote `gorm:"foreignKey:QuoteId"`
    OwnerId         *uint
    Owner           *User `gorm:"foreignKey:OwnerId"`
     Items           []OrderItem `gorm:"foreignKey:ItemsFromOrderId"`
    ContractId         *uint
    Contract           *Contract `gorm:"foreignKey:ContractId"`
    PriceBookId         *uint
    PriceBook           *PriceBook `gorm:"foreignKey:PriceBookId"`
    Status                      OrderStatus

// parent associations as their child

}

