package model

import (
    "time"
    "gorm.io/gorm"
)

//==============================================================
// JobRequisition Declaration
//==============================================================
type JobRequisition struct {
    gorm.Model
     RequisitionNumber                                    string
    Title                                    string
    Openings                                                            string
    TargetStartDate                                                            time.Time
    DepartmentId         *uint
    Department           *Department `gorm:"foreignKey:DepartmentId"`
    HiringManagerId         *uint
    HiringManager           *Employee `gorm:"foreignKey:HiringManagerId"`
    RecruiterId         *uint
    Recruiter           *Employee `gorm:"foreignKey:RecruiterId"`
    JobProfileId         *uint
    JobProfile           *JobProfile `gorm:"foreignKey:JobProfileId"`
     Candidates           []Candidate `gorm:"foreignKey:CandidatesFromJobRequisitionId"`
     Interviews           []Interview `gorm:"foreignKey:InterviewsFromJobRequisitionId"`
     Offers           []Offer `gorm:"foreignKey:OffersFromJobRequisitionId"`
    Status                      RequisitionStatus
    Priority                      RequisitionPriority

// parent associations as their child

}

