package model

import (
    "time"
    "gorm.io/gorm"
)

//==============================================================
// DispositionReview Declaration
//==============================================================
type DispositionReview struct {
    gorm.Model
     ReviewDate                                                            time.Time
    Reviewer                                    string
    Notes                                    string
    RecordId         *uint
    Record           *Record_ `gorm:"foreignKey:RecordId"`
    RetentionScheduleId         *uint
    RetentionSchedule           *RetentionSchedule `gorm:"foreignKey:RetentionScheduleId"`
    Outcome                      DispositionOutcome

// parent associations as their child

}

