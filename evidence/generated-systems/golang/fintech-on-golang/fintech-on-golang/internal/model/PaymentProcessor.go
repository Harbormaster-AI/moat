package model

import (
    "gorm.io/gorm"
)

//==============================================================
// PaymentProcessor Declaration
//==============================================================
type PaymentProcessor struct {
    gorm.Model
     Name                                    string
    ProcessorCode                                    string
    NetworkSupport                                    string
     Institutions           []FinancialInstitution `gorm:"foreignKey:InstitutionsFromPaymentProcessorId"`
     Contracts           []PaymentContract `gorm:"foreignKey:ContractsFromPaymentProcessorId"`
     Settlements           []SettlementBatch `gorm:"foreignKey:SettlementsFromPaymentProcessorId"`

// parent associations as their child

}

