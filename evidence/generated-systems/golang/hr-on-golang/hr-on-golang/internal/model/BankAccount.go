package model

import (
    "gorm.io/gorm"
)

//==============================================================
// BankAccount Declaration
//==============================================================
type BankAccount struct {
    gorm.Model
     AccountHolder                                    string
    BankName                                    string
    Iban                                    string
    Bic                                    string
    AccountNumber                                    string
    RoutingNumber                                    string

// parent associations as their child

}

