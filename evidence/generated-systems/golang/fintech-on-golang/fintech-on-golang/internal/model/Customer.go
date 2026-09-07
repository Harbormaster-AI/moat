package model

import (
    "time"
    "gorm.io/gorm"
)

//==============================================================
// Customer Declaration
//==============================================================
type Customer struct {
    gorm.Model
     FirstName                                    string
    LastName                                    string
    DateOfBirth                                                            time.Time
    Email                                                            string
    Phone                                                            string
    Address                                                            string
    TaxId                                                            string
    RiskScore                                                            string
    InstitutionId         *uint
    Institution           *FinancialInstitution `gorm:"foreignKey:InstitutionId"`
     Accounts           []Account `gorm:"foreignKey:AccountsFromCustomerId"`
     Wallets           []Wallet `gorm:"foreignKey:WalletsFromCustomerId"`
     Cards           []PaymentCard `gorm:"foreignKey:CardsFromCustomerId"`
     KycProfiles           []KYCProfile `gorm:"foreignKey:KycProfilesFromCustomerId"`
     Consents           []Consent `gorm:"foreignKey:ConsentsFromCustomerId"`
     Agreements           []Agreement `gorm:"foreignKey:AgreementsFromCustomerId"`
     LoanApplications           []LoanApplication `gorm:"foreignKey:LoanApplicationsFromCustomerId"`
     Loans           []Loan `gorm:"foreignKey:LoansFromCustomerId"`
     Portfolios           []InvestmentPortfolio `gorm:"foreignKey:PortfoliosFromCustomerId"`
     Disputes           []Dispute `gorm:"foreignKey:DisputesFromCustomerId"`
    CustomerType                      CustomerType

// parent associations as their child

}

