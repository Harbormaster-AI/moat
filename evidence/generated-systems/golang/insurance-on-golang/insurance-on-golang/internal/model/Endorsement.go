package model

import (
    "time"
    "gorm.io/gorm"
)

//==============================================================
// Endorsement Declaration
//==============================================================
type Endorsement struct {
    gorm.Model
     EndorsementNumber                                    string
    EffectiveDate                                                            time.Time
    Description                                    string
    PolicyId         *uint
    Policy           *Policy `gorm:"foreignKey:PolicyId"`

// parent associations as their child

}

