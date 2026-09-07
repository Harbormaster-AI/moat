package model

import (
    "time"
    "gorm.io/gorm"
)

//==============================================================
// ComplianceAlert Declaration
//==============================================================
type ComplianceAlert struct {
    gorm.Model
     AlertCode                                    string
    RaisedAt                                                            time.Time
    Notes                                    string
    ScreeningId         *uint
    Screening           *Screening `gorm:"foreignKey:ScreeningId"`
    TransactionId         *uint
    Transaction           *Transaction `gorm:"foreignKey:TransactionId"`
    Severity                      AlertSeverity
    Status                      AlertStatus

// parent associations as their child

}

