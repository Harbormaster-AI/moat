package model

import (
    "time"
    "gorm.io/gorm"
)

//==============================================================
// Goal Declaration
//==============================================================
type Goal struct {
    gorm.Model
     Title                                    string
    Description                                    string
    TargetDate                                                            time.Time
    Weight                                                            string
    EmployeeId         *uint
    Employee           *Employee `gorm:"foreignKey:EmployeeId"`
    CycleId         *uint
    Cycle           *PerformanceCycle `gorm:"foreignKey:CycleId"`
    ParentGoalId         *uint
    ParentGoal           *Goal `gorm:"foreignKey:ParentGoalId"`
     ChildGoals           []Goal `gorm:"foreignKey:ChildGoalsFromGoalId"`
    Status                      GoalStatus

// parent associations as their child

}

