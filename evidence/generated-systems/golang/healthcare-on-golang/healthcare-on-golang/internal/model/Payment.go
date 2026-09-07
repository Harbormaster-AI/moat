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
     PaymentNumber                                    string
    Amount                                                            string
    PaymentDate                                                            time.Time
    InvoiceId         *uint
    Invoice           *Invoice `gorm:"foreignKey:InvoiceId"`
    PayerId         *uint
    Payer           *InsurancePayer `gorm:"foreignKey:PayerId"`
    Method                      PaymentMethod

// parent associations as their child

}

