package model

import (
    "time"
    "gorm.io/gorm"
)

//==============================================================
// AccountStatement Declaration
//==============================================================
type AccountStatement struct {
    gorm.Model
     StatementNumber                                    string
    PeriodStart                                                            time.Time
    PeriodEnd                                                            time.Time
    OpeningBalance                                                            string
    ClosingBalance                                                            string
    GeneratedAt                                                            time.Time
    AccountId         *uint
    Account           *Account `gorm:"foreignKey:AccountId"`

// parent associations as their child

}

