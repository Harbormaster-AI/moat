package model

import (
    "gorm.io/gorm"
)

//==============================================================
// Beneficiary Declaration
//==============================================================
type Beneficiary struct {
    gorm.Model
     Name                                    string
    AccountIdentifier                                                            string
    Iban                                                            string
    Bic                                                            string
    Address                                                            string
    CustomerId         *uint
    Customer           *Customer `gorm:"foreignKey:CustomerId"`

// parent associations as their child

}

