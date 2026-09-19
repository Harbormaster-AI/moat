class DirectDebitMandate < ApplicationRecord
  enum Scheme: [:SEPA, :ACH, :BACS, :BECS]
  enum Status: [:Active, :Suspended, :Cancelled, :Expired]


  composed_of :dateTime,
    class_name: "DateTime",
    mapping: [
      %w[dateTime_value value]
    ]

  has_many :Account, class_name: 'Account'
  has_many :Creditor, class_name: 'Creditor'

end
