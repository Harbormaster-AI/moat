package model

import (
    "time"
    "gorm.io/gorm"
)

//==============================================================
// Offer Declaration
//==============================================================
type Offer struct {
    gorm.Model
     OfferNumber                                    string
    ProposedStartDate                                                            time.Time
    BaseSalary                                                            string
    SignOnBonus                                                            string
    RequisitionId         *uint
    Requisition           *JobRequisition `gorm:"foreignKey:RequisitionId"`
    CandidateId         *uint
    Candidate           *Candidate `gorm:"foreignKey:CandidateId"`
    ApprovedById         *uint
    ApprovedBy           *Employee `gorm:"foreignKey:ApprovedById"`
    ContractId         *uint
    Contract           *EmploymentContract `gorm:"foreignKey:ContractId"`
    Status                      OfferStatus

// parent associations as their child

}

