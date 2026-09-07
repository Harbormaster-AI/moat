package model

import (
    "time"
    "gorm.io/gorm"
)

//==============================================================
// PerformanceCycle Declaration
//==============================================================
type PerformanceCycle struct {
    gorm.Model
     Name                                    string
    StartDate                                                            time.Time
    EndDate                                                            time.Time
    OrganizationId         *uint
    Organization           *Organization `gorm:"foreignKey:OrganizationId"`
     Reviews           []PerformanceReview `gorm:"foreignKey:ReviewsFromPerformanceCycleId"`
     Goals           []Goal `gorm:"foreignKey:GoalsFromPerformanceCycleId"`
    Status                      CycleStatus

// parent associations as their child

}

