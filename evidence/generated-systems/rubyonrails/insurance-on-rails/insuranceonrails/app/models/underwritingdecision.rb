class UnderwritingDecision < ApplicationRecord
  enum Decision: [:Approve, :ConditionalApprove, :Refer, :Decline]


  has_many :Quote, class_name: 'Quote'
  has_many :Underwriter, class_name: 'Underwriter'

end
