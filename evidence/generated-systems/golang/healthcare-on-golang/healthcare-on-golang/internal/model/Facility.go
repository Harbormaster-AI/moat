package model

import (
    "gorm.io/gorm"
)

//==============================================================
// Facility Declaration
//==============================================================
type Facility struct {
    gorm.Model
     Name                                    string
    FacilityCode                                    string
    Address                                                            string
    HealthSystemId         *uint
    HealthSystem           *HealthSystem `gorm:"foreignKey:HealthSystemId"`
     Departments           []Department `gorm:"foreignKey:DepartmentsFromFacilityId"`
     CareTeams           []CareTeam `gorm:"foreignKey:CareTeamsFromFacilityId"`
     Laboratories           []Laboratory `gorm:"foreignKey:LaboratoriesFromFacilityId"`
     ImagingCenters           []ImagingCenter `gorm:"foreignKey:ImagingCentersFromFacilityId"`
     Pharmacies           []Pharmacy `gorm:"foreignKey:PharmaciesFromFacilityId"`
     InventoryItems           []InventoryItem `gorm:"foreignKey:InventoryItemsFromFacilityId"`
    FacilityType                      FacilityType

// parent associations as their child

}

