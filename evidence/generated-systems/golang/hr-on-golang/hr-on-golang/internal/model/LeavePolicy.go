package model

import (
    "gorm.io/gorm"
)

//==============================================================
// LeavePolicy Declaration
//==============================================================
type LeavePolicy struct {
    gorm.Model
     Name                                    string
    AccrualRate                                                            string
    CarryoverAllowed                                    bool
    MaxBalance                                                            string
    OrganizationId         *uint
    Organization           *Organization `gorm:"foreignKey:OrganizationId"`
     LeaveRequests           []LeaveRequest `gorm:"foreignKey:LeaveRequestsFromLeavePolicyId"`
    LeaveCategory                      LeaveCategory
    AccrualUnit                      AccrualUnit

// parent associations as their child

}

