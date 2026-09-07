package model

import (
    "gorm.io/gorm"
)

//==============================================================
// PriceBook Declaration
//==============================================================
type PriceBook struct {
    gorm.Model
     Name                                    string
    AsActive                                    bool
    Description                                    string
    OrganizationId         *uint
    Organization           *Organization `gorm:"foreignKey:OrganizationId"`
     Entries           []PriceBookEntry `gorm:"foreignKey:EntriesFromPriceBookId"`
     Quotes           []Quote `gorm:"foreignKey:QuotesFromPriceBookId"`
     Orders           []Order `gorm:"foreignKey:OrdersFromPriceBookId"`

// parent associations as their child

}

