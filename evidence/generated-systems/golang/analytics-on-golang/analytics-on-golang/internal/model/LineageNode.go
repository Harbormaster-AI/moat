package model

import (
    "gorm.io/gorm"
)

//==============================================================
// LineageNode Declaration
//==============================================================
type LineageNode struct {
    gorm.Model
     Name                                    string
    QualifiedName                                    string
    WorkspaceId         *uint
    Workspace           *AnalyticsWorkspace `gorm:"foreignKey:WorkspaceId"`
     Inputs           []LineageNode `gorm:"foreignKey:InputsFromLineageNodeId"`
     Outputs           []LineageNode `gorm:"foreignKey:OutputsFromLineageNodeId"`
     Datasets           []DataSet `gorm:"foreignKey:DatasetsFromLineageNodeId"`
     Models           []Model_ `gorm:"foreignKey:ModelsFromLineageNodeId"`
     Pipelines           []DataPipeline `gorm:"foreignKey:PipelinesFromLineageNodeId"`
     Dashboards           []Dashboard `gorm:"foreignKey:DashboardsFromLineageNodeId"`
     Reports           []Report `gorm:"foreignKey:ReportsFromLineageNodeId"`
    NodeType                      LineageNodeType

// parent associations as their child

}

