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
    IssueDate                                                            time.Time
    DueDate                                                            time.Time
    Total                                                            string
    Currency                                    string
    MerchantId         *uint
    Merchant           *Merchant `gorm:"foreignKey:MerchantId"`
     Payments           []PaymentOrder `gorm:"foreignKey:PaymentsFromInvoiceId"`
    Status                      InvoiceStatus

// parent associations as their child

}

