package model

import (
    "gorm.io/gorm"
)

//==============================================================
// RecommendationScenario Declaration
//==============================================================
type RecommendationScenario struct {
    gorm.Model
     Name                                    string
    Objective                                    string
     Models           []Model_ `gorm:"foreignKey:ModelsFromRecommendationScenarioId"`
     Datasets           []DataSet `gorm:"foreignKey:DatasetsFromRecommendationScenarioId"`
     Experiments           []Experiment `gorm:"foreignKey:ExperimentsFromRecommendationScenarioId"`
     Alerts           []Alert `gorm:"foreignKey:AlertsFromRecommendationScenarioId"`
    RecommendationType                      RecommendationType

// parent associations as their child

}

