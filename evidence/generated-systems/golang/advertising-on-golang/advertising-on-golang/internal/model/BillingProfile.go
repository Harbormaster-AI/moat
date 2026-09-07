package model

import (
    "gorm.io/gorm"
)

//==============================================================
// BillingProfile Declaration
//==============================================================
type BillingProfile struct {
    gorm.Model
     BillingName                                    string
    TaxId                                    string
    BillingAddress                                                            string
    AdvertiserId         *uint
    Advertiser           *Advertiser `gorm:"foreignKey:AdvertiserId"`
     PaymentMethods           []PaymentMethod `gorm:"foreignKey:PaymentMethodsFromBillingProfileId"`
     AdAccounts           []AdAccount `gorm:"foreignKey:AdAccountsFromBillingProfileId"`
    PaymentTerms                      PaymentTerms

// parent associations as their child

}

