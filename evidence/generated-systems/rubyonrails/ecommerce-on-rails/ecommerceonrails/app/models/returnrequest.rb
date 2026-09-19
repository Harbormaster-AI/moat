class ReturnRequest < ApplicationRecord
  enum Status: [:Requested, :Approved, :Rejected, :InTransit, :Received, :Refunded, :Closed]


  composed_of :money,
    class_name: "Money",
    mapping: [
      ${$mapping}, 
      %w[money_currency currency]
    ]

  has_many :Order, class_name: 'Order'
  has_many :Items, class_name: 'ReturnItem'
  has_many :Refund, class_name: 'Refund'
  has_many :Shipment, class_name: 'Shipment'

end
