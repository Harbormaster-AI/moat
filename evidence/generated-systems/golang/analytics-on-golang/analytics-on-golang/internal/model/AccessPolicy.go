package model

import (
    "gorm.io/gorm"
)

//==============================================================
// AccessPolicy Declaration
//==============================================================
type AccessPolicy struct {
    gorm.Model
     Name                                    string
    SubjectName                                    string
    WorkspaceId         *uint
    Workspace           *AnalyticsWorkspace `gorm:"foreignKey:WorkspaceId"`
     Datasets           []DataSet `gorm:"foreignKey:DatasetsFromAccessPolicyId"`
     Dashboards           []Dashboard `gorm:"foreignKey:DashboardsFromAccessPolicyId"`
     Reports           []Report `gorm:"foreignKey:ReportsFromAccessPolicyId"`
     Models           []Model_ `gorm:"foreignKey:ModelsFromAccessPolicyId"`
     FeatureSets           []FeatureSet `gorm:"foreignKey:FeatureSetsFromAccessPolicyId"`
    AccessLevel                      AccessLevel
    SubjectType                      SubjectType

// parent associations as their child

}

