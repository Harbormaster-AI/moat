package model

import (
    "gorm.io/gorm"
)

//==============================================================
// RetentionSchedule Declaration
//==============================================================
type RetentionSchedule struct {
    gorm.Model
     Name                                    string
    RetentionPeriodMonths                                                            string
     Repositories           []RecordsRepository `gorm:"foreignKey:RepositoriesFromRetentionScheduleId"`
     Records           []Record_ `gorm:"foreignKey:RecordsFromRetentionScheduleId"`
     Exceptions           []Exception_ `gorm:"foreignKey:ExceptionsFromRetentionScheduleId"`
     DispositionReviews           []DispositionReview `gorm:"foreignKey:DispositionReviewsFromRetentionScheduleId"`
    RetentionTrigger                      RetentionTrigger
    DispositionAction                      DispositionAction
    Status                      RetentionStatus

// parent associations as their child

}

