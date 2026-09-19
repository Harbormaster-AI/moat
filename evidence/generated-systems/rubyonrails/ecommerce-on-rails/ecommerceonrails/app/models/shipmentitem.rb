class ShipmentItem < ApplicationRecord


  has_many :Shipment, class_name: 'Shipment'
  has_many :OrderLine, class_name: 'OrderLine'

end
