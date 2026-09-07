package model

import (
    "time"
    "gorm.io/gorm"
)

//==============================================================
// KYCProfile Declaration
//==============================================================
type KYCProfile struct {
    gorm.Model
     ProfileId                                    string
    CreatedAt                                                            time.Time
    CustomerId         *uint
    Customer           *Customer `gorm:"foreignKey:CustomerId"`
     Documents           []KYCDocument `gorm:"foreignKey:DocumentsFromKYCProfileId"`
     Screenings           []Screening `gorm:"foreignKey:ScreeningsFromKYCProfileId"`
     Addresses           []VerifiedAddress `gorm:"foreignKey:AddressesFromKYCProfileId"`
    Status                      KYCStatus
    VerificationLevel                      VerificationLevel

// parent associations as their child

}

