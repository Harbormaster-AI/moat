class ClaimReserve < ApplicationRecord
  enum ReserveType: [:Indemnity, :Expense, :Legal, :Medical]
  enum Status: [:Open, :Released, :Increased, :Decreased, :Closed]


  composed_of :money,
    class_name: "Money",
    mapping: [
      ${$mapping}, 
      %w[money_currency currency]
    ]

  has_many :Claim, class_name: 'Claim'
  has_many :Exposure, class_name: 'Exposure'

end
