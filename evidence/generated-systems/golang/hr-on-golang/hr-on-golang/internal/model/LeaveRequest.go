package model

import (
    "time"
    "gorm.io/gorm"
)

//==============================================================
// LeaveRequest Declaration
//==============================================================
type LeaveRequest struct {
    gorm.Model
     RequestNumber                                    string
    StartDate                                                            time.Time
    EndDate                                                            time.Time
    Reason                                    string
    Hours                                                            string
    EmployeeId         *uint
    Employee           *Employee `gorm:"foreignKey:EmployeeId"`
    LeavePolicyId         *uint
    LeavePolicy           *LeavePolicy `gorm:"foreignKey:LeavePolicyId"`
     Approvals           []Approval `gorm:"foreignKey:ApprovalsFromLeaveRequestId"`
    Status                      LeaveStatus

// parent associations as their child

}

