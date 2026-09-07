package model

import (
    "gorm.io/gorm"
)

//==============================================================
// SemanticModel Declaration
//==============================================================
type SemanticModel struct {
    gorm.Model
     Name                                    string
    Version                                    string
    Grain                                    string
     Datasets           []DataSet `gorm:"foreignKey:DatasetsFromSemanticModelId"`
     Metrics           []Metric `gorm:"foreignKey:MetricsFromSemanticModelId"`
     Dimensions           []Dimension `gorm:"foreignKey:DimensionsFromSemanticModelId"`
     Measures           []Measure `gorm:"foreignKey:MeasuresFromSemanticModelId"`
     GlossaryTerms           []BusinessGlossaryTerm `gorm:"foreignKey:GlossaryTermsFromSemanticModelId"`

// parent associations as their child

}

