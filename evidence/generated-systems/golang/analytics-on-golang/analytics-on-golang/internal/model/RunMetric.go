package model

import (
    "gorm.io/gorm"
)

//==============================================================
// RunMetric Declaration
//==============================================================
type RunMetric struct {
    gorm.Model
     Name                                    string
    Value                                                            string
    TrainingRunId         *uint
    TrainingRun           *TrainingRun `gorm:"foreignKey:TrainingRunId"`
    MetricId         *uint
    Metric           *Metric `gorm:"foreignKey:MetricId"`
    DatasetId         *uint
    Dataset           *DataSet `gorm:"foreignKey:DatasetId"`

// parent associations as their child

}

