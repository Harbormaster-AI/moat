package model

import (
    "gorm.io/gorm"
)

//==============================================================
// Report Declaration
//==============================================================
type Report struct {
    gorm.Model
     Title                                    string
    Audience                                    string
    WorkspaceId         *uint
    Workspace           *AnalyticsWorkspace `gorm:"foreignKey:WorkspaceId"`
     Visualizations           []Visualization `gorm:"foreignKey:VisualizationsFromReportId"`
     Datasets           []DataSet `gorm:"foreignKey:DatasetsFromReportId"`
     SemanticModels           []SemanticModel `gorm:"foreignKey:SemanticModelsFromReportId"`
     Queries           []BIQuery `gorm:"foreignKey:QueriesFromReportId"`
     Tags           []Tag `gorm:"foreignKey:TagsFromReportId"`
    Status                      ReportStatus

// parent associations as their child

}

