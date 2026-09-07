package model

import (
    "time"
    "gorm.io/gorm"
)

//==============================================================
// PerformanceReview Declaration
//==============================================================
type PerformanceReview struct {
    gorm.Model
     ReviewNumber                                    string
    ReviewDate                                                            time.Time
    ReviewerComments                                    string
    EmployeeId         *uint
    Employee           *Employee `gorm:"foreignKey:EmployeeId"`
    ReviewerId         *uint
    Reviewer           *Employee `gorm:"foreignKey:ReviewerId"`
    CycleId         *uint
    Cycle           *PerformanceCycle `gorm:"foreignKey:CycleId"`
     CompetencyRatings           []CompetencyRating `gorm:"foreignKey:CompetencyRatingsFromPerformanceReviewId"`
     Goals           []Goal `gorm:"foreignKey:GoalsFromPerformanceReviewId"`
    Rating                      PerformanceRating
    Status                      ReviewStatus

// parent associations as their child

}

