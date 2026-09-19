class Supplier < ApplicationRecord
  enum SupplierTier: [:Tier1, :Tier2, :Tier3]
  enum PaymentTerms: [:Net30, :Net45, :Net60, :Prepaid, :COD]


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
  has_many :Items, class_name: 'Item'
  has_many :PurchaseOrders, class_name: 'PurchaseOrder'

end
