package model

import (
    "gorm.io/gorm"
)

//==============================================================
// BillingAccount Declaration
//==============================================================
type BillingAccount struct {
    gorm.Model
     AccountNumber                                    string
    Balance                                                            string
    CustomerId         *uint
    Customer           *Customer `gorm:"foreignKey:CustomerId"`
     Policies           []Policy `gorm:"foreignKey:PoliciesFromBillingAccountId"`
     Invoices           []Invoice `gorm:"foreignKey:InvoicesFromBillingAccountId"`
     Payments           []Payment `gorm:"foreignKey:PaymentsFromBillingAccountId"`
    Status                      BillingStatus

// parent associations as their child

}

