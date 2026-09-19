class Shipment < ApplicationRecord
  enum Status: [:Pending, :Packed, :Shipped, :InTransit, :Delivered, :Delayed, :Returned, :Cancelled]
  enum Carrier: [:UPS, :FedEx, :USPS, :DHL, :RoyalMail, :CanadaPost, :LocalCourier, :Other]


  composed_of :address,
    class_name: "Address",
    mapping: [
      ${$mapping}, 
      ${$mapping}, 
      ${$mapping}, 
      ${$mapping}, 
      %w[address_country country]
    ]

  has_many :Order, class_name: 'Order'
  has_many :ShipmentItems, class_name: 'ShipmentItem'
  has_many :FulfillmentCenter, class_name: 'FulfillmentCenter'

end
