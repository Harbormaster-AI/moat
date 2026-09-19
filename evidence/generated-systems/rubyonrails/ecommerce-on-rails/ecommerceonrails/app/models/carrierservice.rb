class CarrierService < ApplicationRecord
  enum Carrier: [:UPS, :FedEx, :USPS, :DHL, :RoyalMail, :CanadaPost, :LocalCourier, :Other]
  enum ServiceLevel: [:Economy, :Standard, :Express, :Priority, :NextDay]


  has_many :ShippingMethods, class_name: 'ShippingMethod'

end
