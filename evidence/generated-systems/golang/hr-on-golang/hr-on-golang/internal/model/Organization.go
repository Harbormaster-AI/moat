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
    RegistrationCountry                                    string
    Website                                    string
     Departments           []Department `gorm:"foreignKey:DepartmentsFromOrganizationId"`
     Locations           []Location `gorm:"foreignKey:LocationsFromOrganizationId"`
     JobFamilies           []JobFamily `gorm:"foreignKey:JobFamiliesFromOrganizationId"`
     BenefitPlans           []BenefitPlan `gorm:"foreignKey:BenefitPlansFromOrganizationId"`
     CostCenters           []CostCenter `gorm:"foreignKey:CostCentersFromOrganizationId"`
     PayrollCalendars           []PayrollCalendar `gorm:"foreignKey:PayrollCalendarsFromOrganizationId"`

// parent associations as their child

}

