package model

import (
    "time"
    "gorm.io/gorm"
)

//==============================================================
// Screening Declaration
//==============================================================
type Screening struct {
    gorm.Model
     Name                                    string
    CompletedDate                                                            time.Time
    ApplicationId         *uint
    Application           *JobApplication `gorm:"foreignKey:ApplicationId"`
    Status                      BackgroundCheckStatus

// parent associations as their child

}

