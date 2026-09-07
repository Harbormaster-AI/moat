package model

import (
    "time"
    "gorm.io/gorm"
)

//==============================================================
// Payment Declaration
//==============================================================
type Payment struct {
    gorm.Model
     PaymentReference                                    string
    Amount                                                            string
    PaymentDate                                                            time.Time
    InvoiceId         *uint
    Invoice           *Invoice `gorm:"foreignKey:InvoiceId"`
    BillingAccountId         *uint
    BillingAccount           *BillingAccount `gorm:"foreignKey:BillingAccountId"`
    PolicyId         *uint
    Policy           *Policy `gorm:"foreignKey:PolicyId"`
    Method                      PaymentMethod
    Status                      PaymentStatus

// parent associations as their child

}

