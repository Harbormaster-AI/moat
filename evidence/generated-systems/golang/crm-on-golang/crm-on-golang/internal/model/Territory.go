package model

import (
    "gorm.io/gorm"
)

//==============================================================
// Territory Declaration
//==============================================================
type Territory struct {
    gorm.Model
     Name                                    string
    Region                                    string
    OrganizationId         *uint
    Organization           *Organization `gorm:"foreignKey:OrganizationId"`
     Accounts           []Account `gorm:"foreignKey:AccountsFromTerritoryId"`
     Users           []User `gorm:"foreignKey:UsersFromTerritoryId"`
    TerritoryType                      TerritoryType

// parent associations as their child

}

