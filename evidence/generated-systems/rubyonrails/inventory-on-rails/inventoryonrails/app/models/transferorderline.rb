class TransferOrderLine < ApplicationRecord
  enum UnitOfMeasure: [:Each, :Case, :Pallet, :Dozen, :Gram, :Kilogram, :Pound, :Ounce, :Milliliter, :Liter, :CubicMeter, :Meter, :Foot, :SquareMeter]
  enum StockStatus: [:Available, :Reserved, :Damaged, :Hold, :Quarantined, :InTransit, :PendingInspection]


  has_many :TransferOrder, class_name: 'TransferOrder'
  has_many :Sku, class_name: 'StockKeepingUnit'
  has_many :Lot, class_name: 'Lot'
  has_many :SerialNumbers, class_name: 'SerialNumber'
  has_many :FromLocation, class_name: 'StorageLocation'
  has_many :ToLocation, class_name: 'StorageLocation'

end
