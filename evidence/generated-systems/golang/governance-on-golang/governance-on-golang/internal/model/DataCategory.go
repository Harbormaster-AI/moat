package model

import (
    "gorm.io/gorm"
)

//==============================================================
// DataCategory Declaration
//==============================================================
type DataCategory struct {
    gorm.Model
     Name                                    string
    Description                                    string
     ProcessingActivities           []DataProcessingActivity `gorm:"foreignKey:ProcessingActivitiesFromDataCategoryId"`
     Records           []Record_ `gorm:"foreignKey:RecordsFromDataCategoryId"`
     DataBreaches           []DataBreach `gorm:"foreignKey:DataBreachesFromDataCategoryId"`
    Classification                      DataClassificationLevel

// parent associations as their child

}

