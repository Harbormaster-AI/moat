class BuildSchedule < ApplicationRecord
  enum Status: [:Draft, :Published, :Revised, :Closed]


  has_many :ProductionOrders, class_name: 'ProductionOrder'

end
