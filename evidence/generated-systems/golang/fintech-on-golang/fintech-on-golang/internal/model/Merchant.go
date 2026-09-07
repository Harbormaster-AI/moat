package model

import (
    "gorm.io/gorm"
)

//==============================================================
// Merchant Declaration
//==============================================================
type Merchant struct {
    gorm.Model
     Name                                    string
    Mcc                                    string
    Url                                    string
    Country                                    string
    SettlementCurrency                                    string
     Terminals           []Terminal `gorm:"foreignKey:TerminalsFromMerchantId"`
     PaymentContracts           []PaymentContract `gorm:"foreignKey:PaymentContractsFromMerchantId"`
     Payouts           []Payout `gorm:"foreignKey:PayoutsFromMerchantId"`
     Settlements           []SettlementBatch `gorm:"foreignKey:SettlementsFromMerchantId"`
     Disputes           []Dispute `gorm:"foreignKey:DisputesFromMerchantId"`
     Invoices           []Invoice `gorm:"foreignKey:InvoicesFromMerchantId"`

// parent associations as their child

}

