package model

import (
    "time"
    "gorm.io/gorm"
)

//==============================================================
// DirectDebitMandate Declaration
//==============================================================
type DirectDebitMandate struct {
    gorm.Model
     MandateId                                    string
    SignedAt                                                            time.Time
    AccountId         *uint
    Account           *Account `gorm:"foreignKey:AccountId"`
    CreditorId         *uint
    Creditor           *Creditor `gorm:"foreignKey:CreditorId"`
    Scheme                      DirectDebitScheme
    Status                      MandateStatus

// parent associations as their child

}

