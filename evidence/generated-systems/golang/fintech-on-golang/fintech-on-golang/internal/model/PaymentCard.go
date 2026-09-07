package model

import (
    "gorm.io/gorm"
)

//==============================================================
// PaymentCard Declaration
//==============================================================
type PaymentCard struct {
    gorm.Model
     CardToken                                                            string
    MaskedPan                                    string
    ExpiryMonth                                                            string
    ExpiryYear                                                            string
    CardholderName                                    string
    CustomerId         *uint
    Customer           *Customer `gorm:"foreignKey:CustomerId"`
    AccountId         *uint
    Account           *Account `gorm:"foreignKey:AccountId"`
     Tokenizations           []CardTokenization `gorm:"foreignKey:TokenizationsFromPaymentCardId"`
     Disputes           []Dispute `gorm:"foreignKey:DisputesFromPaymentCardId"`
    Scheme                      CardScheme
    Status                      CardStatus

// parent associations as their child

}

