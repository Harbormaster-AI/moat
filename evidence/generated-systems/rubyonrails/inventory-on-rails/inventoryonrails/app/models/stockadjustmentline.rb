class StockAdjustmentLine < ApplicationRecord
  enum UnitOfMeasure: [:Each, :Case, :Pallet, :Dozen, :Gram, :Kilogram, :Pound, :Ounce, :Milliliter, :Liter, :CubicMeter, :Meter, :Foot, :SquareMeter]
  enum StockStatus: [:Available, :Reserved, :Damaged, :Hold, :Quarantined, :InTransit, :PendingInspection]


  has_many :Adjustment, class_name: 'StockAdjustment'
  has_many :Sku, class_name: 'StockKeepingUnit'
  has_many :Lot, class_name: 'Lot'
  has_many :Location, class_name: 'StorageLocation'
  has_many :SerialNumbers, class_name: 'SerialNumber'

end
