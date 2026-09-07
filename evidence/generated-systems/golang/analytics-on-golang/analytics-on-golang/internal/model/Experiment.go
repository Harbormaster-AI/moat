package model

import (
    "gorm.io/gorm"
)

//==============================================================
// Experiment Declaration
//==============================================================
type Experiment struct {
    gorm.Model
     Name                                    string
    Objective                                    string
    WorkspaceId         *uint
    Workspace           *AnalyticsWorkspace `gorm:"foreignKey:WorkspaceId"`
     TrainingRuns           []TrainingRun `gorm:"foreignKey:TrainingRunsFromExperimentId"`
     Models           []Model_ `gorm:"foreignKey:ModelsFromExperimentId"`
     Notebooks           []Notebook `gorm:"foreignKey:NotebooksFromExperimentId"`
    Status                      ExperimentStatus

// parent associations as their child

}

