package model

import (
    "gorm.io/gorm"
)

//==============================================================
// JobFamily Declaration
//==============================================================
type JobFamily struct {
    gorm.Model
     Name                                    string
    Description                                    string
    OrganizationId         *uint
    Organization           *Organization `gorm:"foreignKey:OrganizationId"`
     JobProfiles           []JobProfile `gorm:"foreignKey:JobProfilesFromJobFamilyId"`

// parent associations as their child

}

