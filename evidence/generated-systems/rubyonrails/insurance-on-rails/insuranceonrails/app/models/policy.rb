class Policy < ApplicationRecord
  enum Status: [:Quoted, :Active, :Lapsed, :Cancelled, :Expired, :PendingCancel, :PendingReinstatement]
  enum PaymentPlan: [:Annual, :SemiAnnual, :Quarterly, :Monthly, :PayInFull]


  composed_of :policyNumber,
    class_name: "PolicyNumber",
    mapping: [
      %w[policyNumber_value value]
    ]

  composed_of :dateRange,
    class_name: "DateRange",
    mapping: [
      ${$mapping}, 
      %w[dateRange_endDate endDate]
    ]

  composed_of :money,
    class_name: "Money",
    mapping: [
      ${$mapping}, 
      %w[money_currency currency]
    ]

  has_many :Insurer, class_name: 'Insurer'
  has_many :Customer, class_name: 'Customer'
  has_many :Product, class_name: 'InsuranceProduct'
  has_many :Agent, class_name: 'Agent'
  has_many :Coverages, class_name: 'PolicyCoverage'
  has_many :InsuredObjects, class_name: 'InsuredObject'
  has_many :Endorsements, class_name: 'Endorsement'
  has_many :BillingAccount, class_name: 'BillingAccount'
  has_many :Beneficiaries, class_name: 'Beneficiary'
  has_many :Claims, class_name: 'Claim'
  has_many :ReinsuranceAgreements, class_name: 'ReinsuranceAgreement'

end
