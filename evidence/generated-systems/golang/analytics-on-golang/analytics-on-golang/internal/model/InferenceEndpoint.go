package model

import (
    "gorm.io/gorm"
)

//==============================================================
// InferenceEndpoint Declaration
//==============================================================
type InferenceEndpoint struct {
    gorm.Model
     Name                                    string
    EndpointUrl                                    string
    TrafficShare                                                            string
    ModelVersionId         *uint
    ModelVersion           *ModelVersion `gorm:"foreignKey:ModelVersionId"`
    WorkspaceId         *uint
    Workspace           *AnalyticsWorkspace `gorm:"foreignKey:WorkspaceId"`
     Predictions           []Prediction `gorm:"foreignKey:PredictionsFromInferenceEndpointId"`
    Mode                      InferenceMode

// parent associations as their child

}

