package model

import (
    "time"
    "gorm.io/gorm"
)

//==============================================================
// ImagingReport Declaration
//==============================================================
type ImagingReport struct {
    gorm.Model
     ReportNumber                                    string
    Impression                                    string
    ReportedDate                                                            time.Time
    ImagingOrderId         *uint
    ImagingOrder           *ImagingOrder `gorm:"foreignKey:ImagingOrderId"`
    ClinicianId         *uint
    Clinician           *Clinician `gorm:"foreignKey:ClinicianId"`
    EncounterId         *uint
    Encounter           *Encounter `gorm:"foreignKey:EncounterId"`
    ImagingCenterId         *uint
    ImagingCenter           *ImagingCenter `gorm:"foreignKey:ImagingCenterId"`
    Status                      ResultStatus

// parent associations as their child

}

