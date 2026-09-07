package model

import (
    "time"
    "gorm.io/gorm"
)

//==============================================================
// Transaction Declaration
//==============================================================
type Transaction struct {
    gorm.Model
     Amount                                                            string
    Fee                                                            string
    ExchangeRate                                                            string
    CreatedAt                                                            time.Time
    CompletedAt                                                            time.Time
    Narrative                                    string
    AccountId         *uint
    Account           *Account `gorm:"foreignKey:AccountId"`
    WalletId         *uint
    Wallet           *Wallet `gorm:"foreignKey:WalletId"`
    PaymentOrderId         *uint
    PaymentOrder           *PaymentOrder `gorm:"foreignKey:PaymentOrderId"`
    MerchantId         *uint
    Merchant           *Merchant `gorm:"foreignKey:MerchantId"`
    CardId         *uint
    Card           *PaymentCard `gorm:"foreignKey:CardId"`
     RelatedTransactions           []Transaction `gorm:"foreignKey:RelatedTransactionsFromTransactionId"`
     Alerts           []ComplianceAlert `gorm:"foreignKey:AlertsFromTransactionId"`
    TransactionType                      TransactionType
    Status                      TransactionStatus

// parent associations as their child

}

