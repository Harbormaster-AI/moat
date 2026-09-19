class Invoice < ApplicationRecord
  enum Status: [:Draft, :Issued, :PartiallyPaid, :Paid, :Overdue, :Cancelled]


  composed_of :money,
    class_name: "Money",
    mapping: [
      ${$mapping}, 
      %w[money_currency currency]
    ]

  has_many :Patient, class_name: 'Patient'
  has_many :Claim, class_name: 'Claim'
  has_many :Payments, class_name: 'Payment'

end
