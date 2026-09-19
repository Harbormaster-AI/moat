class ReinsuranceAgreement < ApplicationRecord
  enum ReinsuranceType: [:Treaty, :Facultative]
  enum TreatyType: [:QuotaShare, :Surplus, :ExcessOfLoss, :StopLoss]


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

  composed_of :money,
    class_name: "Money",
    mapping: [
      ${$mapping}, 
      %w[money_currency currency]
    ]

  composed_of :percentage,
    class_name: "Percentage",
    mapping: [
      %w[percentage_value value]
    ]

  has_many :Insurer, class_name: 'Insurer'
  has_many :Policies, class_name: 'Policy'

end
