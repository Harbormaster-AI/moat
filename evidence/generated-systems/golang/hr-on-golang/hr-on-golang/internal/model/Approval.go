package model

import (
    "time"
    "gorm.io/gorm"
)

//==============================================================
// Approval Declaration
//==============================================================
type Approval struct {
    gorm.Model
     ApproverComment                                    string
    ActionDate                                                            time.Time
    ApproverId         *uint
    Approver           *Employee `gorm:"foreignKey:ApproverId"`
    TimesheetId         *uint
    Timesheet           *Timesheet `gorm:"foreignKey:TimesheetId"`
    LeaveRequestId         *uint
    LeaveRequest           *LeaveRequest `gorm:"foreignKey:LeaveRequestId"`
    Status                      ApprovalStatus

// parent associations as their child

}

