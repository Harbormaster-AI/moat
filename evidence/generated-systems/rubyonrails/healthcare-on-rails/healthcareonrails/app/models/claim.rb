class Claim < ApplicationRecord
  enum Status: [:Submitted, :InProcess, :Paid, :Denied, :Adjusted, :Void]


  composed_of :money,
    class_name: "Money",
    mapping: [
      ${$mapping}, 
      %w[money_currency currency]
    ]

  has_many :Patient, class_name: 'Patient'
  has_many :Coverage, class_name: 'Coverage'
  has_many :Encounter, class_name: 'Encounter'
  has_many :Invoices, class_name: 'Invoice'
  has_many :Payer, class_name: 'InsurancePayer'

end
