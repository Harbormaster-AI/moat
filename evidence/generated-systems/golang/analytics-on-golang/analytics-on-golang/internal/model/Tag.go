package model

import (
    "gorm.io/gorm"
)

//==============================================================
// Tag Declaration
//==============================================================
type Tag struct {
    gorm.Model
     Name                                    string
     Datasets           []DataSet `gorm:"foreignKey:DatasetsFromTagId"`
     Models           []Model_ `gorm:"foreignKey:ModelsFromTagId"`
     ModelVersions           []ModelVersion `gorm:"foreignKey:ModelVersionsFromTagId"`
     Dashboards           []Dashboard `gorm:"foreignKey:DashboardsFromTagId"`
     Reports           []Report `gorm:"foreignKey:ReportsFromTagId"`
     FeatureSets           []FeatureSet `gorm:"foreignKey:FeatureSetsFromTagId"`
     Metrics           []Metric `gorm:"foreignKey:MetricsFromTagId"`
    Category                      TagCategory

// parent associations as their child

}

