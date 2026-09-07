package model

import (
    "time"
    "gorm.io/gorm"
)

//==============================================================
// Issue Declaration
//==============================================================
type Issue struct {
    gorm.Model
     Title                                    string
    OpenedDate                                                            time.Time
    ClosedDate                                                            time.Time
    RiskId         *uint
    Risk           *Risk `gorm:"foreignKey:RiskId"`
    FindingId         *uint
    Finding           *AuditFinding `gorm:"foreignKey:FindingId"`
     CorrectiveActions           []CorrectiveAction `gorm:"foreignKey:CorrectiveActionsFromIssueId"`
    ControlId         *uint
    Control           *Control `gorm:"foreignKey:ControlId"`
    IssueType                      IssueType
    Priority                      Priority
    Status                      IssueStatus

// parent associations as their child

}

