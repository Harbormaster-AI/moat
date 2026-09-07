package model

import (
    "gorm.io/gorm"
)

//==============================================================
// Organization Declaration
//==============================================================
type Organization struct {
    gorm.Model
     Name                                    string
    DefaultCurrency                                    string
    DefaultLocale                                                            string
    Website                                                            string
     Users           []User `gorm:"foreignKey:UsersFromOrganizationId"`
     Accounts           []Account `gorm:"foreignKey:AccountsFromOrganizationId"`
     Teams           []Team `gorm:"foreignKey:TeamsFromOrganizationId"`
     Territories           []Territory `gorm:"foreignKey:TerritoriesFromOrganizationId"`
     Products           []Product `gorm:"foreignKey:ProductsFromOrganizationId"`
     PriceBooks           []PriceBook `gorm:"foreignKey:PriceBooksFromOrganizationId"`
     Campaigns           []Campaign `gorm:"foreignKey:CampaignsFromOrganizationId"`

// parent associations as their child

}

