package model

import (
    "time"
    "gorm.io/gorm"
)

//==============================================================
// AuditEngagement Declaration
//==============================================================
type AuditEngagement struct {
    gorm.Model
     Title                                    string
    StartDate                                                            time.Time
    EndDate                                                            time.Time
    AuditProgramId         *uint
    AuditProgram           *AuditProgram `gorm:"foreignKey:AuditProgramId"`
     BusinessUnits           []BusinessUnit `gorm:"foreignKey:BusinessUnitsFromAuditEngagementId"`
     ControlTests           []ControlTest_ `gorm:"foreignKey:ControlTestsFromAuditEngagementId"`
     Workpapers           []AuditWorkpaper `gorm:"foreignKey:WorkpapersFromAuditEngagementId"`
     Findings           []AuditFinding `gorm:"foreignKey:FindingsFromAuditEngagementId"`
    Status                      AuditStatus

// parent associations as their child

}

