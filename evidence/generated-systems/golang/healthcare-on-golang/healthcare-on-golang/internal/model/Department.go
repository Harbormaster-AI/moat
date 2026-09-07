package model

import (
    "gorm.io/gorm"
)

//==============================================================
// Department Declaration
//==============================================================
type Department struct {
    gorm.Model
     Name                                    string
    FacilityId         *uint
    Facility           *Facility `gorm:"foreignKey:FacilityId"`
     CareTeams           []CareTeam `gorm:"foreignKey:CareTeamsFromDepartmentId"`
    DepartmentType                      DepartmentType

// parent associations as their child

}

