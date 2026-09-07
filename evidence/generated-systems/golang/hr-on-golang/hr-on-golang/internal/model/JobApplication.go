package model

import (
    "time"
    "gorm.io/gorm"
)

//==============================================================
// JobApplication Declaration
//==============================================================
type JobApplication struct {
    gorm.Model
     ApplicationNumber                                    string
    AppliedDate                                                            time.Time
    ResumeUrl                                    string
    CandidateId         *uint
    Candidate           *Candidate `gorm:"foreignKey:CandidateId"`
    RequisitionId         *uint
    Requisition           *JobRequisition `gorm:"foreignKey:RequisitionId"`
     Screenings           []Screening `gorm:"foreignKey:ScreeningsFromJobApplicationId"`
    Status                      ApplicationStatus

// parent associations as their child

}

