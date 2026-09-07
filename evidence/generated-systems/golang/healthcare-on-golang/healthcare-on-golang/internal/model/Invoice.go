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
    TotalAmount                                                            string
    DueDate                                                            time.Time
    PatientId         *uint
    Patient           *Patient `gorm:"foreignKey:PatientId"`
    ClaimId         *uint
    Claim           *Claim `gorm:"foreignKey:ClaimId"`
     Payments           []Payment `gorm:"foreignKey:PaymentsFromInvoiceId"`
    Status                      InvoiceStatus

// parent associations as their child

}

