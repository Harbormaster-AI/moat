package model

import (
    "gorm.io/gorm"
)

//==============================================================
// ThirdParty Declaration
//==============================================================
type ThirdParty struct {
    gorm.Model
     Name                                    string
    TaxId                                    string
    Address                                                            string
     Subrogations           []SubrogationRecovery `gorm:"foreignKey:SubrogationsFromThirdPartyId"`
    PartyType                      ThirdPartyType

// parent associations as their child

}

