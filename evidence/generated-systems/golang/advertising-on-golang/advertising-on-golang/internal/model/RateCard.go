package model

import (
    "time"
    "gorm.io/gorm"
)

//==============================================================
// RateCard Declaration
//==============================================================
type RateCard struct {
    gorm.Model
     Name                                    string
    EffectiveDate                                                            time.Time
    Currency                                    string
    PublisherId         *uint
    Publisher           *Publisher `gorm:"foreignKey:PublisherId"`
     Rates           []Rate `gorm:"foreignKey:RatesFromRateCardId"`

// parent associations as their child

}

