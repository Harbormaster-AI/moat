package model

import (
    "time"
    "gorm.io/gorm"
)

//==============================================================
// RepaymentSchedule Declaration
//==============================================================
type RepaymentSchedule struct {
    gorm.Model
     InstallmentNumber                                                            string
    DueDate                                                            time.Time
    AmountDue                                                            string
    PrincipalDue                                                            string
    InterestDue                                                            string
    LoanId         *uint
    Loan           *Loan `gorm:"foreignKey:LoanId"`
     Payments           []Transaction `gorm:"foreignKey:PaymentsFromRepaymentScheduleId"`
    Status                      InstallmentStatus

// parent associations as their child

}

