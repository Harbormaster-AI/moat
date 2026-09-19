class PlannedOrder < ApplicationRecord
  enum OrderType: [:WorkOrder, :PurchaseRequisition, :TransferOrder]
  enum Status: [:Planned, :Firmed, :Released, :Cancelled]


  composed_of :quantity,
    class_name: "Quantity",
    mapping: [
      ${$mapping}, 
      %w[quantity_unit unit]
    ]

  has_many :MrpRun, class_name: 'MRPRun'
  has_many :Item, class_name: 'Item'
  has_many :Plant, class_name: 'Plant'

end
