package model

import (
    "time"
    "gorm.io/gorm"
)

//==============================================================
// KYCDocument Declaration
//==============================================================
type KYCDocument struct {
    gorm.Model
     Reference                                                            string
    IssuedCountry                                    string
    ExpirationDate                                                            time.Time
    KycProfileId         *uint
    KycProfile           *KYCProfile `gorm:"foreignKey:KycProfileId"`
    DocumentType                      KYCDocumentType
    Status                      DocumentStatus

// parent associations as their child

}

