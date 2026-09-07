package model

import (
    "gorm.io/gorm"
)

//==============================================================
// DataSet Declaration
//==============================================================
type DataSet struct {
    gorm.Model
     Name                                    string
    SchemaVersion                                    string
    RefreshSchedule                                                            string
    Sensitive                                    bool
    WorkspaceId         *uint
    Workspace           *AnalyticsWorkspace `gorm:"foreignKey:WorkspaceId"`
     Sources           []DataSource `gorm:"foreignKey:SourcesFromDataSetId"`
     Pipelines           []DataPipeline `gorm:"foreignKey:PipelinesFromDataSetId"`
     SemanticModels           []SemanticModel `gorm:"foreignKey:SemanticModelsFromDataSetId"`
     Dimensions           []Dimension `gorm:"foreignKey:DimensionsFromDataSetId"`
     Measures           []Measure `gorm:"foreignKey:MeasuresFromDataSetId"`
     Metrics           []Metric `gorm:"foreignKey:MetricsFromDataSetId"`
     QualityRules           []QualityRule `gorm:"foreignKey:QualityRulesFromDataSetId"`
    LineageNodeId         *uint
    LineageNode           *LineageNode `gorm:"foreignKey:LineageNodeId"`
     Tags           []Tag `gorm:"foreignKey:TagsFromDataSetId"`
    DataFormat                      DataFormat

// parent associations as their child

}

