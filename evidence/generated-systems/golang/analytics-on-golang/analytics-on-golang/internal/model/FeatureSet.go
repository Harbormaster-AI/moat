package model

import (
    "gorm.io/gorm"
)

//==============================================================
// FeatureSet Declaration
//==============================================================
type FeatureSet struct {
    gorm.Model
     Name                                    string
    RefreshSchedule                                                            string
    WorkspaceId         *uint
    Workspace           *AnalyticsWorkspace `gorm:"foreignKey:WorkspaceId"`
     Features           []Feature `gorm:"foreignKey:FeaturesFromFeatureSetId"`
     Datasets           []DataSet `gorm:"foreignKey:DatasetsFromFeatureSetId"`
     Models           []Model_ `gorm:"foreignKey:ModelsFromFeatureSetId"`
     ModelVersions           []ModelVersion `gorm:"foreignKey:ModelVersionsFromFeatureSetId"`
     Tags           []Tag `gorm:"foreignKey:TagsFromFeatureSetId"`
    StoreType                      FeatureStoreType

// parent associations as their child

}

