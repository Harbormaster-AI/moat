package model

import (
    "gorm.io/gorm"
)

//==============================================================
// CarePlan Declaration
//==============================================================
type CarePlan struct {
    gorm.Model
     PlanNumber                                    string
    GoalSummary                                    string
    PatientId         *uint
    Patient           *Patient `gorm:"foreignKey:PatientId"`
     Encounters           []Encounter `gorm:"foreignKey:EncountersFromCarePlanId"`
     Tasks           []CareTask `gorm:"foreignKey:TasksFromCarePlanId"`
    CareTeamId         *uint
    CareTeam           *CareTeam `gorm:"foreignKey:CareTeamId"`
    Status                      CarePlanStatus

// parent associations as their child

}

