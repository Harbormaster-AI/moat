package model

import (
    "time"
    "gorm.io/gorm"
)

//==============================================================
// Termination Declaration
//==============================================================
type Termination struct {
    gorm.Model
     TerminationNumber                                    string
    TerminationDate                                                            time.Time
    Notes                                    string
    EligibleForRehire                                    bool
    EmployeeId         *uint
    Employee           *Employee `gorm:"foreignKey:EmployeeId"`
    AssignmentId         *uint
    Assignment           *EmploymentAssignment `gorm:"foreignKey:AssignmentId"`
    Reason                      TerminationReason
    Type                      TerminationType

// parent associations as their child

}

