package model

import (
    "gorm.io/gorm"
)

//==============================================================
// User Declaration
//==============================================================
type User struct {
    gorm.Model
     FirstName                                    string
    LastName                                    string
    Email                                                            string
    AgencyId         *uint
    Agency           *Agency `gorm:"foreignKey:AgencyId"`
     Teams           []Team `gorm:"foreignKey:TeamsFromUserId"`
     AdAccounts           []AdAccount `gorm:"foreignKey:AdAccountsFromUserId"`
    Role                      AccountRole

// parent associations as their child

}

