class Cart < ApplicationRecord
  enum Status: [:Active, :Merged, :Ordered, :Abandoned]


  composed_of :address,
    class_name: "Address",
    mapping: [
      ${$mapping}, 
      ${$mapping}, 
      ${$mapping}, 
      ${$mapping}, 
      %w[address_country country]
    ]

  composed_of :address,
    class_name: "Address",
    mapping: [
      ${$mapping}, 
      ${$mapping}, 
      ${$mapping}, 
      ${$mapping}, 
      %w[address_country country]
    ]

  has_many :Customer, class_name: 'Customer'
  has_many :Channel, class_name: 'Channel'
  has_many :Items, class_name: 'CartItem'
  has_many :AppliedPromotions, class_name: 'Promotion'

end
