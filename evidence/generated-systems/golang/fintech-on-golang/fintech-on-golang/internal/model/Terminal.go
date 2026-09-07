package model

import (
    "gorm.io/gorm"
)

//==============================================================
// Terminal Declaration
//==============================================================
type Terminal struct {
    gorm.Model
     Location                                                            string
    MerchantId         *uint
    Merchant           *Merchant `gorm:"foreignKey:MerchantId"`
    Type                      TerminalType
    Status                      TerminalStatus

// parent associations as their child

}

