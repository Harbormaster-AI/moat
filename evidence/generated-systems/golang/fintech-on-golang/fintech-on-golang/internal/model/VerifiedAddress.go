package model

import (
    "time"
    "gorm.io/gorm"
)

//==============================================================
// VerifiedAddress Declaration
//==============================================================
type VerifiedAddress struct {
    gorm.Model
     Address                                                            string
    VerifiedAt                                                            time.Time
    KycProfileId         *uint
    KycProfile           *KYCProfile `gorm:"foreignKey:KycProfileId"`
    VerificationStatus                      VerificationStatus

// parent associations as their child

}

