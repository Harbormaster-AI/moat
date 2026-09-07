package model

import (
    "gorm.io/gorm"
)

//==============================================================
// AnalyticsWorkspace Declaration
//==============================================================
type AnalyticsWorkspace struct {
    gorm.Model
     Name                                    string
    BusinessDomain                                    string
    OwnerTeam                                    string
     Datasets           []DataSet `gorm:"foreignKey:DatasetsFromAnalyticsWorkspaceId"`
     DataSources           []DataSource `gorm:"foreignKey:DataSourcesFromAnalyticsWorkspaceId"`
     Pipelines           []DataPipeline `gorm:"foreignKey:PipelinesFromAnalyticsWorkspaceId"`
     Dashboards           []Dashboard `gorm:"foreignKey:DashboardsFromAnalyticsWorkspaceId"`
     Reports           []Report `gorm:"foreignKey:ReportsFromAnalyticsWorkspaceId"`
     Notebooks           []Notebook `gorm:"foreignKey:NotebooksFromAnalyticsWorkspaceId"`
     Models           []Model_ `gorm:"foreignKey:ModelsFromAnalyticsWorkspaceId"`
     FeatureSets           []FeatureSet `gorm:"foreignKey:FeatureSetsFromAnalyticsWorkspaceId"`
     Policies           []AccessPolicy `gorm:"foreignKey:PoliciesFromAnalyticsWorkspaceId"`
     LineageNodes           []LineageNode `gorm:"foreignKey:LineageNodesFromAnalyticsWorkspaceId"`
    GovernanceTier                      GovernanceTier

// parent associations as their child

}

