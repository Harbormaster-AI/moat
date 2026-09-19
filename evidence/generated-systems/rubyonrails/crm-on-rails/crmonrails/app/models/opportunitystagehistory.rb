class OpportunityStageHistory < ApplicationRecord
  enum FromStage: [:Qualification, :Discovery, :Proposal, :Negotiation, :ClosedWon, :ClosedLost]
  enum ToStage: [:Qualification, :Discovery, :Proposal, :Negotiation, :ClosedWon, :ClosedLost]


  has_many :Opportunity, class_name: 'Opportunity'
  has_many :ChangedBy, class_name: 'User'

end
