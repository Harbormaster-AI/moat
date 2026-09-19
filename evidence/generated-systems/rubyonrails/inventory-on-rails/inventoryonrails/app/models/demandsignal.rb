class DemandSignal < ApplicationRecord
  enum DemandType: [:SalesOrder, :WorkOrder, :TransferOrder, :Forecast, :SampleRequest]


  has_many :Sku, class_name: 'StockKeepingUnit'
  has_many :Reservations, class_name: 'Reservation'

end
