class WorkCenter < ApplicationRecord
  enum WorkCenterType: [:Machining, :Assembly, :Painting, :Packaging, :Test, :Warehouse]


  composed_of :percentage,
    class_name: "Percentage",
    mapping: [
      %w[percentage_value value]
    ]

  has_many :ProductionLine, class_name: 'ProductionLine'
  has_many :Assets, class_name: 'Asset'
  has_many :MaintenanceOrders, class_name: 'MaintenanceOrder'

end
