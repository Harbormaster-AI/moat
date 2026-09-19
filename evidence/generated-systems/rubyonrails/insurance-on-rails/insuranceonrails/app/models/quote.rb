class Quote < ApplicationRecord


  composed_of :money,
    class_name: "Money",
    mapping: [
      ${$mapping}, 
      %w[money_currency currency]
    ]

  has_many :Application, class_name: 'Application'
  has_many :UnderwritingDecisions, class_name: 'UnderwritingDecision'
  has_many :Policy, class_name: 'Policy'

end
