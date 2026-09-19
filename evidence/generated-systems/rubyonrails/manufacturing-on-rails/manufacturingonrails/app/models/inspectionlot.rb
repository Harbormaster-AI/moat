class InspectionLot < ApplicationRecord
  enum InspectionType: [:Incoming, :InProcess, :Final, :Audit]
  enum Status: [:Open, :InProgress, :Completed, :Accepted, :Rejected]


  composed_of :quantity,
    class_name: "Quantity",
    mapping: [
      ${$mapping}, 
      %w[quantity_unit unit]
    ]

  has_many :Item, class_name: 'Item'
  has_many :WorkOrder, class_name: 'WorkOrder'
  has_many :GoodsReceipt, class_name: 'GoodsReceipt'
  has_many :Results, class_name: 'InspectionResult'

end
