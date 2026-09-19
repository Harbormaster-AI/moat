class Customer < ApplicationRecord
  enum CustomerType: [:Individual, :Business]


  composed_of :email,
    class_name: "Email",
    mapping: [
      %w[email_value value]
    ]

  composed_of :phoneNumber,
    class_name: "PhoneNumber",
    mapping: [
      %w[phoneNumber_value value]
    ]

  composed_of :address,
    class_name: "Address",
    mapping: [
      ${$mapping}, 
      ${$mapping}, 
      ${$mapping}, 
      ${$mapping}, 
      %w[address_country country]
    ]

  composed_of :taxId,
    class_name: "TaxId",
    mapping: [
      %w[taxId_value value]
    ]

  composed_of :riskScore,
    class_name: "RiskScore",
    mapping: [
      %w[riskScore_value value]
    ]

  has_many :Institution, class_name: 'FinancialInstitution'
  has_many :Accounts, class_name: 'Account'
  has_many :Wallets, class_name: 'Wallet'
  has_many :Cards, class_name: 'PaymentCard'
  has_many :KycProfiles, class_name: 'KYCProfile'
  has_many :Consents, class_name: 'Consent'
  has_many :Agreements, class_name: 'Agreement'
  has_many :LoanApplications, class_name: 'LoanApplication'
  has_many :Loans, class_name: 'Loan'
  has_many :Portfolios, class_name: 'InvestmentPortfolio'
  has_many :Disputes, class_name: 'Dispute'

end
