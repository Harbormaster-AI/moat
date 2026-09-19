class Terminal < ApplicationRecord
  enum Type: [:POS, :mPOS, :ECommerce]
  enum Status: [:Active, :Inactive, :Decommissioned]


  composed_of :address,
    class_name: "Address",
    mapping: [
      ${$mapping}, 
      ${$mapping}, 
      ${$mapping}, 
      ${$mapping}, 
      %w[address_country country]
    ]

  has_many :Merchant, class_name: 'Merchant'

end
