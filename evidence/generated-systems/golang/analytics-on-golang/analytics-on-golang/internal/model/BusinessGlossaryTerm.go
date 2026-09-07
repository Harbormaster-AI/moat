package model

import (
    "gorm.io/gorm"
)

//==============================================================
// BusinessGlossaryTerm Declaration
//==============================================================
type BusinessGlossaryTerm struct {
    gorm.Model
     Term                                    string
    Definition                                    string
    Steward                                    string
     RelatedTerms           []BusinessGlossaryTerm `gorm:"foreignKey:RelatedTermsFromBusinessGlossaryTermId"`
     Metrics           []Metric `gorm:"foreignKey:MetricsFromBusinessGlossaryTermId"`
     Datasets           []DataSet `gorm:"foreignKey:DatasetsFromBusinessGlossaryTermId"`
     Dimensions           []Dimension `gorm:"foreignKey:DimensionsFromBusinessGlossaryTermId"`
     Measures           []Measure `gorm:"foreignKey:MeasuresFromBusinessGlossaryTermId"`

// parent associations as their child

}

