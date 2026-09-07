package model

import (
    "gorm.io/gorm"
)

//==============================================================
// ClinicalOrder Declaration
//==============================================================
type ClinicalOrder struct {
    gorm.Model
     OrderNumber                                    string
    PatientId         *uint
    Patient           *Patient `gorm:"foreignKey:PatientId"`
    EncounterId         *uint
    Encounter           *Encounter `gorm:"foreignKey:EncounterId"`
    OrderingClinicianId         *uint
    OrderingClinician           *Clinician `gorm:"foreignKey:OrderingClinicianId"`
     MedicationOrders           []MedicationOrder `gorm:"foreignKey:MedicationOrdersFromClinicalOrderId"`
     LaboratoryOrders           []LaboratoryOrder `gorm:"foreignKey:LaboratoryOrdersFromClinicalOrderId"`
     ImagingOrders           []ImagingOrder `gorm:"foreignKey:ImagingOrdersFromClinicalOrderId"`
     ProcedureOrders           []ProcedureOrder `gorm:"foreignKey:ProcedureOrdersFromClinicalOrderId"`
     Authorizations           []Authorization `gorm:"foreignKey:AuthorizationsFromClinicalOrderId"`
    Status                      OrderStatus
    OrderType                      ClinicalOrderType
    Priority                      Priority

// parent associations as their child

}

