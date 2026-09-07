package model

import (
    "time"
    "gorm.io/gorm"
)

//==============================================================
// Prediction Declaration
//==============================================================
type Prediction struct {
    gorm.Model
     ReferenceKey                                    string
    PredictedAt                                                            time.Time
    Score                                                            string
    EndpointId         *uint
    Endpoint           *InferenceEndpoint `gorm:"foreignKey:EndpointId"`
    ModelVersionId         *uint
    ModelVersion           *ModelVersion `gorm:"foreignKey:ModelVersionId"`
    DatasetId         *uint
    Dataset           *DataSet `gorm:"foreignKey:DatasetId"`

// parent associations as their child

}

