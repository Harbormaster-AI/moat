package model

import (
    "time"
    "gorm.io/gorm"
)

//==============================================================
// AuditFinding Declaration
//==============================================================
type AuditFinding struct {
    gorm.Model
     Title                                    string
    Description                                    string
    DueDate                                                            time.Time
    EngagementId         *uint
    Engagement           *AuditEngagement `gorm:"foreignKey:EngagementId"`
    WorkpaperId         *uint
    Workpaper           *AuditWorkpaper `gorm:"foreignKey:WorkpaperId"`
     CorrectiveActions           []CorrectiveAction `gorm:"foreignKey:CorrectiveActionsFromAuditFindingId"`
     RelatedRisks           []Risk `gorm:"foreignKey:RelatedRisksFromAuditFindingId"`
     RelatedControls           []Control `gorm:"foreignKey:RelatedControlsFromAuditFindingId"`
     Issues           []Issue `gorm:"foreignKey:IssuesFromAuditFindingId"`
    Severity                      FindingSeverity
    Status                      FindingStatus

// parent associations as their child

}

