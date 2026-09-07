package model

import (
    "time"
    "gorm.io/gorm"
)

//==============================================================
// Interview Declaration
//==============================================================
type Interview struct {
    gorm.Model
     InterviewDate                                                            time.Time
    Feedback                                    string
    RequisitionId         *uint
    Requisition           *JobRequisition `gorm:"foreignKey:RequisitionId"`
    CandidateId         *uint
    Candidate           *Candidate `gorm:"foreignKey:CandidateId"`
     Interviewers           []Employee `gorm:"foreignKey:InterviewersFromInterviewId"`
    Stage                      InterviewStage
    Result                      InterviewResult

// parent associations as their child

}

