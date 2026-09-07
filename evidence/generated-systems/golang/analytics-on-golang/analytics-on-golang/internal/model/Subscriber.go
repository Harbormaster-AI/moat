package model

import (
    "gorm.io/gorm"
)

//==============================================================
// Subscriber Declaration
//==============================================================
type Subscriber struct {
    gorm.Model
     Name                                    string
    Address                                    string
     Alerts           []Alert `gorm:"foreignKey:AlertsFromSubscriberId"`
    Channel                      NotificationChannel

// parent associations as their child

}

