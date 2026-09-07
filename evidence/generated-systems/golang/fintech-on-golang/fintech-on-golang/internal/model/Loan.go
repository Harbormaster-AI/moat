package model

import (
    "time"
    "gorm.io/gorm"
)

//==============================================================
// Loan Declaration
//==============================================================
type Loan struct {
    gorm.Model
     LoanNumber                                    string
    Principal                                                            string
    InterestRate                                                            string
    OriginationDate                                                            time.Time
    MaturityDate                                                            time.Time
    CustomerId         *uint
    Customer           *Customer `gorm:"foreignKey:CustomerId"`
     Schedule           []RepaymentSchedule `gorm:"foreignKey:ScheduleFromLoanId"`
     Collateral           []Collateral `gorm:"foreignKey:CollateralFromLoanId"`
     Transactions           []LoanTransaction `gorm:"foreignKey:TransactionsFromLoanId"`
    RateType                      InterestRateType
    Status                      LoanStatus

// parent associations as their child

}

