package model

import (
    "gorm.io/gorm"
)

//==============================================================
// Model Declaration
//==============================================================
type Model struct {
    gorm.Model
     Name                                    string
    TaskDescription                                    string
    WorkspaceId         *uint
    Workspace           *AnalyticsWorkspace `gorm:"foreignKey:WorkspaceId"`
     Versions           []ModelVersion `gorm:"foreignKey:VersionsFromModelId"`
     FeatureSets           []FeatureSet `gorm:"foreignKey:FeatureSetsFromModelId"`
     Experiments           []Experiment `gorm:"foreignKey:ExperimentsFromModelId"`
     Tags           []Tag `gorm:"foreignKey:TagsFromModelId"`
    ModelType                      ModelType

// parent associations as their child

}

