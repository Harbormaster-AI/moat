package model

import (
    "gorm.io/gorm"
)

//==============================================================
// Shift Declaration
//==============================================================
type Shift struct {
    gorm.Model
     ShiftName                                    string
    StartTime                                    string
    EndTime                                    string
    PlantId         *uint
    Plant           *Plant `gorm:"foreignKey:PlantId"`
     Assignments           []ShiftAssignment `gorm:"foreignKey:AssignmentsFromShiftId"`
    ShiftType                      ShiftType

// parent associations as their child

}

