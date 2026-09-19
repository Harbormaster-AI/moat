class Customer < ApplicationRecord
  enum CustomerType: [:Distributor, :OEM, :Retailer, :Direct]


  composed_of :address,
    class_name: "Address",
    mapping: [
      ${$mapping}, 
      ${$mapping}, 
      ${$mapping}, 
      ${$mapping}, 
      %w[address_country country]
    ]

  has_many :Enterprises, class_name: 'Enterprise'
  has_many :SalesOrders, class_name: 'SalesOrder'

end
