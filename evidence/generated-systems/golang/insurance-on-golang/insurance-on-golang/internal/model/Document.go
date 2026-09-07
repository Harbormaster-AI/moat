package model

import (
    "time"
    "gorm.io/gorm"
)

//==============================================================
// Document Declaration
//==============================================================
type Document struct {
    gorm.Model
     FileName                                    string
    UploadedDate                                                            time.Time
    PolicyId         *uint
    Policy           *Policy `gorm:"foreignKey:PolicyId"`
    ClaimId         *uint
    Claim           *Claim `gorm:"foreignKey:ClaimId"`
    ApplicationId         *uint
    Application           *Application `gorm:"foreignKey:ApplicationId"`
    CustomerId         *uint
    Customer           *Customer `gorm:"foreignKey:CustomerId"`
    DocumentType                      DocumentType

// parent associations as their child

}

