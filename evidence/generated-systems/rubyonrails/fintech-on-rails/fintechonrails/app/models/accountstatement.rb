class AccountStatement < ApplicationRecord


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

  composed_of :dateTime,
    class_name: "DateTime",
    mapping: [
      %w[dateTime_value value]
    ]

  has_many :Account, class_name: 'Account'

end
