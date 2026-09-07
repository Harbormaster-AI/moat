package model

import (
    "time"
    "gorm.io/gorm"
)

//==============================================================
// CorrectiveAction Declaration
//==============================================================
type CorrectiveAction struct {
    gorm.Model
     CapaNumber                                    string
    RootCause                                    string
    CorrectiveAction                                    string
    VerificationDate                                                            time.Time
    NonconformanceId         *uint
    Nonconformance           *Nonconformance `gorm:"foreignKey:NonconformanceId"`
    OwnerId         *uint
    Owner           *Employee `gorm:"foreignKey:OwnerId"`
    Status                      CAPAStatus

// parent associations as their child

}

