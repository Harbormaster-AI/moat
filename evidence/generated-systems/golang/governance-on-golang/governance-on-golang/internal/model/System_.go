package model

import (
    "gorm.io/gorm"
)

//==============================================================
// System_ Declaration
//==============================================================
type System_ struct {
    gorm.Model
     Name                                    string
    OwnerDepartment                                    string
     ProcessingActivities           []DataProcessingActivity `gorm:"foreignKey:ProcessingActivitiesFromSystem_Id"`
     RecordsRepositories           []RecordsRepository `gorm:"foreignKey:RecordsRepositoriesFromSystem_Id"`
    SystemType                      SystemType

// parent associations as their child

}

