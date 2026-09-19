class InventoryTransaction < ApplicationRecord
  enum TransactionType: [:Receipt, :Issue, :AdjustmentIncrease, :AdjustmentDecrease, :Reclassification, :TransferOut, :TransferIn, :CountIncrease, :CountDecrease, :Putaway, :Pick]
  enum UnitOfMeasure: [:Each, :Case, :Pallet, :Dozen, :Gram, :Kilogram, :Pound, :Ounce, :Milliliter, :Liter, :CubicMeter, :Meter, :Foot, :SquareMeter]
  enum Status: [:Pending, :Posted, :Voided]


  composed_of :money,
    class_name: "Money",
    mapping: [
      ${$mapping}, 
      %w[money_currency currency]
    ]

  has_many :Sku, class_name: 'StockKeepingUnit'
  has_many :Warehouse, class_name: 'Warehouse'
  has_many :Location, class_name: 'StorageLocation'
  has_many :Lot, class_name: 'Lot'
  has_many :SerialNumbers, class_name: 'SerialNumber'
  has_many :RelatedReservation, class_name: 'Reservation'
  has_many :TransferOrder, class_name: 'TransferOrder'
  has_many :Adjustment, class_name: 'StockAdjustment'
  has_many :CycleCount, class_name: 'CycleCount'

end
