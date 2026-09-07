package model

import (
    "gorm.io/gorm"
)

//==============================================================
// FraudScenario Declaration
//==============================================================
type FraudScenario struct {
    gorm.Model
     Name                                    string
    RiskAppetite                                    string
     Models           []Model_ `gorm:"foreignKey:ModelsFromFraudScenarioId"`
     Datasets           []DataSet `gorm:"foreignKey:DatasetsFromFraudScenarioId"`
     Alerts           []Alert `gorm:"foreignKey:AlertsFromFraudScenarioId"`
     Signals           []FraudSignal `gorm:"foreignKey:SignalsFromFraudScenarioId"`
    DetectionType                      FraudDetectionType

// parent associations as their child

}

