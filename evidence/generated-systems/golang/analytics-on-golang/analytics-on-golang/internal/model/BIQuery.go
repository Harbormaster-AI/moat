package model

import (
    "gorm.io/gorm"
)

//==============================================================
// BIQuery Declaration
//==============================================================
type BIQuery struct {
    gorm.Model
     Name                                    string
    Text                                    string
    WorkspaceId         *uint
    Workspace           *AnalyticsWorkspace `gorm:"foreignKey:WorkspaceId"`
     Datasets           []DataSet `gorm:"foreignKey:DatasetsFromBIQueryId"`
     Reports           []Report `gorm:"foreignKey:ReportsFromBIQueryId"`
     Dashboards           []Dashboard `gorm:"foreignKey:DashboardsFromBIQueryId"`
     Notebooks           []Notebook `gorm:"foreignKey:NotebooksFromBIQueryId"`
    Dialect                      SQLDialect

// parent associations as their child

}

