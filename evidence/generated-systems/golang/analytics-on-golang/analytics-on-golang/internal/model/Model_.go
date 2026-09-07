package model

import (
    "gorm.io/gorm"
)

//==============================================================
// Model_ Declaration
//==============================================================
type Model_ struct {
    gorm.Model
     Name                                    string
    TaskDescription                                    string
    WorkspaceId         *uint
    Workspace           *AnalyticsWorkspace `gorm:"foreignKey:WorkspaceId"`
     Versions           []ModelVersion `gorm:"foreignKey:VersionsFromModel_Id"`
     FeatureSets           []FeatureSet `gorm:"foreignKey:FeatureSetsFromModel_Id"`
     Experiments           []Experiment `gorm:"foreignKey:ExperimentsFromModel_Id"`
     Tags           []Tag `gorm:"foreignKey:TagsFromModel_Id"`
    ModelType                      ModelType

// parent associations as their child

}

