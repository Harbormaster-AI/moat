package model

import (
    "time"
    "gorm.io/gorm"
)

//==============================================================
// Application Declaration
//==============================================================
type Application struct {
    gorm.Model
     ApplicationNumber                                    string
    SubmissionDate                                                            time.Time
    CustomerId         *uint
    Customer           *Customer `gorm:"foreignKey:CustomerId"`
    ProductId         *uint
    Product           *InsuranceProduct `gorm:"foreignKey:ProductId"`
    DistributorId         *uint
    Distributor           *Distributor `gorm:"foreignKey:DistributorId"`
     Quotes           []Quote `gorm:"foreignKey:QuotesFromApplicationId"`
    SelectedQuoteId         *uint
    SelectedQuote           *Quote `gorm:"foreignKey:SelectedQuoteId"`
    Status                      ApplicationStatus

// parent associations as their child

}

