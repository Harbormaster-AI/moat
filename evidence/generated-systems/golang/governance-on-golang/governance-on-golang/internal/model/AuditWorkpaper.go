package model

import (
    "gorm.io/gorm"
)

//==============================================================
// AuditWorkpaper Declaration
//==============================================================
type AuditWorkpaper struct {
    gorm.Model
     WorkpaperRef                                    string
    Subject                                    string
    WorkpaperUrl                                                            string
    EngagementId         *uint
    Engagement           *AuditEngagement `gorm:"foreignKey:EngagementId"`
     Evidence           []Evidence `gorm:"foreignKey:EvidenceFromAuditWorkpaperId"`
     Findings           []AuditFinding `gorm:"foreignKey:FindingsFromAuditWorkpaperId"`

// parent associations as their child

}

