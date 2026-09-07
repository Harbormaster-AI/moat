package model

import (
    "time"
    "gorm.io/gorm"
)

//==============================================================
// Timesheet Declaration
//==============================================================
type Timesheet struct {
    gorm.Model
     PeriodStart                                                            time.Time
    PeriodEnd                                                            time.Time
    SubmissionDate                                                            time.Time
    EmployeeId         *uint
    Employee           *Employee `gorm:"foreignKey:EmployeeId"`
     TimeEntries           []TimeEntry `gorm:"foreignKey:TimeEntriesFromTimesheetId"`
     Approvals           []Approval `gorm:"foreignKey:ApprovalsFromTimesheetId"`
    Status                      TimesheetStatus

// parent associations as their child

}

