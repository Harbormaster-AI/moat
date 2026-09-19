class StockKeepingUnit < ApplicationRecord
  enum ItemType: [:FinishedGood, :Component, :RawMaterial, :Packaging, :SparePart, :Consumable]
  enum UnitOfMeasure: [:Each, :Case, :Pallet, :Dozen, :Gram, :Kilogram, :Pound, :Ounce, :Milliliter, :Liter, :CubicMeter, :Meter, :Foot, :SquareMeter]


  composed_of :sKU,
    class_name: "SKU",
    mapping: [
      %w[sKU_value value]
    ]

  has_many :InventoryItems, class_name: 'InventoryItem'
  has_many :UomConversions, class_name: 'UoMConversion'
  has_many :ReplenishmentPolicies, class_name: 'ReplenishmentPolicy'
  has_many :Lots, class_name: 'Lot'
  has_many :SerialNumbers, class_name: 'SerialNumber'

end
