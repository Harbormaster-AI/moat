class Contract < ApplicationRecord
  enum Status: [:Draft, :Active, :Expiring, :Terminated, :Archived]


  composed_of :uRL,
    class_name: "URL",
    mapping: [
      %w[uRL_value value]
    ]

  has_many :ThirdParty, class_name: 'ThirdParty'
  has_many :Obligations, class_name: 'Obligation'
  has_many :DataProcessingActivities, class_name: 'DataProcessingActivity'
  has_many :Matter, class_name: 'Matter'

end
