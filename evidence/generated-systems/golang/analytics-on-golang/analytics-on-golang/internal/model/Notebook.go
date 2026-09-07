package model

import (
    "gorm.io/gorm"
)

//==============================================================
// Notebook Declaration
//==============================================================
type Notebook struct {
    gorm.Model
     Title                                    string
    Repository                                                            string
    WorkspaceId         *uint
    Workspace           *AnalyticsWorkspace `gorm:"foreignKey:WorkspaceId"`
     Datasets           []DataSet `gorm:"foreignKey:DatasetsFromNotebookId"`
     Experiments           []Experiment `gorm:"foreignKey:ExperimentsFromNotebookId"`
     Queries           []BIQuery `gorm:"foreignKey:QueriesFromNotebookId"`
    Language                      NotebookLanguage

// parent associations as their child

}

