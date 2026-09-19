class MaintenancePlan < ApplicationRecord
  enum Strategy: [:TimeBased, :UsageBased, :ConditionBased, :Predictive, :Corrective]


  composed_of :timeDuration,
    class_name: "TimeDuration",
    mapping: [
      ${$mapping}, 
      %w[timeDuration_unit unit]
    ]

  has_many :Asset, class_name: 'Asset'
  has_many :MaintenanceOrders, class_name: 'MaintenanceOrder'

end
