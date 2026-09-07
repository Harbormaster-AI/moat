package model

import (
    "time"
    "gorm.io/gorm"
)

//==============================================================
// Consent Declaration
//==============================================================
type Consent struct {
    gorm.Model
     GrantedAt                                                            time.Time
    ExpiresAt                                                            time.Time
    Scope                                    string
    CustomerId         *uint
    Customer           *Customer `gorm:"foreignKey:CustomerId"`
    ApiClientId         *uint
    ApiClient           *APIClient `gorm:"foreignKey:ApiClientId"`
    ConsentType                      ConsentType
    Status                      ConsentStatus

// parent associations as their child

}

