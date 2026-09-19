class SalesOrder < ApplicationRecord
  enum Status: [:Draft, :Confirmed, :Allocated, :InProduction, :Shipped, :Invoiced, :Closed, :Cancelled]


  composed_of :money,
    class_name: "Money",
    mapping: [
      ${$mapping}, 
      %w[money_currency currency]
    ]

  has_many :Customer, class_name: 'Customer'
  has_many :Plant, class_name: 'Plant'
  has_many :Lines, class_name: 'SalesOrderLine'
  has_many :WorkOrders, class_name: 'WorkOrder'

end
