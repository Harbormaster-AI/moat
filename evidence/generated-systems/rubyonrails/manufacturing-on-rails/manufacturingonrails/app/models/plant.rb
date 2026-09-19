class Plant < ApplicationRecord


  composed_of :address,
    class_name: "Address",
    mapping: [
      ${$mapping}, 
      ${$mapping}, 
      ${$mapping}, 
      ${$mapping}, 
      %w[address_country country]
    ]

  has_many :Enterprise, class_name: 'Enterprise'
  has_many :ProductionLines, class_name: 'ProductionLine'
  has_many :WorkCenters, class_name: 'WorkCenter'
  has_many :Warehouses, class_name: 'Warehouse'
  has_many :Assets, class_name: 'Asset'
  has_many :ProductionSchedules, class_name: 'ProductionSchedule'

end
