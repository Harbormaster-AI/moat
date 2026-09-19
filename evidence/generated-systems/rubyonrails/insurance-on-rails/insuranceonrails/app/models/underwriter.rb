class Underwriter < ApplicationRecord


  composed_of :money,
    class_name: "Money",
    mapping: [
      ${$mapping}, 
      %w[money_currency currency]
    ]

  has_many :Decisions, class_name: 'UnderwritingDecision'
  has_many :Insurer, class_name: 'Insurer'

end
