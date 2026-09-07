package model

import (
    "gorm.io/gorm"
)

//==============================================================
// Creditor Declaration
//==============================================================
type Creditor struct {
    gorm.Model
     Name                                    string
    Bic                                                            string
    Address                                                            string
     Mandates           []DirectDebitMandate `gorm:"foreignKey:MandatesFromCreditorId"`

// parent associations as their child

}

