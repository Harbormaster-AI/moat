package model

import (
    "gorm.io/gorm"
)

//==============================================================
// DSP Declaration
//==============================================================
type DSP struct {
    gorm.Model
     Name                                    string
    Website                                    string
    Region                                    string
     AdAccounts           []AdAccount `gorm:"foreignKey:AdAccountsFromDSPId"`

// parent associations as their child

}

