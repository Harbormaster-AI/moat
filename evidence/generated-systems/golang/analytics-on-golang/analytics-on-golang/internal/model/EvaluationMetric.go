package model

import (
    "gorm.io/gorm"
)

//==============================================================
// EvaluationMetric Declaration
//==============================================================
type EvaluationMetric struct {
    gorm.Model
     Name                                    string
    Value                                                            string
    ModelVersionId         *uint
    ModelVersion           *ModelVersion `gorm:"foreignKey:ModelVersionId"`
    MetricId         *uint
    Metric           *Metric `gorm:"foreignKey:MetricId"`
    DatasetId         *uint
    Dataset           *DataSet `gorm:"foreignKey:DatasetId"`

// parent associations as their child

}

