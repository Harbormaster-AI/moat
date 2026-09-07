package model

import (
    "time"
    "gorm.io/gorm"
)

//==============================================================
// OnboardingTask Declaration
//==============================================================
type OnboardingTask struct {
    gorm.Model
     TaskNumber                                    string
    Name                                    string
    DueDate                                                            time.Time
    EmployeeId         *uint
    Employee           *Employee `gorm:"foreignKey:EmployeeId"`
    AssignedToId         *uint
    AssignedTo           *Employee `gorm:"foreignKey:AssignedToId"`
     Dependencies           []OnboardingTask `gorm:"foreignKey:DependenciesFromOnboardingTaskId"`
    RelatedOfferId         *uint
    RelatedOffer           *Offer `gorm:"foreignKey:RelatedOfferId"`
    Status                      OnboardingTaskStatus

// parent associations as their child

}

