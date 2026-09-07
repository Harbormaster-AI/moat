package model

import (
    "gorm.io/gorm"
)

//==============================================================
// CareTeam Declaration
//==============================================================
type CareTeam struct {
    gorm.Model
     Name                                    string
    DepartmentId         *uint
    Department           *Department `gorm:"foreignKey:DepartmentId"`
     Clinicians           []Clinician `gorm:"foreignKey:CliniciansFromCareTeamId"`
     Patients           []Patient `gorm:"foreignKey:PatientsFromCareTeamId"`
    CareSetting                      CareSettingType

// parent associations as their child

}

