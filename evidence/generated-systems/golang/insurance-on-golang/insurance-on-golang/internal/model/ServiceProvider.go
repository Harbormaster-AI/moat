package model

import (
    "gorm.io/gorm"
)

//==============================================================
// ServiceProvider Declaration
//==============================================================
type ServiceProvider struct {
    gorm.Model
     Name                                    string
    TaxId                                    string
     Claims           []Claim `gorm:"foreignKey:ClaimsFromServiceProviderId"`
    ProviderType                      ServiceProviderType
    NetworkStatus                      NetworkStatus

// parent associations as their child

}

