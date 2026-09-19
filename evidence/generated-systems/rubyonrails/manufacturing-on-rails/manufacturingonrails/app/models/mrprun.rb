class MRPRun < ApplicationRecord
  enum Status: [:Started, :Completed, :Failed, :Cancelled]


  has_many :Plant, class_name: 'Plant'
  has_many :PlannedOrders, class_name: 'PlannedOrder'

end
