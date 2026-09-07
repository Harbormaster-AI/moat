package model

import (
    "time"
    "gorm.io/gorm"
)

//==============================================================
// CorrectiveAction Declaration
//==============================================================
type CorrectiveAction struct {
    gorm.Model
     ActionTitle                                    string
    Owner                                    string
    TargetDate                                                            time.Time
    FindingId         *uint
    Finding           *AuditFinding `gorm:"foreignKey:FindingId"`
    IssueId         *uint
    Issue           *Issue `gorm:"foreignKey:IssueId"`
    Status                      ActionStatus

// parent associations as their child

}

