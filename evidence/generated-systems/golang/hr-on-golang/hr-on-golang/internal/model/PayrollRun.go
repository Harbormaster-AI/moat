package model

import (
    "time"
    "gorm.io/gorm"
)

//==============================================================
// PayrollRun Declaration
//==============================================================
type PayrollRun struct {
    gorm.Model
     RunNumber                                    string
    PeriodStart                                                            time.Time
    PeriodEnd                                                            time.Time
    PaymentDate                                                            time.Time
    PayrollCalendarId         *uint
    PayrollCalendar           *PayrollCalendar `gorm:"foreignKey:PayrollCalendarId"`
     PayrollItems           []PayrollItem `gorm:"foreignKey:PayrollItemsFromPayrollRunId"`
    Status                      PayrollStatus

// parent associations as their child

}

