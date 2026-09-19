class WorkOrder < ApplicationRecord
  enum Status: [:Planned, :Released, :InProcess, :Hold, :Completed, :Closed, :Cancelled]


  composed_of :quantity,
    class_name: "Quantity",
    mapping: [
      ${$mapping}, 
      %w[quantity_unit unit]
    ]

  has_many :Item, class_name: 'Item'
  has_many :Plant, class_name: 'Plant'
  has_many :Routing, class_name: 'Routing'
  has_many :Bom, class_name: 'BOM'
  has_many :ProductionSchedule, class_name: 'ProductionSchedule'
  has_many :SalesOrder, class_name: 'SalesOrder'

end
