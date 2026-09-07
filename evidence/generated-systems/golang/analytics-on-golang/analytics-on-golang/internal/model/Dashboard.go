package model

import (
    "gorm.io/gorm"
)

//==============================================================
// Dashboard Declaration
//==============================================================
type Dashboard struct {
    gorm.Model
     Title                                    string
    Theme                                    string
    WorkspaceId         *uint
    Workspace           *AnalyticsWorkspace `gorm:"foreignKey:WorkspaceId"`
     Visualizations           []Visualization `gorm:"foreignKey:VisualizationsFromDashboardId"`
     Reports           []Report `gorm:"foreignKey:ReportsFromDashboardId"`
     Datasets           []DataSet `gorm:"foreignKey:DatasetsFromDashboardId"`
     Alerts           []Alert `gorm:"foreignKey:AlertsFromDashboardId"`
     Queries           []BIQuery `gorm:"foreignKey:QueriesFromDashboardId"`
     Tags           []Tag `gorm:"foreignKey:TagsFromDashboardId"`
    Status                      DashboardStatus

// parent associations as their child

}

