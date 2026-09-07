package model

import (
    "gorm.io/gorm"
)

//==============================================================
// Collateral Declaration
//==============================================================
type Collateral struct {
    gorm.Model
     Description                                    string
    Value                                                            string
    LoanId         *uint
    Loan           *Loan `gorm:"foreignKey:LoanId"`
    CollateralType                      CollateralType

// parent associations as their child

}

