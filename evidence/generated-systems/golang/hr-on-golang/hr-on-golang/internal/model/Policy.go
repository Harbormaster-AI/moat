package model

import (
    "time"
    "gorm.io/gorm"
)

//==============================================================
// Policy Declaration
//==============================================================
type Policy struct {
    gorm.Model
     PolicyNumber                                    string
    Name                                    string
    EffectiveDate                                                            time.Time
    Description                                    string
    OrganizationId         *uint
    Organization           *Organization `gorm:"foreignKey:OrganizationId"`
     Acknowledgements           []PolicyAcknowledgement `gorm:"foreignKey:AcknowledgementsFromPolicyId"`

// parent associations as their child

}

