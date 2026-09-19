class Asset < ApplicationRecord
  enum AssetStatus: [:Commissioned, :Available, :InMaintenance, :Down, :Retired]


  has_many :Plant, class_name: 'Plant'
  has_many :WorkCenter, class_name: 'WorkCenter'
  has_many :MaintenanceOrders, class_name: 'MaintenanceOrder'
  has_many :MaintenancePlans, class_name: 'MaintenancePlan'

end
