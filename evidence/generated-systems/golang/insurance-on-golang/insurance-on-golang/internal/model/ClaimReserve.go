package model

import (
    "time"
    "gorm.io/gorm"
)

//==============================================================
// ClaimReserve Declaration
//==============================================================
type ClaimReserve struct {
    gorm.Model
     Amount                                                            string
    SetDate                                                            time.Time
    ClaimId         *uint
    Claim           *Claim `gorm:"foreignKey:ClaimId"`
    ExposureId         *uint
    Exposure           *Exposure `gorm:"foreignKey:ExposureId"`
    ReserveType                      ReserveType
    Status                      ReserveStatus

// parent associations as their child

}

