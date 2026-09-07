package model

import (
    "gorm.io/gorm"
)

//==============================================================
// PayrollCalendar Declaration
//==============================================================
type PayrollCalendar struct {
    gorm.Model
     Name                                    string
    Country                                    string
    OrganizationId         *uint
    Organization           *Organization `gorm:"foreignKey:OrganizationId"`
     PayrollRuns           []PayrollRun `gorm:"foreignKey:PayrollRunsFromPayrollCalendarId"`
     Employees           []Employee `gorm:"foreignKey:EmployeesFromPayrollCalendarId"`
    PayFrequency                      PayFrequency

// parent associations as their child

}

