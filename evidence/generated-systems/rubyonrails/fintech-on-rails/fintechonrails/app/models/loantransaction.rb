class LoanTransaction < ApplicationRecord
  enum Type: [:Disbursement, :Repayment, :Interest, :Fee, :Reversal]
  enum Status: [:Pending, :Posted, :Reversed]


  composed_of :transactionId,
    class_name: "TransactionId",
    mapping: [
      %w[transactionId_value value]
    ]

  composed_of :money,
    class_name: "Money",
    mapping: [
      ${$mapping}, 
      %w[money_currency currency]
    ]

  has_many :Loan, class_name: 'Loan'

end
