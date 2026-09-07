package model

import (
    "time"
    "gorm.io/gorm"
)

//==============================================================
// Exception_ Declaration
//==============================================================
type Exception_ struct {
    gorm.Model
     Title                                    string
    Justification                                    string
    StartDate                                                            time.Time
    EndDate                                                            time.Time
    RetentionScheduleId         *uint
    RetentionSchedule           *RetentionSchedule `gorm:"foreignKey:RetentionScheduleId"`
    PolicyId         *uint
    Policy           *Policy `gorm:"foreignKey:PolicyId"`
    ControlId         *uint
    Control           *Control `gorm:"foreignKey:ControlId"`
    RiskId         *uint
    Risk           *Risk `gorm:"foreignKey:RiskId"`
    ExceptionType                      ExceptionType
    Status                      ExceptionStatus

// parent associations as their child

}

