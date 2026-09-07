package model

import (
    "time"
    "gorm.io/gorm"
)

//==============================================================
// LoanTransaction Declaration
//==============================================================
type LoanTransaction struct {
    gorm.Model
     TransactionId                                                            string
    Amount                                                            string
    PostingDate                                                            time.Time
    LoanId         *uint
    Loan           *Loan `gorm:"foreignKey:LoanId"`
    Type                      LoanTransactionType
    Status                      PostingStatus

// parent associations as their child

}

