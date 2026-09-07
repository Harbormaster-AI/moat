package model

import (
    "gorm.io/gorm"
)

//==============================================================
// ModelVersion Declaration
//==============================================================
type ModelVersion struct {
    gorm.Model
     Version                                    string
    Model_Id         *uint
    Model_           *Model_ `gorm:"foreignKey:Model_Id"`
    TrainingRunId         *uint
    TrainingRun           *TrainingRun `gorm:"foreignKey:TrainingRunId"`
     EvaluationMetrics           []EvaluationMetric `gorm:"foreignKey:EvaluationMetricsFromModelVersionId"`
     Deployments           []InferenceEndpoint `gorm:"foreignKey:DeploymentsFromModelVersionId"`
     FeatureSets           []FeatureSet `gorm:"foreignKey:FeatureSetsFromModelVersionId"`
     Datasets           []DataSet `gorm:"foreignKey:DatasetsFromModelVersionId"`
    Lifecycle                      ModelLifecycle
    TrainingStatus                      TrainingStatus

// parent associations as their child

}

