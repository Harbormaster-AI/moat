class CycleCountEntry < ApplicationRecord
  enum StockStatus: [:Available, :Reserved, :Damaged, :Hold, :Quarantined, :InTransit, :PendingInspection]


  has_many :CycleCount, class_name: 'CycleCount'
  has_many :Sku, class_name: 'StockKeepingUnit'
  has_many :Lot, class_name: 'Lot'
  has_many :Location, class_name: 'StorageLocation'
  has_many :SerialNumbers, class_name: 'SerialNumber'

end
