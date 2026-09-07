package model

import (
    "gorm.io/gorm"
)

//==============================================================
// Measure Declaration
//==============================================================
type Measure struct {
    gorm.Model
     Name                                    string
    Format                                    string
    SemanticModelId         *uint
    SemanticModel           *SemanticModel `gorm:"foreignKey:SemanticModelId"`
     Datasets           []DataSet `gorm:"foreignKey:DatasetsFromMeasureId"`
     GlossaryTerms           []BusinessGlossaryTerm `gorm:"foreignKey:GlossaryTermsFromMeasureId"`
    Aggregation                      AggregationType

// parent associations as their child

}

