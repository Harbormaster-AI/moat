package model

import (
    "gorm.io/gorm"
)

//==============================================================
// ThirdParty Declaration
//==============================================================
type ThirdParty struct {
    gorm.Model
     Name                                    string
    Country                                    string
    ContactEmail                                                            string
    OrganizationId         *uint
    Organization           *Organization `gorm:"foreignKey:OrganizationId"`
     ProcessingActivities           []DataProcessingActivity `gorm:"foreignKey:ProcessingActivitiesFromThirdPartyId"`
     Assessments           []ThirdPartyAssessment `gorm:"foreignKey:AssessmentsFromThirdPartyId"`
     Contracts           []Contract `gorm:"foreignKey:ContractsFromThirdPartyId"`
     Obligations           []Obligation `gorm:"foreignKey:ObligationsFromThirdPartyId"`
     DataBreaches           []DataBreach `gorm:"foreignKey:DataBreachesFromThirdPartyId"`
    ThirdPartyType                      ThirdPartyType
    Criticality                      VendorCriticality

// parent associations as their child

}

