package model

import (
    "time"
    "gorm.io/gorm"
)

//==============================================================
// PolicyAcknowledgement Declaration
//==============================================================
type PolicyAcknowledgement struct {
    gorm.Model
     AcknowledgementDate                                                            time.Time
    PolicyId         *uint
    Policy           *Policy `gorm:"foreignKey:PolicyId"`
    EmployeeId         *uint
    Employee           *Employee `gorm:"foreignKey:EmployeeId"`
    Status                      AcknowledgementStatus

// parent associations as their child

}

