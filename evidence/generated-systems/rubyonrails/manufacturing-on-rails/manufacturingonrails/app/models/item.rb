class Item < ApplicationRecord
  enum ItemType: [:FinishedGood, :Subassembly, :Component, :RawMaterial, :Consumable, :Service]
  enum ProcurementType: [:MakeToStock, :MakeToOrder, :Purchase, :Kanban, :Outsourced]
  enum UnitOfMeasure: [:Each, :Kilogram, :Gram, :Pound, :Liter, :Meter, :Centimeter, :Millimeter, :Hour, :Minute, :Box, :Pallet]
  enum LifecycleStatus: [:Active, :PendingApproval, :Discontinued, :Obsolete]


  composed_of :money,
    class_name: "Money",
    mapping: [
      ${$mapping}, 
      %w[money_currency currency]
    ]

  composed_of :measurement,
    class_name: "Measurement",
    mapping: [
      ${$mapping}, 
      %w[measurement_unit unit]
    ]

  has_many :BusinessUnit, class_name: 'BusinessUnit'
  has_many :Boms, class_name: 'BOM'
  has_many :Routings, class_name: 'Routing'
  has_many :Suppliers, class_name: 'Supplier'
  has_many :QualitySpecifications, class_name: 'QualitySpecification'
  has_many :InventoryItems, class_name: 'InventoryItem'

end
