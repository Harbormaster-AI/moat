package model

import (
    "time"
    "gorm.io/gorm"
)

//==============================================================
// ControlTest_ Declaration
//==============================================================
type ControlTest_ struct {
    gorm.Model
     Name                                    string
    TestPeriodStart                                                            time.Time
    TestPeriodEnd                                                            time.Time
    SampleSize                                                            string
    ControlId         *uint
    Control           *Control `gorm:"foreignKey:ControlId"`
     Evidence           []Evidence `gorm:"foreignKey:EvidenceFromControlTest_Id"`
    EngagementId         *uint
    Engagement           *AuditEngagement `gorm:"foreignKey:EngagementId"`
    TestType                      TestType
    Effectiveness                      ControlEffectiveness
    Status                      TestStatus

// parent associations as their child

}

