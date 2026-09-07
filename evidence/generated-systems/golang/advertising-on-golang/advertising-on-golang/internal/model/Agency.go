package model

import (
    "gorm.io/gorm"
)

//==============================================================
// Agency Declaration
//==============================================================
type Agency struct {
    gorm.Model
     Name                                    string
    LegalName                                    string
    HeadquartersCountry                                    string
    Website                                    string
     Advertisers           []Advertiser `gorm:"foreignKey:AdvertisersFromAgencyId"`
     Teams           []Team `gorm:"foreignKey:TeamsFromAgencyId"`
     Users           []User `gorm:"foreignKey:UsersFromAgencyId"`
     InsertionOrders           []InsertionOrder `gorm:"foreignKey:InsertionOrdersFromAgencyId"`

// parent associations as their child

}

