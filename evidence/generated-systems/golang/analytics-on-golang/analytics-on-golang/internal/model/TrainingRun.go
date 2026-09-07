package model

import (
    "time"
    "gorm.io/gorm"
)

//==============================================================
// TrainingRun Declaration
//==============================================================
type TrainingRun struct {
    gorm.Model
     RunLabel                                    string
    StartedAt                                                            time.Time
    CompletedAt                                                            time.Time
    ExperimentId         *uint
    Experiment           *Experiment `gorm:"foreignKey:ExperimentId"`
    ModelVersionId         *uint
    ModelVersion           *ModelVersion `gorm:"foreignKey:ModelVersionId"`
     InputDatasets           []DataSet `gorm:"foreignKey:InputDatasetsFromTrainingRunId"`
     Features           []Feature `gorm:"foreignKey:FeaturesFromTrainingRunId"`
     RunMetrics           []RunMetric `gorm:"foreignKey:RunMetricsFromTrainingRunId"`
     RunParameters           []RunParameter `gorm:"foreignKey:RunParametersFromTrainingRunId"`
    Status                      TrainingStatus

// parent associations as their child

}

