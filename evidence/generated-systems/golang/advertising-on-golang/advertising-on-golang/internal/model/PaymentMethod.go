package model

import (
    "gorm.io/gorm"
)

//==============================================================
// PaymentMethod Declaration
//==============================================================
type PaymentMethod struct {
    gorm.Model
     Last4                                    string
    CardholderName                                    string
    BillingAddress                                                            string
    BillingProfileId         *uint
    BillingProfile           *BillingProfile `gorm:"foreignKey:BillingProfileId"`
    MethodType                      PaymentMethodType

// parent associations as their child

}

