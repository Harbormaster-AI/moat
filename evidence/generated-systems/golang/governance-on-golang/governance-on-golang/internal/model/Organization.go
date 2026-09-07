package model

import (
    "gorm.io/gorm"
)

//==============================================================
// Organization Declaration
//==============================================================
type Organization struct {
    gorm.Model
     Name                                    string
    LegalName                                    string
    Jurisdiction                                    string
    IndustrySector                                    string
     GovernanceBodies           []GovernanceBody `gorm:"foreignKey:GovernanceBodiesFromOrganizationId"`
     Policies           []Policy `gorm:"foreignKey:PoliciesFromOrganizationId"`
     Risks           []Risk `gorm:"foreignKey:RisksFromOrganizationId"`
     ThirdParties           []ThirdParty `gorm:"foreignKey:ThirdPartiesFromOrganizationId"`
     RecordsRepositories           []RecordsRepository `gorm:"foreignKey:RecordsRepositoriesFromOrganizationId"`
     DataProcessingActivities           []DataProcessingActivity `gorm:"foreignKey:DataProcessingActivitiesFromOrganizationId"`
     CompliancePrograms           []ComplianceProgram `gorm:"foreignKey:ComplianceProgramsFromOrganizationId"`
     AuditPrograms           []AuditProgram `gorm:"foreignKey:AuditProgramsFromOrganizationId"`
     BusinessUnits           []BusinessUnit `gorm:"foreignKey:BusinessUnitsFromOrganizationId"`
     Matters           []Matter `gorm:"foreignKey:MattersFromOrganizationId"`
     DataBreaches           []DataBreach `gorm:"foreignKey:DataBreachesFromOrganizationId"`

// parent associations as their child

}

