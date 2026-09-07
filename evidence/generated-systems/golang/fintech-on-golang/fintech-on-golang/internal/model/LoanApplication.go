package model

import (
    "time"
    "gorm.io/gorm"
)

//==============================================================
// LoanApplication Declaration
//==============================================================
type LoanApplication struct {
    gorm.Model
     ApplicationNumber                                    string
    AmountRequested                                                            string
    TermMonths                                                            string
    SubmittedAt                                                            time.Time
    CustomerId         *uint
    Customer           *Customer `gorm:"foreignKey:CustomerId"`
    RiskAssessmentId         *uint
    RiskAssessment           *RiskAssessment `gorm:"foreignKey:RiskAssessmentId"`
    LoanId         *uint
    Loan           *Loan `gorm:"foreignKey:LoanId"`
    Product                      LoanProductType
    Purpose                      LoanPurpose
    Status                      ApplicationStatus

// parent associations as their child

}

