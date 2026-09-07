package model

import (
    "gorm.io/gorm"
)

//==============================================================
// Metric Declaration
//==============================================================
type Metric struct {
    gorm.Model
     Name                                    string
    Expression                                    string
    Unit                                    string
    SemanticModelId         *uint
    SemanticModel           *SemanticModel `gorm:"foreignKey:SemanticModelId"`
     Datasets           []DataSet `gorm:"foreignKey:DatasetsFromMetricId"`
     GlossaryTerms           []BusinessGlossaryTerm `gorm:"foreignKey:GlossaryTermsFromMetricId"`
     Alerts           []Alert `gorm:"foreignKey:AlertsFromMetricId"`
     Visualizations           []Visualization `gorm:"foreignKey:VisualizationsFromMetricId"`
    MetricType                      MetricType

// parent associations as their child

}

