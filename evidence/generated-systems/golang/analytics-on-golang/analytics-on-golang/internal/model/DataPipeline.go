package model

import (
    "gorm.io/gorm"
)

//==============================================================
// DataPipeline Declaration
//==============================================================
type DataPipeline struct {
    gorm.Model
     Name                                    string
    Schedule                                                            string
    WorkspaceId         *uint
    Workspace           *AnalyticsWorkspace `gorm:"foreignKey:WorkspaceId"`
     Tasks           []DataTask `gorm:"foreignKey:TasksFromDataPipelineId"`
     Sources           []DataSource `gorm:"foreignKey:SourcesFromDataPipelineId"`
     Outputs           []DataSet `gorm:"foreignKey:OutputsFromDataPipelineId"`
    LineageNodeId         *uint
    LineageNode           *LineageNode `gorm:"foreignKey:LineageNodeId"`
    TriggerType                      PipelineTriggerType
    Status                      PipelineStatus

// parent associations as their child

}

