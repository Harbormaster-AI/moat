package model

import (
    "time"
    "gorm.io/gorm"
)

//==============================================================
// Invoice Declaration
//==============================================================
type Invoice struct {
    gorm.Model
     InvoiceNumber                                    string
    DueDate                                                            time.Time
    TotalDue                                                            string
    BillingAccountId         *uint
    BillingAccount           *BillingAccount `gorm:"foreignKey:BillingAccountId"`
    PolicyId         *uint
    Policy           *Policy `gorm:"foreignKey:PolicyId"`
     Payments           []Payment `gorm:"foreignKey:PaymentsFromInvoiceId"`
    Status                      InvoiceStatus

// parent associations as their child

}

