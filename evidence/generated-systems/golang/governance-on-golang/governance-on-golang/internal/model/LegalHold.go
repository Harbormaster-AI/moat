package model

import (
    "time"
    "gorm.io/gorm"
)

//==============================================================
// LegalHold Declaration
//==============================================================
type LegalHold struct {
    gorm.Model
     Name                                    string
    Reason                                    string
    IssuedDate                                                            time.Time
    ReleaseDate                                                            time.Time
     Repositories           []RecordsRepository `gorm:"foreignKey:RepositoriesFromLegalHoldId"`
     Records           []Record_ `gorm:"foreignKey:RecordsFromLegalHoldId"`
    MatterId         *uint
    Matter           *Matter `gorm:"foreignKey:MatterId"`
    HoldStatus                      LegalHoldStatus

// parent associations as their child

}

