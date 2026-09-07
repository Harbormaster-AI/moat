package model

import (
    "time"
    "gorm.io/gorm"
)

//==============================================================
// Claim Declaration
//==============================================================
type Claim struct {
    gorm.Model
     ClaimNumber                                                            string
    NoticeDate                                                            time.Time
    LossDate                                                            time.Time
    ReportedBy                                    string
    PolicyId         *uint
    Policy           *Policy `gorm:"foreignKey:PolicyId"`
    CustomerId         *uint
    Customer           *Customer `gorm:"foreignKey:CustomerId"`
    AdjusterId         *uint
    Adjuster           *Adjuster `gorm:"foreignKey:AdjusterId"`
    IncidentId         *uint
    Incident           *Incident `gorm:"foreignKey:IncidentId"`
     Exposures           []Exposure `gorm:"foreignKey:ExposuresFromClaimId"`
     Reserves           []ClaimReserve `gorm:"foreignKey:ReservesFromClaimId"`
     ClaimPayments           []ClaimPayment `gorm:"foreignKey:ClaimPaymentsFromClaimId"`
     ServiceProviders           []ServiceProvider `gorm:"foreignKey:ServiceProvidersFromClaimId"`
     Subrogations           []SubrogationRecovery `gorm:"foreignKey:SubrogationsFromClaimId"`
    Status                      ClaimStatus
    LossCause                      CauseOfLoss

// parent associations as their child

}

