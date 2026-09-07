package model

import (
    "gorm.io/gorm"
)

//==============================================================
// TaxWithholding Declaration
//==============================================================
type TaxWithholding struct {
    gorm.Model
     TaxId                                                            string
    Allowances                                                            string
    AdditionalAmount                                                            string
    EmployeeId         *uint
    Employee           *Employee `gorm:"foreignKey:EmployeeId"`
    FilingStatus                      FilingStatus

// parent associations as their child

}

