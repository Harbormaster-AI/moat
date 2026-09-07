package model

import (
    "time"
    "gorm.io/gorm"
)

//==============================================================
// CareTask Declaration
//==============================================================
type CareTask struct {
    gorm.Model
     Description                                    string
    DueDate                                                            time.Time
    CarePlanId         *uint
    CarePlan           *CarePlan `gorm:"foreignKey:CarePlanId"`
    AssignedToId         *uint
    AssignedTo           *Clinician `gorm:"foreignKey:AssignedToId"`
    EncounterId         *uint
    Encounter           *Encounter `gorm:"foreignKey:EncounterId"`
    Status                      TaskStatus
    Priority                      Priority

// parent associations as their child

}

