package model

import (
    "gorm.io/gorm"
)

//==============================================================
// Feature Declaration
//==============================================================
type Feature struct {
    gorm.Model
     Name                                    string
    Description                                    string
    FeatureSetId         *uint
    FeatureSet           *FeatureSet `gorm:"foreignKey:FeatureSetId"`
     SourceDatasets           []DataSet `gorm:"foreignKey:SourceDatasetsFromFeatureId"`
     Models           []Model_ `gorm:"foreignKey:ModelsFromFeatureId"`
     TrainingRuns           []TrainingRun `gorm:"foreignKey:TrainingRunsFromFeatureId"`
    DataType                      DataType

// parent associations as their child

}

