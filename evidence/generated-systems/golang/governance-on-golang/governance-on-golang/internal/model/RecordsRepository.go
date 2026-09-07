package model

import (
    "gorm.io/gorm"
)

//==============================================================
// RecordsRepository Declaration
//==============================================================
type RecordsRepository struct {
    gorm.Model
     Name                                    string
    Location                                    string
    OwnerDepartment                                    string
    OrganizationId         *uint
    Organization           *Organization `gorm:"foreignKey:OrganizationId"`
     Records           []Record_ `gorm:"foreignKey:RecordsFromRecordsRepositoryId"`
     Systems           []System_ `gorm:"foreignKey:SystemsFromRecordsRepositoryId"`
     RetentionSchedules           []RetentionSchedule `gorm:"foreignKey:RetentionSchedulesFromRecordsRepositoryId"`
     LegalHolds           []LegalHold `gorm:"foreignKey:LegalHoldsFromRecordsRepositoryId"`
    RepositoryType                      RepositoryType

// parent associations as their child

}

