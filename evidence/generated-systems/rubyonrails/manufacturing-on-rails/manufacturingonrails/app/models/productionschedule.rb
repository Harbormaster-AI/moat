class ProductionSchedule < ApplicationRecord
  enum Status: [:Draft, :Approved, :Frozen, :Completed]


  has_many :Plant, class_name: 'Plant'
  has_many :WorkOrders, class_name: 'WorkOrder'

end
