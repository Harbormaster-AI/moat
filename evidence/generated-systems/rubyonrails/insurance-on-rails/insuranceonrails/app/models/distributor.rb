class Distributor < ApplicationRecord
  enum DistributorType: [:Agency, :Broker, :Direct, :Bancassurance, :AffinityPartner, :OnlineAggregator]


  has_many :Insurers, class_name: 'Insurer'
  has_many :Agents, class_name: 'Agent'
  has_many :Policies, class_name: 'Policy'

end
