class Claim < ApplicationRecord
  enum Status: [:Open, :Closed, :Reopened, :Denied, :PendingInvestigation, :Litigation]
  enum LossCause: [:Collision, :Weather, :MechanicalFailure, :HumanError, :NaturalDisaster, :Theft, :Vandalism, :LiabilityClaim, :Illness]


  composed_of :claimNumber,
    class_name: "ClaimNumber",
    mapping: [
      %w[claimNumber_value value]
    ]

  has_many :Policy, class_name: 'Policy'
  has_many :Customer, class_name: 'Customer'
  has_many :Adjuster, class_name: 'Adjuster'
  has_many :Incident, class_name: 'Incident'
  has_many :Exposures, class_name: 'Exposure'
  has_many :Reserves, class_name: 'ClaimReserve'
  has_many :ClaimPayments, class_name: 'ClaimPayment'
  has_many :ServiceProviders, class_name: 'ServiceProvider'
  has_many :Subrogations, class_name: 'SubrogationRecovery'

end
