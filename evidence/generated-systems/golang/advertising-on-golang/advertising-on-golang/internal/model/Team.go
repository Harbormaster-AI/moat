package model

import (
    "gorm.io/gorm"
)

//==============================================================
// Team Declaration
//==============================================================
type Team struct {
    gorm.Model
     Name                                    string
    AgencyId         *uint
    Agency           *Agency `gorm:"foreignKey:AgencyId"`
     Users           []User `gorm:"foreignKey:UsersFromTeamId"`
     AdAccounts           []AdAccount `gorm:"foreignKey:AdAccountsFromTeamId"`

// parent associations as their child

}

