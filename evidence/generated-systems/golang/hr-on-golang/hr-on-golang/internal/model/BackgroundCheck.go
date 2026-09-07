package model

import (
    "time"
    "gorm.io/gorm"
)

//==============================================================
// BackgroundCheck Declaration
//==============================================================
type BackgroundCheck struct {
    gorm.Model
     CheckNumber                                    string
    Provider                                    string
    CompletedDate                                                            time.Time
    CandidateId         *uint
    Candidate           *Candidate `gorm:"foreignKey:CandidateId"`
    RequisitionId         *uint
    Requisition           *JobRequisition `gorm:"foreignKey:RequisitionId"`
    ReportId         *uint
    Report           *Document `gorm:"foreignKey:ReportId"`
    Status                      BackgroundCheckStatus

// parent associations as their child

}

