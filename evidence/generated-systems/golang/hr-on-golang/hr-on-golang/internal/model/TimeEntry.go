package model

import (
    "time"
    "gorm.io/gorm"
)

//==============================================================
// TimeEntry Declaration
//==============================================================
type TimeEntry struct {
    gorm.Model
     EntryDate                                                            time.Time
    HoursWorked                                                            string
    TimesheetId         *uint
    Timesheet           *Timesheet `gorm:"foreignKey:TimesheetId"`
    EmployeeId         *uint
    Employee           *Employee `gorm:"foreignKey:EmployeeId"`
    CostCenterId         *uint
    CostCenter           *CostCenter `gorm:"foreignKey:CostCenterId"`
    EntryType                      TimeEntryType

// parent associations as their child

}

