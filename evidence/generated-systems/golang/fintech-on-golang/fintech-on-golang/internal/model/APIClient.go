package model

import (
    "gorm.io/gorm"
)

//==============================================================
// APIClient Declaration
//==============================================================
type APIClient struct {
    gorm.Model
     Name                                    string
    ClientId                                    string
    RedirectUri                                    string
     Consents           []Consent `gorm:"foreignKey:ConsentsFromAPIClientId"`
    ClientType                      ClientType

// parent associations as their child

}

