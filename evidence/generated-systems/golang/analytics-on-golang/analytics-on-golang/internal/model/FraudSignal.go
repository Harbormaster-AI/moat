package model

import (
    "gorm.io/gorm"
)

//==============================================================
// FraudSignal Declaration
//==============================================================
type FraudSignal struct {
    gorm.Model
     Name                                    string
    RuleLogic                                    string
    ScenarioId         *uint
    Scenario           *FraudScenario `gorm:"foreignKey:ScenarioId"`
    DatasetId         *uint
    Dataset           *DataSet `gorm:"foreignKey:DatasetId"`
    ModelVersionId         *uint
    ModelVersion           *ModelVersion `gorm:"foreignKey:ModelVersionId"`
    SignalType                      FraudSignalType

// parent associations as their child

}

