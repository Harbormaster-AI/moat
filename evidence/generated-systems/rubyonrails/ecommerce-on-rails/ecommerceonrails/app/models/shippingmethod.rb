class ShippingMethod < ApplicationRecord
  enum MethodType: [:Standard, :Expedited, :Overnight, :SameDay, :Pickup]


  composed_of :money,
    class_name: "Money",
    mapping: [
      ${$mapping}, 
      %w[money_currency currency]
    ]

  has_many :Channels, class_name: 'Channel'
  has_many :CarrierService, class_name: 'CarrierService'

end
