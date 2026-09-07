package model

import (
    "gorm.io/gorm"
)

//==============================================================
// PayrollItem Declaration
//==============================================================
type PayrollItem struct {
    gorm.Model
     Amount                                                            string
    Taxable                                    bool
    PayrollRunId         *uint
    PayrollRun           *PayrollRun `gorm:"foreignKey:PayrollRunId"`
    EmployeeId         *uint
    Employee           *Employee `gorm:"foreignKey:EmployeeId"`
    ItemType                      PayrollItemType

// parent associations as their child

}

