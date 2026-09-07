package model

import (
    "gorm.io/gorm"
)

//==============================================================
// RunParameter Declaration
//==============================================================
type RunParameter struct {
    gorm.Model
     Name                                    string
    Value                                    string
    TrainingRunId         *uint
    TrainingRun           *TrainingRun `gorm:"foreignKey:TrainingRunId"`

// parent associations as their child

}

