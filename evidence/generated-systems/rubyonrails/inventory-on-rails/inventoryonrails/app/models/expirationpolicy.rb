class ExpirationPolicy < ApplicationRecord
  enum RotationMethod: [:FIFO, :LIFO, :FEFO]


  has_many :Sku, class_name: 'StockKeepingUnit'
  has_many :Warehouse, class_name: 'Warehouse'

end
